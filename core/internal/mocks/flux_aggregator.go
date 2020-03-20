// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	assets "chainlink/core/assets"

	abi "github.com/ethereum/go-ethereum/accounts/abi"

	common "github.com/ethereum/go-ethereum/common"

	contracts "chainlink/core/services/eth/contracts"

	coreeth "chainlink/core/eth"

	eth "chainlink/core/services/eth"

	mock "github.com/stretchr/testify/mock"
)

// FluxAggregator is an autogenerated mock type for the FluxAggregator type
type FluxAggregator struct {
	mock.Mock
}

// ABI provides a mock function with given fields:
func (_m *FluxAggregator) ABI() *abi.ABI {
	ret := _m.Called()

	var r0 *abi.ABI
	if rf, ok := ret.Get(0).(func() *abi.ABI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abi.ABI)
		}
	}

	return r0
}

// Call provides a mock function with given fields: result, methodName, args
func (_m *FluxAggregator) Call(result interface{}, methodName string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, methodName)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, methodName, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EncodeMessageCall provides a mock function with given fields: method, args
func (_m *FluxAggregator) EncodeMessageCall(method string, args ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, ...interface{}) []byte); ok {
		r0 = rf(method, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(method, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAvailableFunds provides a mock function with given fields:
func (_m *FluxAggregator) GetAvailableFunds() (*assets.Link, error) {
	ret := _m.Called()

	var r0 *assets.Link
	if rf, ok := ret.Get(0).(func() *assets.Link); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMethodID provides a mock function with given fields: method
func (_m *FluxAggregator) GetMethodID(method string) ([]byte, error) {
	ret := _m.Called(method)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RoundState provides a mock function with given fields: oracle
func (_m *FluxAggregator) RoundState(oracle common.Address) (contracts.FluxAggregatorRoundState, error) {
	ret := _m.Called(oracle)

	var r0 contracts.FluxAggregatorRoundState
	if rf, ok := ret.Get(0).(func(common.Address) contracts.FluxAggregatorRoundState); ok {
		r0 = rf(oracle)
	} else {
		r0 = ret.Get(0).(contracts.FluxAggregatorRoundState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(oracle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToLogs provides a mock function with given fields: listener
func (_m *FluxAggregator) SubscribeToLogs(listener eth.LogListener) (bool, eth.UnsubscribeFunc) {
	ret := _m.Called(listener)

	var r0 bool
	if rf, ok := ret.Get(0).(func(eth.LogListener) bool); ok {
		r0 = rf(listener)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 eth.UnsubscribeFunc
	if rf, ok := ret.Get(1).(func(eth.LogListener) eth.UnsubscribeFunc); ok {
		r1 = rf(listener)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(eth.UnsubscribeFunc)
		}
	}

	return r0, r1
}

// UnpackLog provides a mock function with given fields: out, event, log
func (_m *FluxAggregator) UnpackLog(out interface{}, event string, log coreeth.Log) error {
	ret := _m.Called(out, event, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, coreeth.Log) error); ok {
		r0 = rf(out, event, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
