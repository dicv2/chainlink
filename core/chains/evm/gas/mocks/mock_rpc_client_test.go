// Code generated by mockery v2.43.2. DO NOT EDIT.

package gas

import (
	big "math/big"

	assets "github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"

	context "context"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"

	mock "github.com/stretchr/testify/mock"

	rollups "github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/rollups"

	types "github.com/smartcontractkit/chainlink/v2/common/fee/types"
)

// EvmFeeEstimator is an autogenerated mock type for the EvmFeeEstimator type
type EvmFeeEstimator struct {
	mock.Mock
}

type EvmFeeEstimator_Expecter struct {
	mock *mock.Mock
}

func (_m *EvmFeeEstimator) EXPECT() *EvmFeeEstimator_Expecter {
	return &EvmFeeEstimator_Expecter{mock: &_m.Mock}
}

// BumpFee provides a mock function with given fields: ctx, originalFee, feeLimit, maxFeePrice, attempts
func (_m *EvmFeeEstimator) BumpFee(ctx context.Context, originalFee EvmFee, feeLimit uint64, maxFeePrice *assets.Wei, attempts []EvmPriorAttempt) (EvmFee, uint64, error) {
	ret := _m.Called(ctx, originalFee, feeLimit, maxFeePrice, attempts)

	if len(ret) == 0 {
		panic("no return value specified for BumpFee")
	}

	var r0 EvmFee
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, EvmFee, uint64, *assets.Wei, []EvmPriorAttempt) (EvmFee, uint64, error)); ok {
		return rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, EvmFee, uint64, *assets.Wei, []EvmPriorAttempt) EvmFee); ok {
		r0 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r0 = ret.Get(0).(EvmFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, EvmFee, uint64, *assets.Wei, []EvmPriorAttempt) uint64); ok {
		r1 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, EvmFee, uint64, *assets.Wei, []EvmPriorAttempt) error); ok {
		r2 = rf(ctx, originalFee, feeLimit, maxFeePrice, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EvmFeeEstimator_BumpFee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BumpFee'
type EvmFeeEstimator_BumpFee_Call struct {
	*mock.Call
}

// BumpFee is a helper method to define mock.On call
//   - ctx context.Context
//   - originalFee EvmFee
//   - feeLimit uint64
//   - maxFeePrice *assets.Wei
//   - attempts []EvmPriorAttempt
func (_e *EvmFeeEstimator_Expecter) BumpFee(ctx interface{}, originalFee interface{}, feeLimit interface{}, maxFeePrice interface{}, attempts interface{}) *EvmFeeEstimator_BumpFee_Call {
	return &EvmFeeEstimator_BumpFee_Call{Call: _e.mock.On("BumpFee", ctx, originalFee, feeLimit, maxFeePrice, attempts)}
}

func (_c *EvmFeeEstimator_BumpFee_Call) Run(run func(ctx context.Context, originalFee EvmFee, feeLimit uint64, maxFeePrice *assets.Wei, attempts []EvmPriorAttempt)) *EvmFeeEstimator_BumpFee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(EvmFee), args[2].(uint64), args[3].(*assets.Wei), args[4].([]EvmPriorAttempt))
	})
	return _c
}

func (_c *EvmFeeEstimator_BumpFee_Call) Return(bumpedFee EvmFee, chainSpecificFeeLimit uint64, err error) *EvmFeeEstimator_BumpFee_Call {
	_c.Call.Return(bumpedFee, chainSpecificFeeLimit, err)
	return _c
}

func (_c *EvmFeeEstimator_BumpFee_Call) RunAndReturn(run func(context.Context, EvmFee, uint64, *assets.Wei, []EvmPriorAttempt) (EvmFee, uint64, error)) *EvmFeeEstimator_BumpFee_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *EvmFeeEstimator) Close() error {
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

// EvmFeeEstimator_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type EvmFeeEstimator_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *EvmFeeEstimator_Expecter) Close() *EvmFeeEstimator_Close_Call {
	return &EvmFeeEstimator_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *EvmFeeEstimator_Close_Call) Run(run func()) *EvmFeeEstimator_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmFeeEstimator_Close_Call) Return(_a0 error) *EvmFeeEstimator_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_Close_Call) RunAndReturn(run func() error) *EvmFeeEstimator_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetFee provides a mock function with given fields: ctx, calldata, feeLimit, maxFeePrice, opts
