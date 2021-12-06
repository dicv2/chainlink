package resolver

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/smartcontractkit/chainlink/core/services/job"
)

type SpecResolver struct {
	j job.Job
}

func NewSpec(j job.Job) *SpecResolver {
	return &SpecResolver{j: j}
}

func (r *SpecResolver) ToCronSpec() (*CronSpecResolver, bool) {
	if r.j.Type != job.Cron {
		return nil, false
	}

	return &CronSpecResolver{spec: *r.j.CronSpec}, true
}

func (r *SpecResolver) ToDirectRequestSpec() (*DirectRequestSpecResolver, bool) {
	if r.j.Type != job.DirectRequest {
		return nil, false
	}

	return &DirectRequestSpecResolver{spec: *r.j.DirectRequestSpec}, true
}

func (r *SpecResolver) ToFluxMonitorSpec() (*FluxMonitorSpecResolver, bool) {
	if r.j.Type != job.FluxMonitor {
		return nil, false
	}

	return &FluxMonitorSpecResolver{spec: *r.j.FluxMonitorSpec}, true
}

func (r *SpecResolver) ToKeeperSpec() (*KeeperSpecResolver, bool) {
	if r.j.Type != job.Keeper {
		return nil, false
	}

	return &KeeperSpecResolver{spec: *r.j.KeeperSpec}, true
}

func (r *SpecResolver) ToOCRSpec() (*OCRSpecResolver, bool) {
	if r.j.Type != job.OffchainReporting {
		return nil, false
	}

	return &OCRSpecResolver{spec: *r.j.OffchainreportingOracleSpec}, true
}

func (r *SpecResolver) ToOCR2Spec() (*OCR2SpecResolver, bool) {
	if r.j.Type != job.OffchainReporting2 {
		return nil, false
	}

	return &OCR2SpecResolver{spec: *r.j.Offchainreporting2OracleSpec}, true
}

func (r *SpecResolver) ToVRFSpec() (*VRFSpecResolver, bool) {
	if r.j.Type != job.VRF {
		return nil, false
	}

	return &VRFSpecResolver{spec: *r.j.VRFSpec}, true
}

func (r *SpecResolver) ToWebhookSpec() (*WebhookSpecResolver, bool) {
	if r.j.Type != job.Webhook {
		return nil, false
	}

	return &WebhookSpecResolver{spec: *r.j.WebhookSpec}, true
}

type CronSpecResolver struct {
	spec job.CronSpec
}

func (r *CronSpecResolver) Schedule() string {
	return r.spec.CronSchedule
}

