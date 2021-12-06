package presenters

import (
	"time"

	"gopkg.in/guregu/null.v4"

	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	"github.com/smartcontractkit/chainlink/core/assets"
	clnull "github.com/smartcontractkit/chainlink/core/null"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
)

// JobSpecType defines the the the spec type of the job
type JobSpecType string

func (t JobSpecType) String() string {
	return string(t)
}

const (
	DirectRequestJobSpec     JobSpecType = "directrequest"
	FluxMonitorJobSpec       JobSpecType = "fluxmonitor"
	OffChainReportingJobSpec JobSpecType = "offchainreporting"
	KeeperJobSpec            JobSpecType = "keeper"
	CronJobSpec              JobSpecType = "cron"
	VRFJobSpec               JobSpecType = "vrf"
	WebhookJobSpec           JobSpecType = "webhook"
)

// DirectRequestSpec defines the spec details of a DirectRequest Job
type DirectRequestSpec struct {
	ContractAddress             ethkey.EIP55Address      `json:"contractAddress"`
	MinIncomingConfirmations    clnull.Uint32            `json:"minIncomingConfirmations"`
	MinIncomingConfirmationsEnv bool                     `json:"minIncomingConfirmationsEnv,omitempty"`
	MinContractPayment          *assets.Link             `json:"minContractPaymentLinkJuels"`
	Requesters                  models.AddressCollection `json:"requesters"`
	Initiator                   string                   `json:"initiator"`
	CreatedAt                   time.Time                `json:"createdAt"`
	UpdatedAt                   time.Time                `json:"updatedAt"`
	EVMChainID                  *utils.Big               `json:"evmChainID"`
}

// NewDirectRequestSpec initializes a new DirectRequestSpec from a
// job.DirectRequestSpec
func NewDirectRequestSpec(spec *job.DirectRequestSpec) *DirectRequestSpec {
	return &DirectRequestSpec{
		ContractAddress:             spec.ContractAddress,
		MinIncomingConfirmations:    spec.MinIncomingConfirmations,
		MinIncomingConfirmationsEnv: spec.MinIncomingConfirmationsEnv,
		MinContractPayment:          spec.MinContractPayment,
		Requesters:                  spec.Requesters,
		// This is hardcoded to runlog. When we support other intiators, we need
		// to change this
		Initiator:  "runlog",
		CreatedAt:  spec.CreatedAt,
		UpdatedAt:  spec.UpdatedAt,
		EVMChainID: spec.EVMChainID,
	}
}

// FluxMonitorSpec defines the spec details of a FluxMonitor Job
type FluxMonitorSpec struct {
	ContractAddress     ethkey.EIP55Address `json:"contractAddress"`
	Threshold           float32             `json:"threshold"`
	AbsoluteThreshold   float32             `json:"absoluteThreshold"`
	PollTimerPeriod     string              `json:"pollTimerPeriod"`
	PollTimerDisabled   bool                `json:"pollTimerDisabled"`
	IdleTimerPeriod     string              `json:"idleTimerPeriod"`
	IdleTimerDisabled   bool                `json:"idleTimerDisabled"`
	DrumbeatEnabled     bool                `json:"drumbeatEnabled"`
	DrumbeatSchedule    *string             `json:"drumbeatSchedule"`
	DrumbeatRandomDelay *string             `json:"drumbeatRandomDelay"`
	MinPayment          *assets.Link        `json:"minPayment"`
	CreatedAt           time.Time           `json:"createdAt"`
	UpdatedAt           time.Time           `json:"updatedAt"`
	EVMChainID          *utils.Big          `json:"evmChainID"`
}

// NewFluxMonitorSpec initializes a new DirectFluxMonitorSpec from a
// job.FluxMonitorSpec
func NewFluxMonitorSpec(spec *job.FluxMonitorSpec) *FluxMonitorSpec {
	var drumbeatSchedulePtr *string
	if spec.DrumbeatEnabled {
		drumbeatSchedulePtr = &spec.DrumbeatSchedule
	}
	var drumbeatRandomDelayPtr *string
	if spec.DrumbeatRandomDelay > 0 {
		drumbeatRandomDelay := spec.DrumbeatRandomDelay.String()
		drumbeatRandomDelayPtr = &drumbeatRandomDelay
	}
	return &FluxMonitorSpec{
		ContractAddress:     spec.ContractAddress,
		Threshold:           spec.Threshold,
		AbsoluteThreshold:   spec.AbsoluteThreshold,
		PollTimerPeriod:     spec.PollTimerPeriod.String(),
		PollTimerDisabled:   spec.PollTimerDisabled,
		IdleTimerPeriod:     spec.IdleTimerPeriod.String(),
		IdleTimerDisabled:   spec.IdleTimerDisabled,
		DrumbeatEnabled:     spec.DrumbeatEnabled,
		DrumbeatSchedule:    drumbeatSchedulePtr,
		DrumbeatRandomDelay: drumbeatRandomDelayPtr,
		MinPayment:          spec.MinPayment,
		CreatedAt:           spec.CreatedAt,
		UpdatedAt:           spec.UpdatedAt,
		EVMChainID:          spec.EVMChainID,
	}
}

