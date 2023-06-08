// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	commontypes "github.com/smartcontractkit/libocr/commontypes"

	config "github.com/smartcontractkit/chainlink/v2/core/config"

	cosmos "github.com/smartcontractkit/chainlink/v2/core/chains/cosmos"

	ethkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"

	net "net"

	networking "github.com/smartcontractkit/libocr/networking"

	p2pkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"

	solana "github.com/smartcontractkit/chainlink/v2/core/chains/solana"

	starknet "github.com/smartcontractkit/chainlink/v2/core/chains/starknet"

	storemodels "github.com/smartcontractkit/chainlink/v2/core/store/models"

	time "time"

	url "net/url"

	uuid "github.com/google/uuid"

	v2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/v2"

	zapcore "go.uber.org/zap/zapcore"
)

// GeneralConfig is an autogenerated mock type for the GeneralConfig type
type GeneralConfig struct {
	mock.Mock
}

// AppID provides a mock function with given fields:
func (_m *GeneralConfig) AppID() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// AuditLogger provides a mock function with given fields:
func (_m *GeneralConfig) AuditLogger() config.AuditLogger {
	ret := _m.Called()

	var r0 config.AuditLogger
	if rf, ok := ret.Get(0).(func() config.AuditLogger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.AuditLogger)
		}
	}

	return r0
}

// AutoPprof provides a mock function with given fields:
func (_m *GeneralConfig) AutoPprof() config.AutoPprof {
	ret := _m.Called()

	var r0 config.AutoPprof
	if rf, ok := ret.Get(0).(func() config.AutoPprof); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.AutoPprof)
		}
	}

	return r0
}

// AutoPprofEnabled provides a mock function with given fields:
func (_m *GeneralConfig) AutoPprofEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ConfigTOML provides a mock function with given fields:
func (_m *GeneralConfig) ConfigTOML() (string, string) {
	ret := _m.Called()

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// CosmosConfigs provides a mock function with given fields:
func (_m *GeneralConfig) CosmosConfigs() cosmos.CosmosConfigs {
	ret := _m.Called()

	var r0 cosmos.CosmosConfigs
	if rf, ok := ret.Get(0).(func() cosmos.CosmosConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos.CosmosConfigs)
		}
	}

	return r0
}

// CosmosEnabled provides a mock function with given fields:
func (_m *GeneralConfig) CosmosEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Database provides a mock function with given fields:
func (_m *GeneralConfig) Database() config.Database {
	ret := _m.Called()

	var r0 config.Database
	if rf, ok := ret.Get(0).(func() config.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Database)
		}
	}

	return r0
}

// DefaultChainID provides a mock function with given fields:
func (_m *GeneralConfig) DefaultChainID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// EVMConfigs provides a mock function with given fields:
func (_m *GeneralConfig) EVMConfigs() v2.EVMConfigs {
	ret := _m.Called()

	var r0 v2.EVMConfigs
	if rf, ok := ret.Get(0).(func() v2.EVMConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v2.EVMConfigs)
		}
	}

	return r0
}

// EVMEnabled provides a mock function with given fields:
func (_m *GeneralConfig) EVMEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EVMRPCEnabled provides a mock function with given fields:
func (_m *GeneralConfig) EVMRPCEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EthereumHTTPURL provides a mock function with given fields:
func (_m *GeneralConfig) EthereumHTTPURL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// EthereumSecondaryURLs provides a mock function with given fields:
func (_m *GeneralConfig) EthereumSecondaryURLs() []url.URL {
	ret := _m.Called()

	var r0 []url.URL
	if rf, ok := ret.Get(0).(func() []url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]url.URL)
		}
	}

	return r0
}

// EthereumURL provides a mock function with given fields:
func (_m *GeneralConfig) EthereumURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ExplorerAccessKey provides a mock function with given fields:
func (_m *GeneralConfig) ExplorerAccessKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ExplorerSecret provides a mock function with given fields:
func (_m *GeneralConfig) ExplorerSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ExplorerURL provides a mock function with given fields:
func (_m *GeneralConfig) ExplorerURL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// FeatureExternalInitiators provides a mock function with given fields:
func (_m *GeneralConfig) FeatureExternalInitiators() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureFeedsManager provides a mock function with given fields:
func (_m *GeneralConfig) FeatureFeedsManager() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureLogPoller provides a mock function with given fields:
func (_m *GeneralConfig) FeatureLogPoller() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureOffchainReporting provides a mock function with given fields:
func (_m *GeneralConfig) FeatureOffchainReporting() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureOffchainReporting2 provides a mock function with given fields:
func (_m *GeneralConfig) FeatureOffchainReporting2() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeatureUICSAKeys provides a mock function with given fields:
func (_m *GeneralConfig) FeatureUICSAKeys() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FluxMonitor provides a mock function with given fields:
func (_m *GeneralConfig) FluxMonitor() config.FluxMonitor {
	ret := _m.Called()

	var r0 config.FluxMonitor
	if rf, ok := ret.Get(0).(func() config.FluxMonitor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.FluxMonitor)
		}
	}

	return r0
}

