// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	terra "github.com/smartcontractkit/chainlink-terra/pkg/terra"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// TxManager is an autogenerated mock type for the TxManager type
type TxManager struct {
	mock.Mock
}

// Enqueue provides a mock function with given fields: contractID, msg
func (_m *TxManager) Enqueue(contractID string, msg types.Msg) (int64, error) {
	ret := _m.Called(contractID, msg)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, types.Msg) int64); ok {
		r0 = rf(contractID, msg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.Msg) error); ok {
		r1 = rf(contractID, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GasPrice provides a mock function with given fields:
func (_m *TxManager) GasPrice() (types.DecCoin, error) {
	ret := _m.Called()

	var r0 types.DecCoin
	if rf, ok := ret.Get(0).(func() types.DecCoin); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.DecCoin)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMsgs provides a mock function with given fields: ids
func (_m *TxManager) GetMsgs(ids ...int64) (terra.Msgs, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 terra.Msgs
	if rf, ok := ret.Get(0).(func(...int64) terra.Msgs); ok {
		r0 = rf(ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(terra.Msgs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...int64) error); ok {
		r1 = rf(ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTxManager creates a new instance of TxManager. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxManager(t testing.TB) *TxManager {
	mock := &TxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