// OffChainReportingSpec defines the spec details of a OffChainReporting Job
type OffChainReportingSpec struct {
	ContractAddress                           ethkey.EIP55Address  `json:"contractAddress"`
	P2PBootstrapPeers                         pq.StringArray       `json:"p2pBootstrapPeers"`
	IsBootstrapPeer                           bool                 `json:"isBootstrapPeer"`
	EncryptedOCRKeyBundleID                   *models.Sha256Hash   `json:"keyBundleID"`
	TransmitterAddress                        *ethkey.EIP55Address `json:"transmitterAddress"`
	ObservationTimeout                        models.Interval      `json:"observationTimeout"`
	ObservationTimeoutEnv                     bool                 `json:"observationTimeoutEnv,omitempty"`
	BlockchainTimeout                         models.Interval      `json:"blockchainTimeout"`
	BlockchainTimeoutEnv                      bool                 `json:"blockchainTimeoutEnv,omitempty"`
	ContractConfigTrackerSubscribeInterval    models.Interval      `json:"contractConfigTrackerSubscribeInterval"`
	ContractConfigTrackerSubscribeIntervalEnv bool                 `json:"contractConfigTrackerSubscribeIntervalEnv,omitempty"`
	ContractConfigTrackerPollInterval         models.Interval      `json:"contractConfigTrackerPollInterval"`
	ContractConfigTrackerPollIntervalEnv      bool                 `json:"contractConfigTrackerPollIntervalEnv,omitempty"`
	ContractConfigConfirmations               uint16               `json:"contractConfigConfirmations"`
	ContractConfigConfirmationsEnv            bool                 `json:"contractConfigConfirmationsEnv,omitempty"`
	CreatedAt                                 time.Time            `json:"createdAt"`
	UpdatedAt                                 time.Time            `json:"updatedAt"`
	EVMChainID                                *utils.Big           `json:"evmChainID"`
}

// NewOffChainReportingSpec initializes a new OffChainReportingSpec from a
// job.OffchainReportingOracleSpec
func NewOffChainReportingSpec(spec *job.OffchainReportingOracleSpec) *OffChainReportingSpec {
	return &OffChainReportingSpec{
		ContractAddress:                           spec.ContractAddress,
		P2PBootstrapPeers:                         spec.P2PBootstrapPeers,
		IsBootstrapPeer:                           spec.IsBootstrapPeer,
		EncryptedOCRKeyBundleID:                   spec.EncryptedOCRKeyBundleID,
		TransmitterAddress:                        spec.TransmitterAddress,
		ObservationTimeout:                        spec.ObservationTimeout,
		ObservationTimeoutEnv:                     spec.ObservationTimeoutEnv,
		BlockchainTimeout:                         spec.BlockchainTimeout,
		BlockchainTimeoutEnv:                      spec.BlockchainTimeoutEnv,
		ContractConfigTrackerSubscribeInterval:    spec.ContractConfigTrackerSubscribeInterval,
		ContractConfigTrackerSubscribeIntervalEnv: spec.ContractConfigTrackerSubscribeIntervalEnv,
		ContractConfigTrackerPollInterval:         spec.ContractConfigTrackerPollInterval,
		ContractConfigTrackerPollIntervalEnv:      spec.ContractConfigTrackerPollIntervalEnv,
		ContractConfigConfirmations:               spec.ContractConfigConfirmations,
		ContractConfigConfirmationsEnv:            spec.ContractConfigConfirmationsEnv,
		CreatedAt:                                 spec.CreatedAt,
		UpdatedAt:                                 spec.UpdatedAt,
		EVMChainID:                                spec.EVMChainID,
	}
}