// Insecure provides a mock function with given fields:
func (_m *GeneralConfig) Insecure() config.Insecure {
	ret := _m.Called()

	var r0 config.Insecure
	if rf, ok := ret.Get(0).(func() config.Insecure); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Insecure)
		}
	}

	return r0
}

// InsecureFastScrypt provides a mock function with given fields:
func (_m *GeneralConfig) InsecureFastScrypt() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JobPipeline provides a mock function with given fields:
func (_m *GeneralConfig) JobPipeline() config.JobPipeline {
	ret := _m.Called()

	var r0 config.JobPipeline
	if rf, ok := ret.Get(0).(func() config.JobPipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.JobPipeline)
		}
	}

	return r0
}

// Keeper provides a mock function with given fields:
func (_m *GeneralConfig) Keeper() config.Keeper {
	ret := _m.Called()

	var r0 config.Keeper
	if rf, ok := ret.Get(0).(func() config.Keeper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Keeper)
		}
	}

	return r0
}

// KeystorePassword provides a mock function with given fields:
func (_m *GeneralConfig) KeystorePassword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Log provides a mock function with given fields:
func (_m *GeneralConfig) Log() config.Log {
	ret := _m.Called()

	var r0 config.Log
	if rf, ok := ret.Get(0).(func() config.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Log)
		}
	}

	return r0
}

// LogConfiguration provides a mock function with given fields: log
func (_m *GeneralConfig) LogConfiguration(log config.LogfFn) {
	_m.Called(log)
}

// MercuryCredentials provides a mock function with given fields: credName
func (_m *GeneralConfig) MercuryCredentials(credName string) *models.MercuryCredentials {
	ret := _m.Called(credName)

	var r0 *models.MercuryCredentials
	if rf, ok := ret.Get(0).(func(string) *models.MercuryCredentials); ok {
		r0 = rf(credName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MercuryCredentials)
		}
	}

	return r0
}

