// Code generated by mockery v2.43.2. DO NOT EDIT.

package ldapauth

import (
	ldap "github.com/go-ldap/ldap/v3"
	mock "github.com/stretchr/testify/mock"
)

// LDAPConn is an autogenerated mock type for the LDAPConn type
type LDAPConn struct {
	mock.Mock
}

type LDAPConn_Expecter struct {
	mock *mock.Mock
}

func (_m *LDAPConn) EXPECT() *LDAPConn_Expecter {
	return &LDAPConn_Expecter{mock: &_m.Mock}
}

// Bind provides a mock function with given fields: username, password
func (_m *LDAPConn) Bind(username string, password string) error {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LDAPConn_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type LDAPConn_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - username string
//   - password string
func (_e *LDAPConn_Expecter) Bind(username interface{}, password interface{}) *LDAPConn_Bind_Call {
	return &LDAPConn_Bind_Call{Call: _e.mock.On("Bind", username, password)}
}

func (_c *LDAPConn_Bind_Call) Run(run func(username string, password string)) *LDAPConn_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *LDAPConn_Bind_Call) Return(_a0 error) *LDAPConn_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LDAPConn_Bind_Call) RunAndReturn(run func(string, string) error) *LDAPConn_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *LDAPConn) Close() error {
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

// LDAPConn_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type LDAPConn_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *LDAPConn_Expecter) Close() *LDAPConn_Close_Call {
	return &LDAPConn_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *LDAPConn_Close_Call) Run(run func()) *LDAPConn_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *LDAPConn_Close_Call) Return(err error) *LDAPConn_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *LDAPConn_Close_Call) RunAndReturn(run func() error) *LDAPConn_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: searchRequest
func (_m *LDAPConn) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	ret := _m.Called(searchRequest)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *ldap.SearchResult
	var r1 error
	if rf, ok := ret.Get(0).(func(*ldap.SearchRequest) (*ldap.SearchResult, error)); ok {
		return rf(searchRequest)
	}
	if rf, ok := ret.Get(0).(func(*ldap.SearchRequest) *ldap.SearchResult); ok {
		r0 = rf(searchRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ldap.SearchResult)
		}
	}

	if rf, ok := ret.Get(1).(func(*ldap.SearchRequest) error); ok {
		r1 = rf(searchRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LDAPConn_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type LDAPConn_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - searchRequest *ldap.SearchRequest
func (_e *LDAPConn_Expecter) Search(searchRequest interface{}) *LDAPConn_Search_Call {
	return &LDAPConn_Search_Call{Call: _e.mock.On("Search", searchRequest)}
}

func (_c *LDAPConn_Search_Call) Run(run func(searchRequest *ldap.SearchRequest)) *LDAPConn_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ldap.SearchRequest))
	})
	return _c
}

func (_c *LDAPConn_Search_Call) Return(_a0 *ldap.SearchResult, _a1 error) *LDAPConn_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LDAPConn_Search_Call) RunAndReturn(run func(*ldap.SearchRequest) (*ldap.SearchResult, error)) *LDAPConn_Search_Call {
	_c.Call.Return(run)
	return _c
}

// NewLDAPConn creates a new instance of LDAPConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLDAPConn(t interface {
	mock.TestingT
	Cleanup(func())
}) *LDAPConn {
	mock := &LDAPConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