// OffChainReporting2Spec defines the spec details of a OffChainReporting2 Job
type OffChainReporting2Spec struct {
	ContractAddress                        ethkey.EIP55Address  `json:"contractAddress"`
	P2PBootstrapPeers                      pq.StringArray       `json:"p2pBootstrapPeers"`
	IsBootstrapPeer                        bool                 `json:"isBootstrapPeer"`
	EncryptedOCRKeyBundleID                null.String          `json:"keyBundleID"`
	TransmitterAddress                     *ethkey.EIP55Address `json:"transmitterAddress"`
	ObservationTimeout                     models.Interval      `json:"observationTimeout"`
	BlockchainTimeout                      models.Interval      `json:"blockchainTimeout"`
	ContractConfigTrackerSubscribeInterval models.Interval      `json:"contractConfigTrackerSubscribeInterval"`
	ContractConfigTrackerPollInterval      models.Interval      `json:"contractConfigTrackerPollInterval"`
	ContractConfigConfirmations            uint16               `json:"contractConfigConfirmations"`
	CreatedAt                              time.Time            `json:"createdAt"`
	UpdatedAt                              time.Time            `json:"updatedAt"`
}

// NewOffChainReporting2Spec initializes a new OffChainReportingSpec from a
// job.OffchainReporting2OracleSpec
func NewOffChainReporting2Spec(spec *job.OffchainReporting2OracleSpec) *OffChainReporting2Spec {
	return &OffChainReporting2Spec{
		ContractAddress:                        spec.ContractAddress,
		P2PBootstrapPeers:                      spec.P2PBootstrapPeers,
		IsBootstrapPeer:                        spec.IsBootstrapPeer,
		EncryptedOCRKeyBundleID:                spec.EncryptedOCRKeyBundleID,
		TransmitterAddress:                     spec.TransmitterAddress,
		BlockchainTimeout:                      spec.BlockchainTimeout,
		ContractConfigTrackerSubscribeInterval: spec.ContractConfigTrackerSubscribeInterval,
		ContractConfigTrackerPollInterval:      spec.ContractConfigTrackerPollInterval,
		ContractConfigConfirmations:            spec.ContractConfigConfirmations,
		CreatedAt:                              spec.CreatedAt,
		UpdatedAt:                              spec.UpdatedAt,
	}
}

// PipelineSpec defines the spec details of the pipeline
type PipelineSpec struct {
	ID           int32  `json:"id"`
	JobID        int32  `json:"jobID"`
	DotDAGSource string `json:"dotDagSource"`
}

// NewPipelineSpec generates a new PipelineSpec from a pipeline.Spec
func NewPipelineSpec(spec *pipeline.Spec) PipelineSpec {
	return PipelineSpec{
		ID:           spec.ID,
		JobID:        spec.JobID,
		DotDAGSource: spec.DotDagSource,
	}
}

// KeeperSpec defines the spec details of a Keeper Job
type KeeperSpec struct {
	ContractAddress ethkey.EIP55Address `json:"contractAddress"`
	FromAddress     ethkey.EIP55Address `json:"fromAddress"`
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
	EVMChainID      *utils.Big          `json:"evmChainID"`
}

// NewKeeperSpec generates a new KeeperSpec from a job.KeeperSpec
func NewKeeperSpec(spec *job.KeeperSpec) *KeeperSpec {
	return &KeeperSpec{
		ContractAddress: spec.ContractAddress,
		FromAddress:     spec.FromAddress,
		CreatedAt:       spec.CreatedAt,
		UpdatedAt:       spec.UpdatedAt,
		EVMChainID:      spec.EVMChainID,
	}
}

// WebhookSpec defines the spec details of a Webhook Job
type WebhookSpec struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewWebhookSpec generates a new WebhookSpec from a job.WebhookSpec
func NewWebhookSpec(spec *job.WebhookSpec) *WebhookSpec {
	return &WebhookSpec{
		CreatedAt: spec.CreatedAt,
		UpdatedAt: spec.UpdatedAt,
	}
}