// OCR2BlockchainTimeout provides a mock function with given fields:
func (_m *GeneralConfig) OCR2BlockchainTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2CaptureEATelemetry provides a mock function with given fields:
func (_m *GeneralConfig) OCR2CaptureEATelemetry() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCR2ContractConfirmations provides a mock function with given fields:
func (_m *GeneralConfig) OCR2ContractConfirmations() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// OCR2ContractPollInterval provides a mock function with given fields:
func (_m *GeneralConfig) OCR2ContractPollInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2ContractSubscribeInterval provides a mock function with given fields:
func (_m *GeneralConfig) OCR2ContractSubscribeInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2ContractTransmitterTransmitTimeout provides a mock function with given fields:
func (_m *GeneralConfig) OCR2ContractTransmitterTransmitTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2DatabaseTimeout provides a mock function with given fields:
func (_m *GeneralConfig) OCR2DatabaseTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCR2DefaultTransactionQueueDepth provides a mock function with given fields:
func (_m *GeneralConfig) OCR2DefaultTransactionQueueDepth() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// OCR2KeyBundleID provides a mock function with given fields:
func (_m *GeneralConfig) OCR2KeyBundleID() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OCR2SimulateTransactions provides a mock function with given fields:
func (_m *GeneralConfig) OCR2SimulateTransactions() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCR2TraceLogging provides a mock function with given fields:
func (_m *GeneralConfig) OCR2TraceLogging() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRBlockchainTimeout provides a mock function with given fields:
func (_m *GeneralConfig) OCRBlockchainTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRCaptureEATelemetry provides a mock function with given fields:
func (_m *GeneralConfig) OCRCaptureEATelemetry() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRContractPollInterval provides a mock function with given fields:
func (_m *GeneralConfig) OCRContractPollInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRContractSubscribeInterval provides a mock function with given fields:
func (_m *GeneralConfig) OCRContractSubscribeInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRDefaultTransactionQueueDepth provides a mock function with given fields:
func (_m *GeneralConfig) OCRDefaultTransactionQueueDepth() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// OCRKeyBundleID provides a mock function with given fields:
func (_m *GeneralConfig) OCRKeyBundleID() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OCRObservationTimeout provides a mock function with given fields:
func (_m *GeneralConfig) OCRObservationTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OCRSimulateTransactions provides a mock function with given fields:
func (_m *GeneralConfig) OCRSimulateTransactions() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRTraceLogging provides a mock function with given fields:
func (_m *GeneralConfig) OCRTraceLogging() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OCRTransmitterAddress provides a mock function with given fields:
func (_m *GeneralConfig) OCRTransmitterAddress() (ethkey.EIP55Address, error) {
	ret := _m.Called()

	var r0 ethkey.EIP55Address
	var r1 error
	if rf, ok := ret.Get(0).(func() (ethkey.EIP55Address, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() ethkey.EIP55Address); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ethkey.EIP55Address)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// P2PAnnounceIP provides a mock function with given fields:
func (_m *GeneralConfig) P2PAnnounceIP() net.IP {
	ret := _m.Called()

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// P2PAnnouncePort provides a mock function with given fields:
func (_m *GeneralConfig) P2PAnnouncePort() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// P2PBootstrapCheckInterval provides a mock function with given fields:
func (_m *GeneralConfig) P2PBootstrapCheckInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// P2PBootstrapPeers provides a mock function with given fields:
func (_m *GeneralConfig) P2PBootstrapPeers() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// P2PDHTAnnouncementCounterUserPrefix provides a mock function with given fields:
func (_m *GeneralConfig) P2PDHTAnnouncementCounterUserPrefix() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// P2PDHTLookupInterval provides a mock function with given fields:
func (_m *GeneralConfig) P2PDHTLookupInterval() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// P2PEnabled provides a mock function with given fields:
func (_m *GeneralConfig) P2PEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// P2PIncomingMessageBufferSize provides a mock function with given fields:
func (_m *GeneralConfig) P2PIncomingMessageBufferSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// P2PListenIP provides a mock function with given fields:
func (_m *GeneralConfig) P2PListenIP() net.IP {
	ret := _m.Called()

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// P2PListenPort provides a mock function with given fields:
func (_m *GeneralConfig) P2PListenPort() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// P2PListenPortRaw provides a mock function with given fields:
func (_m *GeneralConfig) P2PListenPortRaw() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// P2PNetworkingStack provides a mock function with given fields:
func (_m *GeneralConfig) P2PNetworkingStack() networking.NetworkingStack {
	ret := _m.Called()

	var r0 networking.NetworkingStack
	if rf, ok := ret.Get(0).(func() networking.NetworkingStack); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(networking.NetworkingStack)
	}

	return r0
}

// P2PNetworkingStackRaw provides a mock function with given fields:
func (_m *GeneralConfig) P2PNetworkingStackRaw() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// P2PNewStreamTimeout provides a mock function with given fields:
func (_m *GeneralConfig) P2PNewStreamTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// P2POutgoingMessageBufferSize provides a mock function with given fields:
func (_m *GeneralConfig) P2POutgoingMessageBufferSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// P2PPeerID provides a mock function with given fields:
func (_m *GeneralConfig) P2PPeerID() p2pkey.PeerID {
	ret := _m.Called()

	var r0 p2pkey.PeerID
	if rf, ok := ret.Get(0).(func() p2pkey.PeerID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(p2pkey.PeerID)
	}

	return r0
}

// P2PPeerIDRaw provides a mock function with given fields:
func (_m *GeneralConfig) P2PPeerIDRaw() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// P2PPeerstoreWriteInterval provides a mock function with given fields:
func (_m *GeneralConfig) P2PPeerstoreWriteInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// P2PV2AnnounceAddresses provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2AnnounceAddresses() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// P2PV2Bootstrappers provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2Bootstrappers() []commontypes.BootstrapperLocator {
	ret := _m.Called()

	var r0 []commontypes.BootstrapperLocator
	if rf, ok := ret.Get(0).(func() []commontypes.BootstrapperLocator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]commontypes.BootstrapperLocator)
		}
	}

	return r0
}

