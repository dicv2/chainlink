// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	logpoller "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// LogPoller is an autogenerated mock type for the LogPoller type
type LogPoller struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *LogPoller) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlocksRange provides a mock function with given fields: ctx, numbers
func (_m *LogPoller) GetBlocksRange(ctx context.Context, numbers []uint64) ([]logpoller.LogPollerBlock, error) {
	ret := _m.Called(ctx, numbers)

	if len(ret) == 0 {
		panic("no return value specified for GetBlocksRange")
	}

	var r0 []logpoller.LogPollerBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint64) ([]logpoller.LogPollerBlock, error)); ok {
		return rf(ctx, numbers)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint64) []logpoller.LogPollerBlock); ok {
		r0 = rf(ctx, numbers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.LogPollerBlock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint64) error); ok {
		r1 = rf(ctx, numbers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasFilter provides a mock function with given fields: name
func (_m *LogPoller) HasFilter(name string) bool {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for HasFilter")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *LogPoller) HealthReport() map[string]error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HealthReport")
	}

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *LogPoller) Healthy() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Healthy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexedLogs provides a mock function with given fields: ctx, eventSig, address, topicIndex, topicValues, confs
func (_m *LogPoller) IndexedLogs(ctx context.Context, eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, topicIndex, topicValues, confs)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogs")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, []common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, topicIndex, topicValues, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, []common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, topicIndex, topicValues, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, []common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, topicIndex, topicValues, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsByBlockRange provides a mock function with given fields: ctx, start, end, eventSig, address, topicIndex, topicValues
func (_m *LogPoller) IndexedLogsByBlockRange(ctx context.Context, start int64, end int64, eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, start, end, eventSig, address, topicIndex, topicValues)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsByBlockRange")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, common.Hash, common.Address, int, []common.Hash) ([]logpoller.Log, error)); ok {
		return rf(ctx, start, end, eventSig, address, topicIndex, topicValues)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, common.Hash, common.Address, int, []common.Hash) []logpoller.Log); ok {
		r0 = rf(ctx, start, end, eventSig, address, topicIndex, topicValues)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, common.Hash, common.Address, int, []common.Hash) error); ok {
		r1 = rf(ctx, start, end, eventSig, address, topicIndex, topicValues)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsByTxHash provides a mock function with given fields: ctx, eventSig, address, txHash
func (_m *LogPoller) IndexedLogsByTxHash(ctx context.Context, eventSig common.Hash, address common.Address, txHash common.Hash) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, txHash)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsByTxHash")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, common.Hash) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, common.Hash) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, eventSig, address, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsCreatedAfter provides a mock function with given fields: ctx, eventSig, address, topicIndex, topicValues, after, confs
func (_m *LogPoller) IndexedLogsCreatedAfter(ctx context.Context, eventSig common.Hash, address common.Address, topicIndex int, topicValues []common.Hash, after time.Time, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, topicIndex, topicValues, after, confs)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsCreatedAfter")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, []common.Hash, time.Time, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, topicIndex, topicValues, after, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, []common.Hash, time.Time, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, topicIndex, topicValues, after, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, []common.Hash, time.Time, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, topicIndex, topicValues, after, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsTopicGreaterThan provides a mock function with given fields: ctx, eventSig, address, topicIndex, topicValueMin, confs
func (_m *LogPoller) IndexedLogsTopicGreaterThan(ctx context.Context, eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, topicIndex, topicValueMin, confs)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsTopicGreaterThan")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, topicIndex, topicValueMin, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, topicIndex, topicValueMin, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, topicIndex, topicValueMin, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsTopicRange provides a mock function with given fields: ctx, eventSig, address, topicIndex, topicValueMin, topicValueMax, confs
func (_m *LogPoller) IndexedLogsTopicRange(ctx context.Context, eventSig common.Hash, address common.Address, topicIndex int, topicValueMin common.Hash, topicValueMax common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, topicIndex, topicValueMin, topicValueMax, confs)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsTopicRange")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, topicIndex, topicValueMin, topicValueMax, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, topicIndex, topicValueMin, topicValueMax, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, topicIndex, topicValueMin, topicValueMax, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexedLogsWithSigsExcluding provides a mock function with given fields: ctx, address, eventSigA, eventSigB, topicIndex, fromBlock, toBlock, confs
func (_m *LogPoller) IndexedLogsWithSigsExcluding(ctx context.Context, address common.Address, eventSigA common.Hash, eventSigB common.Hash, topicIndex int, fromBlock int64, toBlock int64, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, address, eventSigA, eventSigB, topicIndex, fromBlock, toBlock, confs)

	if len(ret) == 0 {
		panic("no return value specified for IndexedLogsWithSigsExcluding")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash, common.Hash, int, int64, int64, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, address, eventSigA, eventSigB, topicIndex, fromBlock, toBlock, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash, common.Hash, int, int64, int64, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, address, eventSigA, eventSigB, topicIndex, fromBlock, toBlock, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash, common.Hash, int, int64, int64, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, address, eventSigA, eventSigB, topicIndex, fromBlock, toBlock, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestBlock provides a mock function with given fields: ctx
func (_m *LogPoller) LatestBlock(ctx context.Context) (logpoller.LogPollerBlock, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestBlock")
	}

	var r0 logpoller.LogPollerBlock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (logpoller.LogPollerBlock, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) logpoller.LogPollerBlock); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(logpoller.LogPollerBlock)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestBlockByEventSigsAddrsWithConfs provides a mock function with given fields: ctx, fromBlock, eventSigs, addresses, confs