func (_m *EvmFeeEstimator) GetFee(ctx context.Context, calldata []byte, feeLimit uint64, maxFeePrice *assets.Wei, opts ...types.Opt) (EvmFee, uint64, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, calldata, feeLimit, maxFeePrice)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetFee")
	}

	var r0 EvmFee
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint64, *assets.Wei, ...types.Opt) (EvmFee, uint64, error)); ok {
		return rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, uint64, *assets.Wei, ...types.Opt) EvmFee); ok {
		r0 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r0 = ret.Get(0).(EvmFee)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, uint64, *assets.Wei, ...types.Opt) uint64); ok {
		r1 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, uint64, *assets.Wei, ...types.Opt) error); ok {
		r2 = rf(ctx, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EvmFeeEstimator_GetFee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFee'
type EvmFeeEstimator_GetFee_Call struct {
	*mock.Call
}

// GetFee is a helper method to define mock.On call
//   - ctx context.Context
//   - calldata []byte
//   - feeLimit uint64
//   - maxFeePrice *assets.Wei
//   - opts ...types.Opt
func (_e *EvmFeeEstimator_Expecter) GetFee(ctx interface{}, calldata interface{}, feeLimit interface{}, maxFeePrice interface{}, opts ...interface{}) *EvmFeeEstimator_GetFee_Call {
	return &EvmFeeEstimator_GetFee_Call{Call: _e.mock.On("GetFee",
		append([]interface{}{ctx, calldata, feeLimit, maxFeePrice}, opts...)...)}
}

func (_c *EvmFeeEstimator_GetFee_Call) Run(run func(ctx context.Context, calldata []byte, feeLimit uint64, maxFeePrice *assets.Wei, opts ...types.Opt)) *EvmFeeEstimator_GetFee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]types.Opt, len(args)-4)
		for i, a := range args[4:] {
			if a != nil {
				variadicArgs[i] = a.(types.Opt)
			}
		}
		run(args[0].(context.Context), args[1].([]byte), args[2].(uint64), args[3].(*assets.Wei), variadicArgs...)
	})
	return _c
}

func (_c *EvmFeeEstimator_GetFee_Call) Return(fee EvmFee, chainSpecificFeeLimit uint64, err error) *EvmFeeEstimator_GetFee_Call {
	_c.Call.Return(fee, chainSpecificFeeLimit, err)
	return _c
}

func (_c *EvmFeeEstimator_GetFee_Call) RunAndReturn(run func(context.Context, []byte, uint64, *assets.Wei, ...types.Opt) (EvmFee, uint64, error)) *EvmFeeEstimator_GetFee_Call {
	_c.Call.Return(run)
	return _c
}

// GetMaxCost provides a mock function with given fields: ctx, amount, calldata, feeLimit, maxFeePrice, opts
func (_m *EvmFeeEstimator) GetMaxCost(ctx context.Context, amount assets.Eth, calldata []byte, feeLimit uint64, maxFeePrice *assets.Wei, opts ...types.Opt) (*big.Int, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, amount, calldata, feeLimit, maxFeePrice)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetMaxCost")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, assets.Eth, []byte, uint64, *assets.Wei, ...types.Opt) (*big.Int, error)); ok {
		return rf(ctx, amount, calldata, feeLimit, maxFeePrice, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, assets.Eth, []byte, uint64, *assets.Wei, ...types.Opt) *big.Int); ok {
		r0 = rf(ctx, amount, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, assets.Eth, []byte, uint64, *assets.Wei, ...types.Opt) error); ok {
		r1 = rf(ctx, amount, calldata, feeLimit, maxFeePrice, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvmFeeEstimator_GetMaxCost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMaxCost'
type EvmFeeEstimator_GetMaxCost_Call struct {
	*mock.Call
}

// GetMaxCost is a helper method to define mock.On call
//   - ctx context.Context
//   - amount assets.Eth
//   - calldata []byte
//   - feeLimit uint64
//   - maxFeePrice *assets.Wei
//   - opts ...types.Opt
func (_e *EvmFeeEstimator_Expecter) GetMaxCost(ctx interface{}, amount interface{}, calldata interface{}, feeLimit interface{}, maxFeePrice interface{}, opts ...interface{}) *EvmFeeEstimator_GetMaxCost_Call {
	return &EvmFeeEstimator_GetMaxCost_Call{Call: _e.mock.On("GetMaxCost",
		append([]interface{}{ctx, amount, calldata, feeLimit, maxFeePrice}, opts...)...)}
}

func (_c *EvmFeeEstimator_GetMaxCost_Call) Run(run func(ctx context.Context, amount assets.Eth, calldata []byte, feeLimit uint64, maxFeePrice *assets.Wei, opts ...types.Opt)) *EvmFeeEstimator_GetMaxCost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]types.Opt, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(types.Opt)
			}
		}
		run(args[0].(context.Context), args[1].(assets.Eth), args[2].([]byte), args[3].(uint64), args[4].(*assets.Wei), variadicArgs...)
	})
	return _c
}

func (_c *EvmFeeEstimator_GetMaxCost_Call) Return(_a0 *big.Int, _a1 error) *EvmFeeEstimator_GetMaxCost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EvmFeeEstimator_GetMaxCost_Call) RunAndReturn(run func(context.Context, assets.Eth, []byte, uint64, *assets.Wei, ...types.Opt) (*big.Int, error)) *EvmFeeEstimator_GetMaxCost_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *EvmFeeEstimator) HealthReport() map[string]error {
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

// EvmFeeEstimator_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type EvmFeeEstimator_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *EvmFeeEstimator_Expecter) HealthReport() *EvmFeeEstimator_HealthReport_Call {
	return &EvmFeeEstimator_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *EvmFeeEstimator_HealthReport_Call) Run(run func()) *EvmFeeEstimator_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmFeeEstimator_HealthReport_Call) Return(_a0 map[string]error) *EvmFeeEstimator_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_HealthReport_Call) RunAndReturn(run func() map[string]error) *EvmFeeEstimator_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// L1Oracle provides a mock function with given fields:
