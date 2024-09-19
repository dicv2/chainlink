package llo

import (
	"context"
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocr2plus "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	llotypes "github.com/smartcontractkit/chainlink-common/pkg/types/llo"
	"github.com/smartcontractkit/chainlink-data-streams/llo"
	datastreamsllo "github.com/smartcontractkit/chainlink-data-streams/llo"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/streams"
)

var _ job.ServiceCtx = &delegate{}

type Closer interface {
	Close() error
}

type delegate struct {
	services.StateMachine

	cfg    DelegateConfig
	codecs map[llotypes.ReportFormat]datastreamsllo.ReportCodec

	prrc datastreamsllo.PredecessorRetirementReportCache
	src  datastreamsllo.ShouldRetireCache
	ds   datastreamsllo.DataSource
	t    services.Service

	oracles [2]Closer
}

type LLOConfigTracker interface {
	// services.Service
	// ocr3types.ContractTransmitter[llotypes.ReportInfo]
	ocr2types.ContractConfigTracker
}

type DelegateConfig struct {
	Logger             logger.Logger
	DataSource         sqlutil.DataSource
	Runner             streams.Runner
	Registry           Registry
	JobName            null.String
	CaptureEATelemetry bool

	// LLO
	ChannelDefinitionCache llotypes.ChannelDefinitionCache
	ReportingPluginConfig  datastreamsllo.Config
	RetirementReportCache  RetirementReportCache
	ShouldRetireCache      datastreamsllo.ShouldRetireCache

	// OCR3
	TraceLogging                 bool
	BinaryNetworkEndpointFactory ocr2types.BinaryNetworkEndpointFactory
	V2Bootstrappers              []ocrcommontypes.BootstrapperLocator
	// ContractConfigTrackers is expected to contain both Blue and Green config
	// trackers in that order. One Oracle will be started for each.
	ContractConfigTrackers []ocr2types.ContractConfigTracker
	ContractTransmitter    ocr3types.ContractTransmitter[llotypes.ReportInfo]
	Database               ocr3types.Database
	MonitoringEndpoint     ocrcommontypes.MonitoringEndpoint
	OffchainConfigDigester ocr2types.OffchainConfigDigester
	OffchainKeyring        ocr2types.OffchainKeyring
	OnchainKeyring         ocr3types.OnchainKeyring[llotypes.ReportInfo]
	LocalConfig            ocr2types.LocalConfig
}

func NewDelegate(cfg DelegateConfig) (job.ServiceCtx, error) {
	lggr := logger.Sugared(cfg.Logger).With("jobName", cfg.JobName.ValueOrZero())
	if cfg.DataSource == nil {
		return nil, errors.New("DataSource must not be nil")
	}
	if cfg.Runner == nil {
		return nil, errors.New("Runner must not be nil")
	}
	if cfg.Registry == nil {
		return nil, errors.New("Registry must not be nil")
	}
	if cfg.RetirementReportCache == nil {
		return nil, errors.New("RetirementReportCache must not be nil")
	}
	if cfg.ShouldRetireCache == nil {
		return nil, errors.New("ShouldRetireCache must not be nil")
	}
	codecs := NewCodecs()

	// TODO: Do these services need starting?
	// https://smartcontract-it.atlassian.net/browse/MERC-3386
	var t TelemeterService
	if cfg.CaptureEATelemetry {
		t = NewTelemeterService(lggr, cfg.MonitoringEndpoint)
	} else {
		t = NullTelemeter
	}
	ds := newDataSource(logger.Named(lggr, "DataSource"), cfg.Registry, t)

	return &delegate{services.StateMachine{}, cfg, codecs, cfg.RetirementReportCache, cfg.ShouldRetireCache, ds, t, [2]Closer{}}, nil
}

func (d *delegate) Start(ctx context.Context) error {
	return d.StartOnce("LLODelegate", func() error {
		// create the oracle from config values
		var merr error
		for i, configTracker := range d.cfg.ContractConfigTrackers {
			lggr := logger.Named(d.cfg.Logger, fmt.Sprintf("%d", i))
			switch i {
			case 0:
				lggr = logger.With(lggr, "instanceType", "Blue")
			case 1:
				lggr = logger.With(lggr, "instanceType", "Green")
			}
			ocrLogger := logger.NewOCRWrapper(lggr, d.cfg.TraceLogging, func(msg string) {
				// TODO: do we actually need to DB-persist errors?
			})

			oracle, err := ocr2plus.NewOracle(ocr2plus.OCR3OracleArgs[llotypes.ReportInfo]{
				BinaryNetworkEndpointFactory: d.cfg.BinaryNetworkEndpointFactory,
				V2Bootstrappers:              d.cfg.V2Bootstrappers,
				ContractConfigTracker:        configTracker,
				ContractTransmitter:          d.cfg.ContractTransmitter,
				Database:                     d.cfg.Database,
				LocalConfig:                  d.cfg.LocalConfig,
				Logger:                       ocrLogger,
				MonitoringEndpoint:           d.cfg.MonitoringEndpoint,
				OffchainConfigDigester:       d.cfg.OffchainConfigDigester,
				OffchainKeyring:              d.cfg.OffchainKeyring,
				OnchainKeyring:               d.cfg.OnchainKeyring,
				ReportingPluginFactory: datastreamsllo.NewPluginFactory(
					d.cfg.ReportingPluginConfig, d.prrc, d.src, d.cfg.ChannelDefinitionCache, d.ds, logger.Named(lggr, "LLOReportingPlugin"), llo.StandardOnchainConfigCodec{}, d.codecs,
				),
				MetricsRegisterer: prometheus.WrapRegistererWith(map[string]string{"job_name": d.cfg.JobName.ValueOrZero()}, prometheus.DefaultRegisterer),
			})

			if err != nil {
				return fmt.Errorf("%w: failed to create new OCR oracle", err)
			}

			d.oracles[i] = oracle

			merr = errors.Join(merr, oracle.Start())
		}

		return merr
	})
}

func (d *delegate) Close() error {
	return d.StopOnce("LLODelegate", func() error {
		return errors.Join(d.oracles[0].Close(), d.oracles[1].Close())
	})
}