func (_m *LogPoller) LatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs logpoller.Confirmations) (int64, error) {
	ret := _m.Called(ctx, fromBlock, eventSigs, addresses, confs)

	if len(ret) == 0 {
		panic("no return value specified for LatestBlockByEventSigsAddrsWithConfs")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) (int64, error)); ok {
		return rf(ctx, fromBlock, eventSigs, addresses, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) int64); ok {
		r0 = rf(ctx, fromBlock, eventSigs, addresses, confs)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, fromBlock, eventSigs, addresses, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestLogByEventSigWithConfs provides a mock function with given fields: ctx, eventSig, address, confs
func (_m *LogPoller) LatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs logpoller.Confirmations) (*logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, confs)

	if len(ret) == 0 {
		panic("no return value specified for LatestLogByEventSigWithConfs")
	}

	var r0 *logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, logpoller.Confirmations) (*logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, logpoller.Confirmations) *logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestLogEventSigsAddrsWithConfs provides a mock function with given fields: ctx, fromBlock, eventSigs, addresses, confs
func (_m *LogPoller) LatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, fromBlock, eventSigs, addresses, confs)

	if len(ret) == 0 {
		panic("no return value specified for LatestLogEventSigsAddrsWithConfs")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, fromBlock, eventSigs, addresses, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, fromBlock, eventSigs, addresses, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []common.Hash, []common.Address, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, fromBlock, eventSigs, addresses, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logs provides a mock function with given fields: ctx, start, end, eventSig, address
func (_m *LogPoller) Logs(ctx context.Context, start int64, end int64, eventSig common.Hash, address common.Address) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, start, end, eventSig, address)

	if len(ret) == 0 {
		panic("no return value specified for Logs")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, common.Hash, common.Address) ([]logpoller.Log, error)); ok {
		return rf(ctx, start, end, eventSig, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, common.Hash, common.Address) []logpoller.Log); ok {
		r0 = rf(ctx, start, end, eventSig, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, common.Hash, common.Address) error); ok {
		r1 = rf(ctx, start, end, eventSig, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsCreatedAfter provides a mock function with given fields: ctx, eventSig, address, _a3, confs
func (_m *LogPoller) LogsCreatedAfter(ctx context.Context, eventSig common.Hash, address common.Address, _a3 time.Time, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, _a3, confs)

	if len(ret) == 0 {
		panic("no return value specified for LogsCreatedAfter")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, time.Time, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, _a3, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, time.Time, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, _a3, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, time.Time, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, _a3, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsDataWordBetween provides a mock function with given fields: ctx, eventSig, address, wordIndexMin, wordIndexMax, wordValue, confs
func (_m *LogPoller) LogsDataWordBetween(ctx context.Context, eventSig common.Hash, address common.Address, wordIndexMin int, wordIndexMax int, wordValue common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, wordIndexMin, wordIndexMax, wordValue, confs)

	if len(ret) == 0 {
		panic("no return value specified for LogsDataWordBetween")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, int, common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, wordIndexMin, wordIndexMax, wordValue, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, int, common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, wordIndexMin, wordIndexMax, wordValue, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, int, common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, wordIndexMin, wordIndexMax, wordValue, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsDataWordGreaterThan provides a mock function with given fields: ctx, eventSig, address, wordIndex, wordValueMin, confs
func (_m *LogPoller) LogsDataWordGreaterThan(ctx context.Context, eventSig common.Hash, address common.Address, wordIndex int, wordValueMin common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, wordIndex, wordValueMin, confs)

	if len(ret) == 0 {
		panic("no return value specified for LogsDataWordGreaterThan")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, wordIndex, wordValueMin, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, wordIndex, wordValueMin, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, wordIndex, wordValueMin, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsDataWordRange provides a mock function with given fields: ctx, eventSig, address, wordIndex, wordValueMin, wordValueMax, confs
func (_m *LogPoller) LogsDataWordRange(ctx context.Context, eventSig common.Hash, address common.Address, wordIndex int, wordValueMin common.Hash, wordValueMax common.Hash, confs logpoller.Confirmations) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, eventSig, address, wordIndex, wordValueMin, wordValueMax, confs)

	if len(ret) == 0 {
		panic("no return value specified for LogsDataWordRange")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) ([]logpoller.Log, error)); ok {
		return rf(ctx, eventSig, address, wordIndex, wordValueMin, wordValueMax, confs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) []logpoller.Log); ok {
		r0 = rf(ctx, eventSig, address, wordIndex, wordValueMin, wordValueMax, confs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, common.Address, int, common.Hash, common.Hash, logpoller.Confirmations) error); ok {
		r1 = rf(ctx, eventSig, address, wordIndex, wordValueMin, wordValueMax, confs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsWithSigs provides a mock function with given fields: ctx, start, end, eventSigs, address
func (_m *LogPoller) LogsWithSigs(ctx context.Context, start int64, end int64, eventSigs []common.Hash, address common.Address) ([]logpoller.Log, error) {
	ret := _m.Called(ctx, start, end, eventSigs, address)

	if len(ret) == 0 {
		panic("no return value specified for LogsWithSigs")
	}

	var r0 []logpoller.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, []common.Hash, common.Address) ([]logpoller.Log, error)); ok {
		return rf(ctx, start, end, eventSigs, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, []common.Hash, common.Address) []logpoller.Log); ok {
		r0 = rf(ctx, start, end, eventSigs, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]logpoller.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, []common.Hash, common.Address) error); ok {
		r1 = rf(ctx, start, end, eventSigs, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *LogPoller) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *LogPoller) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterFilter provides a mock function with given fields: ctx, filter
func (_m *LogPoller) RegisterFilter(ctx context.Context, filter logpoller.Filter) error {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, logpoller.Filter) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Replay provides a mock function with given fields: ctx, fromBlock
func (_m *LogPoller) Replay(ctx context.Context, fromBlock int64) error {
	ret := _m.Called(ctx, fromBlock)

	if len(ret) == 0 {
		panic("no return value specified for Replay")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, fromBlock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplayAsync provides a mock function with given fields: fromBlock
func (_m *LogPoller) ReplayAsync(fromBlock int64) {
	_m.Called(fromBlock)
}

// Start provides a mock function with given fields: _a0
func (_m *LogPoller) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnregisterFilter provides a mock function with given fields: ctx, name
func (_m *LogPoller) UnregisterFilter(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for UnregisterFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLogPoller creates a new instance of LogPoller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogPoller(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogPoller {
	mock := &LogPoller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