// CreatedAt resolves the spec's created at timestamp.
func (r *CronSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

type DirectRequestSpecResolver struct {
	spec job.DirectRequestSpec
}

// ContractAddress resolves the spec's contract address.
func (r *DirectRequestSpecResolver) ContractAddress() string {
	return r.spec.ContractAddress.String()
}

// CreatedAt resolves the spec's created at timestamp.
func (r *DirectRequestSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// EVMChainID resolves the spec's evm chain id.
func (r *DirectRequestSpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// MinIncomingConfirmations resolves the spec's min incoming confirmations.
func (r *DirectRequestSpecResolver) MinIncomingConfirmations() int32 {
	if r.spec.MinIncomingConfirmations.Valid {
		return int32(r.spec.MinIncomingConfirmations.Uint32)
	}

	return 0
}

// EVMChainID resolves the spec's evm chain id.
func (r *DirectRequestSpecResolver) MinIncomingConfirmationsEnv() bool {
	return r.spec.MinIncomingConfirmationsEnv
}

// MinContractPayment resolves the spec's evm chain id.
func (r *DirectRequestSpecResolver) MinContractPayment() string {
	return r.spec.MinContractPayment.String()
}

// Requesters resolves the spec's evm chain id.
func (r *DirectRequestSpecResolver) Requesters() *[]string {
	if r.spec.Requesters == nil {
		return nil
	}

	requesters := r.spec.Requesters.ToStrings()

	return &requesters
}

type FluxMonitorSpecResolver struct {
	spec job.FluxMonitorSpec
}

// AbsoluteThreshold resolves the spec's absolute deviation threshold.
func (r *FluxMonitorSpecResolver) AbsoluteThreshold() float64 {
	return float64(r.spec.AbsoluteThreshold)
}

// ContractAddress resolves the spec's contract address.
func (r *FluxMonitorSpecResolver) ContractAddress() string {
	return r.spec.ContractAddress.String()
}

// CreatedAt resolves the spec's created at timestamp.
func (r *FluxMonitorSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// AbsoluteThreshold resolves the spec's absolute threshold.
func (r *FluxMonitorSpecResolver) DrumbeatEnabled() bool {
	return r.spec.DrumbeatEnabled
}

// DrumbeatRandomDelay resolves the spec's drumbeat random delay.
func (r *FluxMonitorSpecResolver) DrumbeatRandomDelay() *string {
	var delay *string
	if r.spec.DrumbeatRandomDelay > 0 {
		drumbeatRandomDelay := r.spec.DrumbeatRandomDelay.String()
		delay = &drumbeatRandomDelay
	}

	return delay
}

// DrumbeatSchedule resolves the spec's drumbeat schedule.
func (r *FluxMonitorSpecResolver) DrumbeatSchedule() *string {
	if r.spec.DrumbeatEnabled {
		return &r.spec.DrumbeatSchedule
	}

	return nil
}

// EVMChainID resolves the spec's evm chain id.
func (r *FluxMonitorSpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// IdleTimerDisabled resolves the spec's idle timer disabled flag.
func (r *FluxMonitorSpecResolver) IdleTimerDisabled() bool {
	return r.spec.IdleTimerDisabled
}

// IdleTimerPeriod resolves the spec's idle timer period.
func (r *FluxMonitorSpecResolver) IdleTimerPeriod() string {
	return r.spec.IdleTimerPeriod.String()
}

// MinPayment resolves the spec's min payment.
func (r *FluxMonitorSpecResolver) MinPayment() *string {
	if r.spec.MinPayment != nil {
		min := r.spec.MinPayment.String()

		return &min
	}
	return nil
}

// PollTimerDisabled resolves the spec's poll timer disabled flag.
func (r *FluxMonitorSpecResolver) PollTimerDisabled() bool {
	return r.spec.PollTimerDisabled
}

// PollTimerPeriod resolves the spec's poll timer period.
func (r *FluxMonitorSpecResolver) PollTimerPeriod() string {
	return r.spec.PollTimerPeriod.String()
}

// Threshold resolves the spec's deviation threshold.
func (r *FluxMonitorSpecResolver) Threshold() float64 {
	return float64(r.spec.Threshold)
}

type KeeperSpecResolver struct {
	spec job.KeeperSpec
}

// ContractAddress resolves the spec's contract address.
func (r *KeeperSpecResolver) ContractAddress() string {
	return r.spec.ContractAddress.String()
}

// CreatedAt resolves the spec's created at timestamp.
func (r *KeeperSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// EVMChainID resolves the spec's evm chain id.
func (r *KeeperSpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// FromAddress resolves the spec's from contract address.
//
// Because VRF has an non required field of the same name, we have to be
// consistent in our return value of using a *string instead of a string even
// though this is a required field for the KeeperSpec.
//
// http://spec.graphql.org/draft/#sec-Field-Selection-Merging
func (r *KeeperSpecResolver) FromAddress() *string {
	addr := r.spec.FromAddress.String()

	return &addr
}

type OCRSpecResolver struct {
	spec job.OffchainReportingOracleSpec
}

// BlockchainTimeout resolves the spec's blockchain timeout.
func (r *OCRSpecResolver) BlockchainTimeout() *string {
	if r.spec.BlockchainTimeout.Duration() == 0 {
		return nil
	}

	timeout := r.spec.BlockchainTimeout.Duration().String()

	return &timeout
}

// BlockchainTimeoutEnv resolves whether the spec's blockchain timeout comes
// from an env var.
func (r *OCRSpecResolver) BlockchainTimeoutEnv() bool {
	return r.spec.BlockchainTimeoutEnv
}

// ContractAddress resolves the spec's contract address.
func (r *OCRSpecResolver) ContractAddress() string {
	return r.spec.ContractAddress.String()
}

// ContractConfigConfirmations resolves the spec's confirmations config.
func (r *OCRSpecResolver) ContractConfigConfirmations() *int32 {
	if r.spec.ContractConfigConfirmations == 0 {
		return nil
	}

	confirmations := int32(r.spec.ContractConfigConfirmations)

	return &confirmations
}

// ContractConfigConfirmationsEnv resolves whether spec's confirmations
// config comes from an env var.
func (r *OCRSpecResolver) ContractConfigConfirmationsEnv() bool {
	return r.spec.ContractConfigConfirmationsEnv
}

// ContractConfigTrackerPollInterval resolves the spec's contract tracker poll
// interval config.
func (r *OCRSpecResolver) ContractConfigTrackerPollInterval() *string {
	if r.spec.ContractConfigTrackerPollInterval.Duration() == 0 {
		return nil
	}

	interval := r.spec.ContractConfigTrackerPollInterval.Duration().String()

	return &interval
}

// ContractConfigTrackerPollIntervalEnv resolves the whether spec's tracker poll
// config comes from an env var.
func (r *OCRSpecResolver) ContractConfigTrackerPollIntervalEnv() bool {
	return r.spec.ContractConfigTrackerPollIntervalEnv
}

// ContractConfigTrackerSubscribeInterval resolves the spec's tracker subscribe
// interval config.
func (r *OCRSpecResolver) ContractConfigTrackerSubscribeInterval() *string {
	if r.spec.ContractConfigTrackerSubscribeInterval.Duration() == 0 {
		return nil
	}

	interval := r.spec.ContractConfigTrackerSubscribeInterval.Duration().String()

	return &interval
}

// ContractConfigTrackerSubscribeIntervalEnv resolves whether spec's tracker
// subscribe interval config comes from an env var.
func (r *OCRSpecResolver) ContractConfigTrackerSubscribeIntervalEnv() bool {
	return r.spec.ContractConfigTrackerSubscribeIntervalEnv
}

// CreatedAt resolves the spec's created at timestamp.
func (r *OCRSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// EVMChainID resolves the spec's evm chain id.
func (r *OCRSpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// IsBootstrapPeer resolves whether spec is a bootstrap peer.
func (r *OCRSpecResolver) IsBootstrapPeer() bool {
	return r.spec.IsBootstrapPeer
}

// KeyBundleID resolves the spec's key bundle id.
func (r *OCRSpecResolver) KeyBundleID() *string {
	if r.spec.EncryptedOCRKeyBundleID == nil {
		return nil
	}

	bundleID := r.spec.EncryptedOCRKeyBundleID.String()

	return &bundleID
}

// ObservationTimeout resolves the spec's observation timeout
func (r *OCRSpecResolver) ObservationTimeout() *string {
	if r.spec.ObservationTimeout.Duration() == 0 {
		return nil
	}

	timeout := r.spec.ObservationTimeout.Duration().String()

	return &timeout
}

// ObservationTimeoutEnv resolves whether spec's observation timeout comes
// from an env var.
func (r *OCRSpecResolver) ObservationTimeoutEnv() bool {
	return r.spec.ObservationTimeoutEnv
}

// P2PBootstrapPeers resolves the spec's p2p bootstrap peers
func (r *OCRSpecResolver) P2PBootstrapPeers() *[]string {
	if len(r.spec.P2PBootstrapPeers) == 0 {
		return nil
	}

	peers := []string(r.spec.P2PBootstrapPeers)

	return &peers
}

// TransmitterAddress resolves the spec's transmitter address
func (r *OCRSpecResolver) TransmitterAddress() *string {
	if r.spec.TransmitterAddress == nil {
		return nil
	}

	addr := r.spec.TransmitterAddress.String()
	return &addr
}

type OCR2SpecResolver struct {
	spec job.OffchainReporting2OracleSpec
}

// BlockchainTimeout resolves the spec's blockchain timeout.
func (r *OCR2SpecResolver) BlockchainTimeout() *string {
	if r.spec.BlockchainTimeout.Duration() == 0 {
		return nil
	}

	timeout := r.spec.BlockchainTimeout.Duration().String()

	return &timeout
}

// ContractAddress resolves the spec's contract address.
func (r *OCR2SpecResolver) ContractAddress() string {
	return r.spec.ContractAddress.String()
}

// ContractConfigConfirmations resolves the spec's confirmations config.
func (r *OCR2SpecResolver) ContractConfigConfirmations() *int32 {
	if r.spec.ContractConfigConfirmations == 0 {
		return nil
	}

	confirmations := int32(r.spec.ContractConfigConfirmations)

	return &confirmations
}

// ContractConfigTrackerPollInterval resolves the spec's contract tracker poll
// interval config.
func (r *OCR2SpecResolver) ContractConfigTrackerPollInterval() *string {
	if r.spec.ContractConfigTrackerPollInterval.Duration() == 0 {
		return nil
	}

	interval := r.spec.ContractConfigTrackerPollInterval.Duration().String()

	return &interval
}

// ContractConfigTrackerSubscribeInterval resolves the spec's tracker subscribe
// interval config.
func (r *OCR2SpecResolver) ContractConfigTrackerSubscribeInterval() *string {
	if r.spec.ContractConfigTrackerSubscribeInterval.Duration() == 0 {
		return nil
	}

	interval := r.spec.ContractConfigTrackerSubscribeInterval.Duration().String()

	return &interval
}

// CreatedAt resolves the spec's created at timestamp.
func (r *OCR2SpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// EVMChainID resolves the spec's evm chain id.
func (r *OCR2SpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// IsBootstrapPeer resolves whether spec is a bootstrap peer.
func (r *OCR2SpecResolver) IsBootstrapPeer() bool {
	return r.spec.IsBootstrapPeer
}

// JuelsPerFeeCoinSource resolves the spec's jeuls per fee coin source
func (r *OCR2SpecResolver) JuelsPerFeeCoinSource() *string {
	if r.spec.JuelsPerFeeCoinPipeline == "" {
		return nil
	}

	return &r.spec.JuelsPerFeeCoinPipeline
}

// KeyBundleID resolves the spec's key bundle id.
func (r *OCR2SpecResolver) KeyBundleID() *string {
	if !r.spec.EncryptedOCRKeyBundleID.Valid {
		return nil
	}

	return &r.spec.EncryptedOCRKeyBundleID.String
}

// MonitoringEndpoint resolves the spec's monitoring endpoint
func (r *OCR2SpecResolver) MonitoringEndpoint() *string {
	if !r.spec.MonitoringEndpoint.Valid {
		return nil
	}

	return &r.spec.MonitoringEndpoint.String
}

// P2PBootstrapPeers resolves the spec's p2p bootstrap peers
func (r *OCR2SpecResolver) P2PBootstrapPeers() *[]string {
	if len(r.spec.P2PBootstrapPeers) == 0 {
		return nil
	}

	peers := []string(r.spec.P2PBootstrapPeers)

	return &peers
}

// TransmitterAddress resolves the spec's transmitter address
func (r *OCR2SpecResolver) TransmitterAddress() *string {
	if r.spec.TransmitterAddress == nil {
		return nil
	}

	addr := r.spec.TransmitterAddress.String()
	return &addr
}

type VRFSpecResolver struct {
	spec job.VRFSpec
}

// MinIncomingConfirmations resolves the spec's min incoming confirmations.
func (r *VRFSpecResolver) MinIncomingConfirmations() int32 {
	return int32(r.spec.MinIncomingConfirmations)
}

// MinIncomingConfirmations resolves the spec's min incoming confirmations.
func (r *VRFSpecResolver) MinIncomingConfirmationsEnv() bool {
	return r.spec.ConfirmationsEnv
}

// CoordinatorAddress resolves the spec's coordinator address.
func (r *VRFSpecResolver) CoordinatorAddress() string {
	return r.spec.CoordinatorAddress.String()
}

// CreatedAt resolves the spec's created at timestamp.
func (r *VRFSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}

// EVMChainID resolves the spec's evm chain id.
func (r *VRFSpecResolver) EVMChainID() *string {
	if r.spec.EVMChainID == nil {
		return nil
	}

	chainID := r.spec.EVMChainID.String()

	return &chainID
}

// FromAddress resolves the spec's from address.
func (r *VRFSpecResolver) FromAddress() *string {
	if r.spec.FromAddress == nil {
		return nil
	}

	addr := r.spec.FromAddress.String()
	return &addr
}

// PollPeriod resolves the spec's poll period.
func (r *VRFSpecResolver) PollPeriod() string {
	return r.spec.PollPeriod.String()
}

// PublicKey resolves the spec's public key.
func (r *VRFSpecResolver) PublicKey() string {
	return r.spec.PublicKey.String()
}

// RequestedConfsDelay resolves the spec's requested conf delay.
func (r *VRFSpecResolver) RequestedConfsDelay() int32 {
	// GraphQL doesn't support 64 bit integers, so we have to cast.
	return int32(r.spec.RequestedConfsDelay)
}

// RequestTimeout resolves the spec's request timeout.
func (r *VRFSpecResolver) RequestTimeout() string {
	return r.spec.RequestTimeout.String()
}

type WebhookSpecResolver struct {
	spec job.WebhookSpec
}

// CreatedAt resolves the spec's created at timestamp.
func (r *WebhookSpecResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.spec.CreatedAt}
}
