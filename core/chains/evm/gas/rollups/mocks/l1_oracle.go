// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	big "math/big"

	assets "github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"

	chaintype "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/chaintype"

	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// L1Oracle is an autogenerated mock type for the L1Oracle type
type L1Oracle struct {
	mock.Mock
}

type L1Oracle_Expecter struct {
	mock *mock.Mock
}

func (_m *L1Oracle) EXPECT() *L1Oracle_Expecter {
	return &L1Oracle_Expecter{mock: &_m.Mock}
}

// ChainType provides a mock function with given fields: ctx
func (_m *L1Oracle) ChainType(ctx context.Context) chaintype.ChainType {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ChainType")
	}

	var r0 chaintype.ChainType
	if rf, ok := ret.Get(0).(func(context.Context) chaintype.ChainType); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(chaintype.ChainType)
	}

	return r0
}

// L1Oracle_ChainType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChainType'
type L1Oracle_ChainType_Call struct {
	*mock.Call
}

// ChainType is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L1Oracle_Expecter) ChainType(ctx interface{}) *L1Oracle_ChainType_Call {
	return &L1Oracle_ChainType_Call{Call: _e.mock.On("ChainType", ctx)}
}

func (_c *L1Oracle_ChainType_Call) Run(run func(ctx context.Context)) *L1Oracle_ChainType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1Oracle_ChainType_Call) Return(_a0 chaintype.ChainType) *L1Oracle_ChainType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_ChainType_Call) RunAndReturn(run func(context.Context) chaintype.ChainType) *L1Oracle_ChainType_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *L1Oracle) Close() error {
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

// L1Oracle_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type L1Oracle_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *L1Oracle_Expecter) Close() *L1Oracle_Close_Call {
	return &L1Oracle_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *L1Oracle_Close_Call) Run(run func()) *L1Oracle_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *L1Oracle_Close_Call) Return(_a0 error) *L1Oracle_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_Close_Call) RunAndReturn(run func() error) *L1Oracle_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GasPrice provides a mock function with given fields: ctx
func (_m *L1Oracle) GasPrice(ctx context.Context) (*assets.Wei, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GasPrice")
	}

	var r0 *assets.Wei
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*assets.Wei, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *assets.Wei); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L1Oracle_GasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GasPrice'
type L1Oracle_GasPrice_Call struct {
	*mock.Call
}

// GasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L1Oracle_Expecter) GasPrice(ctx interface{}) *L1Oracle_GasPrice_Call {
	return &L1Oracle_GasPrice_Call{Call: _e.mock.On("GasPrice", ctx)}
}

func (_c *L1Oracle_GasPrice_Call) Run(run func(ctx context.Context)) *L1Oracle_GasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1Oracle_GasPrice_Call) Return(_a0 *assets.Wei, _a1 error) *L1Oracle_GasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L1Oracle_GasPrice_Call) RunAndReturn(run func(context.Context) (*assets.Wei, error)) *L1Oracle_GasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// GetGasCost provides a mock function with given fields: ctx, tx, blockNum
func (_m *L1Oracle) GetGasCost(ctx context.Context, tx *types.Transaction, blockNum *big.Int) (*assets.Wei, error) {
	ret := _m.Called(ctx, tx, blockNum)

	if len(ret) == 0 {
		panic("no return value specified for GetGasCost")
	}

	var r0 *assets.Wei
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, *big.Int) (*assets.Wei, error)); ok {
		return rf(ctx, tx, blockNum)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, *big.Int) *assets.Wei); ok {
		r0 = rf(ctx, tx, blockNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction, *big.Int) error); ok {
		r1 = rf(ctx, tx, blockNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L1Oracle_GetGasCost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGasCost'
type L1Oracle_GetGasCost_Call struct {
	*mock.Call
}

// GetGasCost is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
//   - blockNum *big.Int
func (_e *L1Oracle_Expecter) GetGasCost(ctx interface{}, tx interface{}, blockNum interface{}) *L1Oracle_GetGasCost_Call {
	return &L1Oracle_GetGasCost_Call{Call: _e.mock.On("GetGasCost", ctx, tx, blockNum)}
}

func (_c *L1Oracle_GetGasCost_Call) Run(run func(ctx context.Context, tx *types.Transaction, blockNum *big.Int)) *L1Oracle_GetGasCost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction), args[2].(*big.Int))
	})
	return _c
}

func (_c *L1Oracle_GetGasCost_Call) Return(_a0 *assets.Wei, _a1 error) *L1Oracle_GetGasCost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L1Oracle_GetGasCost_Call) RunAndReturn(run func(context.Context, *types.Transaction, *big.Int) (*assets.Wei, error)) *L1Oracle_GetGasCost_Call {
	_c.Call.Return(run)
	return _c
}

// HealthReport provides a mock function with given fields:
func (_m *L1Oracle) HealthReport() map[string]error {
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

// L1Oracle_HealthReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthReport'
type L1Oracle_HealthReport_Call struct {
	*mock.Call
}

// HealthReport is a helper method to define mock.On call
func (_e *L1Oracle_Expecter) HealthReport() *L1Oracle_HealthReport_Call {
	return &L1Oracle_HealthReport_Call{Call: _e.mock.On("HealthReport")}
}

func (_c *L1Oracle_HealthReport_Call) Run(run func()) *L1Oracle_HealthReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *L1Oracle_HealthReport_Call) Return(_a0 map[string]error) *L1Oracle_HealthReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_HealthReport_Call) RunAndReturn(run func() map[string]error) *L1Oracle_HealthReport_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *L1Oracle) Name() string {
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

// L1Oracle_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type L1Oracle_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *L1Oracle_Expecter) Name() *L1Oracle_Name_Call {
	return &L1Oracle_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *L1Oracle_Name_Call) Run(run func()) *L1Oracle_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *L1Oracle_Name_Call) Return(_a0 string) *L1Oracle_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_Name_Call) RunAndReturn(run func() string) *L1Oracle_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with given fields:
func (_m *L1Oracle) Ready() error {
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

// L1Oracle_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type L1Oracle_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *L1Oracle_Expecter) Ready() *L1Oracle_Ready_Call {
	return &L1Oracle_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *L1Oracle_Ready_Call) Run(run func()) *L1Oracle_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *L1Oracle_Ready_Call) Return(_a0 error) *L1Oracle_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_Ready_Call) RunAndReturn(run func() error) *L1Oracle_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *L1Oracle) Start(_a0 context.Context) error {
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

// L1Oracle_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type L1Oracle_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *L1Oracle_Expecter) Start(_a0 interface{}) *L1Oracle_Start_Call {
	return &L1Oracle_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *L1Oracle_Start_Call) Run(run func(_a0 context.Context)) *L1Oracle_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L1Oracle_Start_Call) Return(_a0 error) *L1Oracle_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L1Oracle_Start_Call) RunAndReturn(run func(context.Context) error) *L1Oracle_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewL1Oracle creates a new instance of L1Oracle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewL1Oracle(t interface {
	mock.TestingT
	Cleanup(func())
}) *L1Oracle {
	mock := &L1Oracle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
