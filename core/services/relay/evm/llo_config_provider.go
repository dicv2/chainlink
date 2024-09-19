package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/llo"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

// This is only used for the bootstrap job

var _ commontypes.ConfigProvider = (*lloConfigProvider)(nil)

type lloConfigProvider struct {
	services.Service
	eng *services.Engine

	lp              FilterRegisterer
	cp              llo.LLOConfigPollerService
	digester        ocrtypes.OffchainConfigDigester
	runReplay       bool
	replayFromBlock uint64
}

func (l *lloConfigProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return l.digester
}
func (l *lloConfigProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return l.cp
}

func newLLOConfigProvider(
	ctx context.Context,
	lggr logger.Logger,
	chain legacyevm.Chain,
	opts *types.RelayOpts,
) (commontypes.ConfigProvider, error) {
	if !common.IsHexAddress(opts.ContractID) {
		return nil, errors.New("invalid contractID, expected hex address")
	}

	configuratorAddress := common.HexToAddress(opts.ContractID)

	relayConfig, err := opts.RelayConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get relay config: %w", err)
	}
	donID := relayConfig.LLODONID
	if donID == 0 {
		return nil, errors.New("donID must be specified in relayConfig for LLO jobs")
	}
	donIDHash := llo.DonIDToBytes32(donID)

	lp := chain.LogPoller()
	err = lp.RegisterFilter(ctx, logpoller.Filter{Name: lloProviderConfiguratorFilterName(configuratorAddress, donID), EventSigs: []common.Hash{llo.ProductionConfigSet, llo.StagingConfigSet, llo.PromoteStagingConfig}, Topic2: []common.Hash{donIDHash}, Addresses: []common.Address{configuratorAddress}})
	if err != nil {
		return nil, fmt.Errorf("failed to register filter: %w", err)
	}

	configDigester := mercury.NewOffchainConfigDigester(llo.DonIDToBytes32(donID), chain.Config().EVM().ChainID(), configuratorAddress, ocrtypes.ConfigDigestPrefixLLO)
	cp := llo.NewConfigPoller(
		logger.Named(lggr, fmt.Sprintf("LLO-%d", donID)),
		chain.LogPoller(),
		configuratorAddress,
		donID,
		llo.InstanceTypeBlue,
		relayConfig.FromBlock,
	)
	// FIXME: green config tracking on the bootstrapper is NOT SUPPORTED YET
	// Need to generalize the provider to return multiple config trackers
	// This affects a lot of other plugins
	//
	// greenCP := llo.NewConfigPoller(
	//     logger.Named(lggr, fmt.Sprintf("LLO-%d", donID)),
	//     chain.LogPoller(),
	//     configuratorAddress,
	//     donID,
	//     llo.InstanceTypeGreen,
	// )
	p := &lloConfigProvider{nil, nil, lp, cp, configDigester, opts.New, relayConfig.FromBlock}
	p.Service, p.eng = services.Config{
		Name:  "LLOConfigProvider",
		Start: p.start,
		Close: p.cp.Close,
	}.NewServiceEngine(lggr)
	return p, nil
}

func (l *lloConfigProvider) start(ctx context.Context) error {
	if l.runReplay && l.replayFromBlock != 0 {
		// Only replay if it's a brand new job.
		l.eng.Go(func(ctx context.Context) {
			l.eng.Infow("starting replay for config", "fromBlock", l.replayFromBlock)
			if err := l.lp.Replay(ctx, int64(l.replayFromBlock)); err != nil {
				l.eng.Errorw("error replaying for config", "err", err)
			} else {
				l.eng.Infow("completed replaying for config", "replayFromBlock", l.replayFromBlock)
			}
		})
	}
	return l.cp.Start(ctx)
}