// CronSpec defines the spec details of a Cron Job
type CronSpec struct {
	CronSchedule string    `json:"schedule" tom:"schedule"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// NewCronSpec generates a new CronSpec from a job.CronSpec
func NewCronSpec(spec *job.CronSpec) *CronSpec {
	return &CronSpec{
		CronSchedule: spec.CronSchedule,
		CreatedAt:    spec.CreatedAt,
		UpdatedAt:    spec.UpdatedAt,
	}
}

type VRFSpec struct {
	CoordinatorAddress       ethkey.EIP55Address  `json:"coordinatorAddress"`
	PublicKey                secp256k1.PublicKey  `json:"publicKey"`
	FromAddress              *ethkey.EIP55Address `json:"fromAddress"`
	PollPeriod               models.Duration      `json:"pollPeriod"`
	MinIncomingConfirmations uint32               `json:"confirmations"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
	EVMChainID               *utils.Big           `json:"evmChainID"`
}

func NewVRFSpec(spec *job.VRFSpec) *VRFSpec {
	return &VRFSpec{
		CoordinatorAddress:       spec.CoordinatorAddress,
		PublicKey:                spec.PublicKey,
		FromAddress:              spec.FromAddress,
		PollPeriod:               models.MustMakeDuration(spec.PollPeriod),
		MinIncomingConfirmations: spec.MinIncomingConfirmations,
		CreatedAt:                spec.CreatedAt,
		UpdatedAt:                spec.UpdatedAt,
		EVMChainID:               spec.EVMChainID,
	}
}

// JobError represents errors on the job
type JobError struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Occurrences uint      `json:"occurrences"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewJobError(e job.SpecError) JobError {
	return JobError{
		ID:          e.ID,
		Description: e.Description,
		Occurrences: e.Occurrences,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

// JobResource represents a JobResource
type JobResource struct {
	JAID
	Name                   string                  `json:"name"`
	Type                   JobSpecType             `json:"type"`
	SchemaVersion          uint32                  `json:"schemaVersion"`
	MaxTaskDuration        models.Interval         `json:"maxTaskDuration"`
	ExternalJobID          uuid.UUID               `json:"externalJobID"`
	DirectRequestSpec      *DirectRequestSpec      `json:"directRequestSpec"`
	FluxMonitorSpec        *FluxMonitorSpec        `json:"fluxMonitorSpec"`
	CronSpec               *CronSpec               `json:"cronSpec"`
	OffChainReportingSpec  *OffChainReportingSpec  `json:"offChainReportingOracleSpec"`
	OffChainReporting2Spec *OffChainReporting2Spec `json:"offChainReporting2OracleSpec"`
	KeeperSpec             *KeeperSpec             `json:"keeperSpec"`
	VRFSpec                *VRFSpec                `json:"vrfSpec"`
	WebhookSpec            *WebhookSpec            `json:"webhookSpec"`
	PipelineSpec           PipelineSpec            `json:"pipelineSpec"`
	Errors                 []JobError              `json:"errors"`
}

// NewJobResource initializes a new JSONAPI job resource
func NewJobResource(j job.Job) *JobResource {
	resource := &JobResource{
		JAID:            NewJAIDInt32(j.ID),
		Name:            j.Name.ValueOrZero(),
		Type:            JobSpecType(j.Type),
		SchemaVersion:   j.SchemaVersion,
		MaxTaskDuration: j.MaxTaskDuration,
		PipelineSpec:    NewPipelineSpec(j.PipelineSpec),
		ExternalJobID:   j.ExternalJobID,
	}

	switch j.Type {
	case job.DirectRequest:
		resource.DirectRequestSpec = NewDirectRequestSpec(j.DirectRequestSpec)
	case job.FluxMonitor:
		resource.FluxMonitorSpec = NewFluxMonitorSpec(j.FluxMonitorSpec)
	case job.Cron:
		resource.CronSpec = NewCronSpec(j.CronSpec)
	case job.OffchainReporting:
		resource.OffChainReportingSpec = NewOffChainReportingSpec(j.OffchainreportingOracleSpec)
	case job.OffchainReporting2:
		resource.OffChainReporting2Spec = NewOffChainReporting2Spec(j.Offchainreporting2OracleSpec)
	case job.Keeper:
		resource.KeeperSpec = NewKeeperSpec(j.KeeperSpec)
	case job.VRF:
		resource.VRFSpec = NewVRFSpec(j.VRFSpec)
	case job.Webhook:
		resource.WebhookSpec = NewWebhookSpec(j.WebhookSpec)
	}

	jes := []JobError{}
	for _, e := range j.JobSpecErrors {
		jes = append(jes, NewJobError((e)))
	}
	resource.Errors = jes

	return resource
}

// GetName implements the api2go EntityNamer interface
func (r JobResource) GetName() string {
	return "jobs"
}
