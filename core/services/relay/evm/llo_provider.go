package evm

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/llo"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	relaytypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	llotypes "github.com/smartcontractkit/chainlink-common/pkg/types/llo"
)

var _ commontypes.LLOProvider = (*lloProvider)(nil)

type LLOTransmitter interface {
	services.Service
	llotypes.Transmitter
}

type FilterRegisterer interface {
	Replay(ctx context.Context, fromBlock int64) error
	RegisterFilter(ctx context.Context, filter logpoller.Filter) error
}

type lloProvider struct {
	services.Service
	eng *services.Engine

	cps []llo.LLOConfigPollerService

	transmitter            LLOTransmitter
	logger                 logger.Logger
	channelDefinitionCache llotypes.ChannelDefinitionCache
	digester               ocrtypes.OffchainConfigDigester
	shouldRetireCache      llo.ShouldRetireCacheService

	lp              FilterRegisterer
	runReplay       bool
	replayFromBlock uint64

	ms services.MultiStart
}

func lloProviderConfiguratorFilterName(addr common.Address, donID uint32) string {
	return logpoller.FilterName("LLOProvider Configurator", addr.String(), fmt.Sprintf("%d", donID))
}

func NewLLOProvider(
	ctx context.Context,
	transmitter LLOTransmitter,
	lggr logger.Logger,
	chain legacyevm.Chain,
	configuratorAddress common.Address,
	channelDefinitionCache llotypes.ChannelDefinitionCache,
	relayConfig types.RelayConfig,
	relayOpts *types.RelayOpts,
) (relaytypes.LLOProvider, error) {
	donID := relayConfig.LLODONID
	donIDHash := llo.DonIDToBytes32(donID)
	lp := chain.LogPoller()
	lggr = logger.Sugared(lggr).With("configuratorAddress", configuratorAddress, "donID", donID)

	err := lp.RegisterFilter(ctx, logpoller.Filter{Name: lloProviderConfiguratorFilterName(configuratorAddress, donID), EventSigs: []common.Hash{llo.ProductionConfigSet, llo.StagingConfigSet, llo.PromoteStagingConfig}, Topic2: []common.Hash{donIDHash}, Addresses: []common.Address{configuratorAddress}})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter: %w", err)
	}

	configDigester := mercury.NewOffchainConfigDigester(donIDHash, chain.Config().EVM().ChainID(), configuratorAddress, ocrtypes.ConfigDigestPrefixLLO)
	lggr = logger.Named(lggr, fmt.Sprintf("LLO-%d", donID))

	var cps []llo.LLOConfigPollerService
	switch relayConfig.LLOConfigMode {
	case types.LLOConfigModeMercury:
		mcp, err := mercury.NewConfigPoller(
			ctx,
			lggr,
			lp,
			configuratorAddress,
			llo.DonIDToBytes32(donID),
		)
		if err != nil {
			return nil, err
		}
		w := new(mercuryConfigPollerWrapper)
		w.ConfigPoller = mcp
		w.runReplay = relayOpts.New
		cps = []llo.LLOConfigPollerService{w}
	case types.LLOConfigModeBlueGreen:
		blueCP := llo.NewConfigPoller(
			lggr,
			lp,
			configuratorAddress,
			donID,
			llo.InstanceTypeBlue,
			relayConfig.FromBlock,
		)
		greenCP := llo.NewConfigPoller(
			lggr,
			chain.LogPoller(),
			configuratorAddress,
			donID,
			llo.InstanceTypeGreen,
			relayConfig.FromBlock,
		)
		cps = []llo.LLOConfigPollerService{blueCP, greenCP}
	}
	src := llo.NewShouldRetireCache(lggr, chain.LogPoller(), configuratorAddress, donID)

	p := &lloProvider{
		nil,
		nil,
		cps,
		transmitter,
		logger.Named(lggr, "LLOProvider"),
		channelDefinitionCache,
		configDigester,
		src,
		lp,
		relayOpts.New,
		relayConfig.FromBlock,
		services.MultiStart{},
	}

	p.Service, p.eng = services.Config{
		Name:  "LLOProvider",
		Start: p.start,
		Close: p.close,
	}.NewServiceEngine(lggr)

	return p, nil

}