// P2PV2BootstrappersRaw provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2BootstrappersRaw() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// P2PV2DeltaDial provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2DeltaDial() storemodels.Duration {
	ret := _m.Called()

	var r0 storemodels.Duration
	if rf, ok := ret.Get(0).(func() storemodels.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storemodels.Duration)
	}

	return r0
}

// P2PV2DeltaReconcile provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2DeltaReconcile() storemodels.Duration {
	ret := _m.Called()

	var r0 storemodels.Duration
	if rf, ok := ret.Get(0).(func() storemodels.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storemodels.Duration)
	}

	return r0
}

// P2PV2ListenAddresses provides a mock function with given fields:
func (_m *GeneralConfig) P2PV2ListenAddresses() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// PrometheusAuthToken provides a mock function with given fields:
func (_m *GeneralConfig) PrometheusAuthToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PyroscopeAuthToken provides a mock function with given fields:
func (_m *GeneralConfig) PyroscopeAuthToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PyroscopeEnvironment provides a mock function with given fields:
func (_m *GeneralConfig) PyroscopeEnvironment() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PyroscopeServerAddress provides a mock function with given fields:
func (_m *GeneralConfig) PyroscopeServerAddress() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RootDir provides a mock function with given fields:
func (_m *GeneralConfig) RootDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Sentry provides a mock function with given fields:
func (_m *GeneralConfig) Sentry() config.Sentry {
	ret := _m.Called()

	var r0 config.Sentry
	if rf, ok := ret.Get(0).(func() config.Sentry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Sentry)
		}
	}

	return r0
}

// SetLogLevel provides a mock function with given fields: lvl
func (_m *GeneralConfig) SetLogLevel(lvl zapcore.Level) error {
	ret := _m.Called(lvl)

	var r0 error
	if rf, ok := ret.Get(0).(func(zapcore.Level) error); ok {
		r0 = rf(lvl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogSQL provides a mock function with given fields: logSQL
func (_m *GeneralConfig) SetLogSQL(logSQL bool) {
	_m.Called(logSQL)
}

// SetPasswords provides a mock function with given fields: keystore, vrf
func (_m *GeneralConfig) SetPasswords(keystore *string, vrf *string) {
	_m.Called(keystore, vrf)
}

// ShutdownGracePeriod provides a mock function with given fields:
func (_m *GeneralConfig) ShutdownGracePeriod() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// SolanaConfigs provides a mock function with given fields:
func (_m *GeneralConfig) SolanaConfigs() solana.SolanaConfigs {
	ret := _m.Called()

	var r0 solana.SolanaConfigs
	if rf, ok := ret.Get(0).(func() solana.SolanaConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(solana.SolanaConfigs)
		}
	}

	return r0
}

// SolanaEnabled provides a mock function with given fields:
func (_m *GeneralConfig) SolanaEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StarkNetEnabled provides a mock function with given fields:
func (_m *GeneralConfig) StarkNetEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StarknetConfigs provides a mock function with given fields:
func (_m *GeneralConfig) StarknetConfigs() starknet.StarknetConfigs {
	ret := _m.Called()

	var r0 starknet.StarknetConfigs
	if rf, ok := ret.Get(0).(func() starknet.StarknetConfigs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(starknet.StarknetConfigs)
		}
	}

	return r0
}

// TelemetryIngress provides a mock function with given fields:
func (_m *GeneralConfig) TelemetryIngress() config.TelemetryIngress {
	ret := _m.Called()

	var r0 config.TelemetryIngress
	if rf, ok := ret.Get(0).(func() config.TelemetryIngress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.TelemetryIngress)
		}
	}

	return r0
}

// ThresholdKeyShare provides a mock function with given fields:
func (_m *GeneralConfig) ThresholdKeyShare() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VRFPassword provides a mock function with given fields:
func (_m *GeneralConfig) VRFPassword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *GeneralConfig) Validate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDB provides a mock function with given fields:
func (_m *GeneralConfig) ValidateDB() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WebServer provides a mock function with given fields:
func (_m *GeneralConfig) WebServer() config.WebServer {
	ret := _m.Called()

	var r0 config.WebServer
	if rf, ok := ret.Get(0).(func() config.WebServer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.WebServer)
		}
	}

	return r0
}

type mockConstructorTestingTNewGeneralConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralConfig creates a new instance of GeneralConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralConfig(t mockConstructorTestingTNewGeneralConfig) *GeneralConfig {
	mock := &GeneralConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
