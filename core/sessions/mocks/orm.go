// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	auth "github.com/smartcontractkit/chainlink/core/auth"
	bridges "github.com/smartcontractkit/chainlink/core/bridges"

	mock "github.com/stretchr/testify/mock"

	sessions "github.com/smartcontractkit/chainlink/core/sessions"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// AuthorizedUserWithSession provides a mock function with given fields: sessionID
func (_m *ORM) AuthorizedUserWithSession(sessionID string) (sessions.User, error) {
	ret := _m.Called(sessionID)

	var r0 sessions.User
	if rf, ok := ret.Get(0).(func(string) sessions.User); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearNonCurrentSessions provides a mock function with given fields: sessionID
func (_m *ORM) ClearNonCurrentSessions(sessionID string) error {
	ret := _m.Called(sessionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAndSetAuthToken provides a mock function with given fields: user
func (_m *ORM) CreateAndSetAuthToken(user *sessions.User) (*auth.Token, error) {
	ret := _m.Called(user)

	var r0 *auth.Token
	if rf, ok := ret.Get(0).(func(*sessions.User) *auth.Token); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sessions.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSession provides a mock function with given fields: sr
func (_m *ORM) CreateSession(sr sessions.SessionRequest) (string, error) {
	ret := _m.Called(sr)

	var r0 string
	if rf, ok := ret.Get(0).(func(sessions.SessionRequest) string); ok {
		r0 = rf(sr)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(sessions.SessionRequest) error); ok {
		r1 = rf(sr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: user
func (_m *ORM) CreateUser(user *sessions.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAuthToken provides a mock function with given fields: user
func (_m *ORM) DeleteAuthToken(user *sessions.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields:
func (_m *ORM) DeleteUser() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserSession provides a mock function with given fields: sessionID
func (_m *ORM) DeleteUserSession(sessionID string) error {
	ret := _m.Called(sessionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindExternalInitiator provides a mock function with given fields: eia
func (_m *ORM) FindExternalInitiator(eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(eia)

	var r0 *bridges.ExternalInitiator
	if rf, ok := ret.Get(0).(func(*auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auth.Token) error); ok {
		r1 = rf(eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields:
func (_m *ORM) FindUser() (sessions.User, error) {
	ret := _m.Called()

	var r0 sessions.User
	if rf, ok := ret.Get(0).(func() sessions.User); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserWebAuthn provides a mock function with given fields: email
func (_m *ORM) GetUserWebAuthn(email string) ([]sessions.WebAuthn, error) {
	ret := _m.Called(email)

	var r0 []sessions.WebAuthn
	if rf, ok := ret.Get(0).(func(string) []sessions.WebAuthn); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.WebAuthn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveWebAuthn provides a mock function with given fields: token
func (_m *ORM) SaveWebAuthn(token *sessions.WebAuthn) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.WebAuthn) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sessions provides a mock function with given fields: offset, limit
func (_m *ORM) Sessions(offset int, limit int) ([]sessions.Session, error) {
	ret := _m.Called(offset, limit)

	var r0 []sessions.Session
	if rf, ok := ret.Get(0).(func(int, int) []sessions.Session); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAuthToken provides a mock function with given fields: user, token
func (_m *ORM) SetAuthToken(user *sessions.User, token *auth.Token) error {
	ret := _m.Called(user, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User, *auth.Token) error); ok {
		r0 = rf(user, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPassword provides a mock function with given fields: user, newPassword
func (_m *ORM) SetPassword(user *sessions.User, newPassword string) error {
	ret := _m.Called(user, newPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sessions.User, string) error); ok {
		r0 = rf(user, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