func (p *lloProvider) start(ctx context.Context) error {
	// NOTE: Remember that all filters must be registered first for this replay
	// to be effective
	// 1. Replay
	// 2. Start all services
	if p.runReplay && p.replayFromBlock != 0 {
		// Only replay if it's a brand new job.
		p.eng.Go(func(ctx context.Context) {
			p.eng.Infow("starting replay for config", "fromBlock", p.replayFromBlock)
			if err := p.lp.Replay(ctx, int64(p.replayFromBlock)); err != nil {
				p.eng.Errorw("error replaying for config", "err", err)
			} else {
				p.eng.Infow("completed replaying for config", "replayFromBlock", p.replayFromBlock)
			}
		})
	}
	srvs := []services.StartClose{p.transmitter, p.channelDefinitionCache, p.shouldRetireCache}
	for _, cp := range p.cps {
		srvs = append(srvs, cp)
	}
	err := p.ms.Start(ctx, srvs...)
	return err
}

func (p *lloProvider) close() error {
	return p.ms.Close()
}

func (p *lloProvider) Ready() error {
	errs := make([]error, len(p.cps))
	for i, cp := range p.cps {
		errs[i] = cp.Ready()
	}
	errs = append(errs, p.transmitter.Ready(), p.channelDefinitionCache.Ready(), p.shouldRetireCache.Ready())
	return errors.Join(errs...)
}

func (p *lloProvider) Name() string {
	return p.logger.Name()
}

func (p *lloProvider) HealthReport() map[string]error {
	report := map[string]error{}
	for _, cp := range p.cps {
		services.CopyHealth(report, cp.HealthReport())
	}
	services.CopyHealth(report, p.transmitter.HealthReport())
	services.CopyHealth(report, p.channelDefinitionCache.HealthReport())
	services.CopyHealth(report, p.shouldRetireCache.HealthReport())
	return report
}

func (p *lloProvider) ContractConfigTrackers() (cps []ocrtypes.ContractConfigTracker) {
	cps = make([]ocrtypes.ContractConfigTracker, len(p.cps))
	for i, cp := range p.cps {
		cps[i] = cp
	}
	return
}

func (p *lloProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return p.digester
}

func (p *lloProvider) ContractTransmitter() llotypes.Transmitter {
	return p.transmitter
}

func (p *lloProvider) ChannelDefinitionCache() llotypes.ChannelDefinitionCache {
	return p.channelDefinitionCache
}

func (p *lloProvider) ShouldRetireCache() llotypes.ShouldRetireCache {
	return p.shouldRetireCache
}

// wrapper is needed to provide auto-replay support and turn mercury config
// poller into a service
type mercuryConfigPollerWrapper struct {
	*mercury.ConfigPoller
	services.Service
	eng *services.Engine

	runReplay bool
	fromBlock uint64
}

func newMercuryConfigPollerWrapper(lggr logger.Logger, cp *mercury.ConfigPoller, fromBlock uint64, runReplay bool) *mercuryConfigPollerWrapper {
	w := &mercuryConfigPollerWrapper{cp, nil, nil, runReplay, fromBlock}
	w.Service, w.eng = services.Config{
		Name:  "LLOMercuryConfigWrapper",
		Start: w.start,
		Close: w.close,
	}.NewServiceEngine(lggr)
	return w
}

func (w *mercuryConfigPollerWrapper) Start(ctx context.Context) error {
	return w.Service.Start(ctx)
}

func (w *mercuryConfigPollerWrapper) start(ctx context.Context) error {
	if w.runReplay && w.fromBlock != 0 {
		// Only replay if it's a brand new job.
		w.eng.Go(func(ctx context.Context) {
			w.eng.Infow("starting replay for config", "fromBlock", w.fromBlock)
			if err := w.ConfigPoller.Replay(ctx, int64(w.fromBlock)); err != nil {
				w.eng.Errorw("error replaying for config", "err", err)
			} else {
				w.eng.Infow("completed replaying for config", "fromBlock", w.fromBlock)
			}
		})
	}
	return nil
}

func (w *mercuryConfigPollerWrapper) Close() error {
	return w.Service.Close()
}

func (w *mercuryConfigPollerWrapper) close() error {
	return w.ConfigPoller.Close()
}
