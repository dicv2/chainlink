// Code generated by mockery v2.43.2. DO NOT EDIT.

package config

import (
	url "net/url"

	mock "github.com/stretchr/testify/mock"
)

// TelemetryIngressEndpoint is an autogenerated mock type for the TelemetryIngressEndpoint type
type TelemetryIngressEndpoint struct {
	mock.Mock
}

type TelemetryIngressEndpoint_Expecter struct {
	mock *mock.Mock
}

func (_m *TelemetryIngressEndpoint) EXPECT() *TelemetryIngressEndpoint_Expecter {
	return &TelemetryIngressEndpoint_Expecter{mock: &_m.Mock}
}

// ChainID provides a mock function with given fields:
func (_m *TelemetryIngressEndpoint) ChainID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TelemetryIngressEndpoint_ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChainID'
type TelemetryIngressEndpoint_ChainID_Call struct {
	*mock.Call
}

// ChainID is a helper method to define mock.On call
func (_e *TelemetryIngressEndpoint_Expecter) ChainID() *TelemetryIngressEndpoint_ChainID_Call {
	return &TelemetryIngressEndpoint_ChainID_Call{Call: _e.mock.On("ChainID")}
}

func (_c *TelemetryIngressEndpoint_ChainID_Call) Run(run func()) *TelemetryIngressEndpoint_ChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngressEndpoint_ChainID_Call) Return(_a0 string) *TelemetryIngressEndpoint_ChainID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngressEndpoint_ChainID_Call) RunAndReturn(run func() string) *TelemetryIngressEndpoint_ChainID_Call {
	_c.Call.Return(run)
	return _c
}

// Network provides a mock function with given fields:
func (_m *TelemetryIngressEndpoint) Network() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Network")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TelemetryIngressEndpoint_Network_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Network'
type TelemetryIngressEndpoint_Network_Call struct {
	*mock.Call
}

// Network is a helper method to define mock.On call
func (_e *TelemetryIngressEndpoint_Expecter) Network() *TelemetryIngressEndpoint_Network_Call {
	return &TelemetryIngressEndpoint_Network_Call{Call: _e.mock.On("Network")}
}

func (_c *TelemetryIngressEndpoint_Network_Call) Run(run func()) *TelemetryIngressEndpoint_Network_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngressEndpoint_Network_Call) Return(_a0 string) *TelemetryIngressEndpoint_Network_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngressEndpoint_Network_Call) RunAndReturn(run func() string) *TelemetryIngressEndpoint_Network_Call {
	_c.Call.Return(run)
	return _c
}

// ServerPubKey provides a mock function with given fields:
func (_m *TelemetryIngressEndpoint) ServerPubKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ServerPubKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TelemetryIngressEndpoint_ServerPubKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServerPubKey'
type TelemetryIngressEndpoint_ServerPubKey_Call struct {
	*mock.Call
}

// ServerPubKey is a helper method to define mock.On call
func (_e *TelemetryIngressEndpoint_Expecter) ServerPubKey() *TelemetryIngressEndpoint_ServerPubKey_Call {
	return &TelemetryIngressEndpoint_ServerPubKey_Call{Call: _e.mock.On("ServerPubKey")}
}

func (_c *TelemetryIngressEndpoint_ServerPubKey_Call) Run(run func()) *TelemetryIngressEndpoint_ServerPubKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngressEndpoint_ServerPubKey_Call) Return(_a0 string) *TelemetryIngressEndpoint_ServerPubKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngressEndpoint_ServerPubKey_Call) RunAndReturn(run func() string) *TelemetryIngressEndpoint_ServerPubKey_Call {
	_c.Call.Return(run)
	return _c
}

// URL provides a mock function with given fields:
func (_m *TelemetryIngressEndpoint) URL() *url.URL {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for URL")
	}

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

// TelemetryIngressEndpoint_URL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'URL'
type TelemetryIngressEndpoint_URL_Call struct {
	*mock.Call
}

// URL is a helper method to define mock.On call
func (_e *TelemetryIngressEndpoint_Expecter) URL() *TelemetryIngressEndpoint_URL_Call {
	return &TelemetryIngressEndpoint_URL_Call{Call: _e.mock.On("URL")}
}

func (_c *TelemetryIngressEndpoint_URL_Call) Run(run func()) *TelemetryIngressEndpoint_URL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngressEndpoint_URL_Call) Return(_a0 *url.URL) *TelemetryIngressEndpoint_URL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngressEndpoint_URL_Call) RunAndReturn(run func() *url.URL) *TelemetryIngressEndpoint_URL_Call {
	_c.Call.Return(run)
	return _c
}

// NewTelemetryIngressEndpoint creates a new instance of TelemetryIngressEndpoint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTelemetryIngressEndpoint(t interface {
	mock.TestingT
	Cleanup(func())
}) *TelemetryIngressEndpoint {
	mock := &TelemetryIngressEndpoint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