func (_m *EvmFeeEstimator) L1Oracle() rollups.L1Oracle {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for L1Oracle")
	}

	var r0 rollups.L1Oracle
	if rf, ok := ret.Get(0).(func() rollups.L1Oracle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rollups.L1Oracle)
		}
	}

	return r0
}

// EvmFeeEstimator_L1Oracle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'L1Oracle'
type EvmFeeEstimator_L1Oracle_Call struct {
	*mock.Call
}

// L1Oracle is a helper method to define mock.On call
func (_e *EvmFeeEstimator_Expecter) L1Oracle() *EvmFeeEstimator_L1Oracle_Call {
	return &EvmFeeEstimator_L1Oracle_Call{Call: _e.mock.On("L1Oracle")}
}

func (_c *EvmFeeEstimator_L1Oracle_Call) Run(run func()) *EvmFeeEstimator_L1Oracle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmFeeEstimator_L1Oracle_Call) Return(_a0 rollups.L1Oracle) *EvmFeeEstimator_L1Oracle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_L1Oracle_Call) RunAndReturn(run func() rollups.L1Oracle) *EvmFeeEstimator_L1Oracle_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *EvmFeeEstimator) Name() string {
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

// EvmFeeEstimator_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type EvmFeeEstimator_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *EvmFeeEstimator_Expecter) Name() *EvmFeeEstimator_Name_Call {
	return &EvmFeeEstimator_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *EvmFeeEstimator_Name_Call) Run(run func()) *EvmFeeEstimator_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmFeeEstimator_Name_Call) Return(_a0 string) *EvmFeeEstimator_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_Name_Call) RunAndReturn(run func() string) *EvmFeeEstimator_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *EvmFeeEstimator) OnNewLongestChain(ctx context.Context, head *evmtypes.Head) {
	_m.Called(ctx, head)
}

// EvmFeeEstimator_OnNewLongestChain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnNewLongestChain'
type EvmFeeEstimator_OnNewLongestChain_Call struct {
	*mock.Call
}

// OnNewLongestChain is a helper method to define mock.On call
//   - ctx context.Context
//   - head *evmtypes.Head
func (_e *EvmFeeEstimator_Expecter) OnNewLongestChain(ctx interface{}, head interface{}) *EvmFeeEstimator_OnNewLongestChain_Call {
	return &EvmFeeEstimator_OnNewLongestChain_Call{Call: _e.mock.On("OnNewLongestChain", ctx, head)}
}

func (_c *EvmFeeEstimator_OnNewLongestChain_Call) Run(run func(ctx context.Context, head *evmtypes.Head)) *EvmFeeEstimator_OnNewLongestChain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*evmtypes.Head))
	})
	return _c
}

func (_c *EvmFeeEstimator_OnNewLongestChain_Call) Return() *EvmFeeEstimator_OnNewLongestChain_Call {
	_c.Call.Return()
	return _c
}

func (_c *EvmFeeEstimator_OnNewLongestChain_Call) RunAndReturn(run func(context.Context, *evmtypes.Head)) *EvmFeeEstimator_OnNewLongestChain_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *EvmFeeEstimator) Ready() error {
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

// EvmFeeEstimator_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type EvmFeeEstimator_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *EvmFeeEstimator_Expecter) Ready() *EvmFeeEstimator_Ready_Call {
	return &EvmFeeEstimator_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *EvmFeeEstimator_Ready_Call) Run(run func()) *EvmFeeEstimator_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EvmFeeEstimator_Ready_Call) Return(_a0 error) *EvmFeeEstimator_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_Ready_Call) RunAndReturn(run func() error) *EvmFeeEstimator_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *EvmFeeEstimator) Start(_a0 context.Context) error {
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

// EvmFeeEstimator_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type EvmFeeEstimator_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *EvmFeeEstimator_Expecter) Start(_a0 interface{}) *EvmFeeEstimator_Start_Call {
	return &EvmFeeEstimator_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *EvmFeeEstimator_Start_Call) Run(run func(_a0 context.Context)) *EvmFeeEstimator_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *EvmFeeEstimator_Start_Call) Return(_a0 error) *EvmFeeEstimator_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EvmFeeEstimator_Start_Call) RunAndReturn(run func(context.Context) error) *EvmFeeEstimator_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewEvmFeeEstimator creates a new instance of EvmFeeEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvmFeeEstimator(t interface {
	mock.TestingT
	Cleanup(func())
}) *EvmFeeEstimator {
	mock := &EvmFeeEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
