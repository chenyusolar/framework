// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	http "github.com/chenyusolar/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Context is an autogenerated mock type for the Context type
type Context struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *Context) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Deadline provides a mock function with given fields:
func (_m *Context) Deadline() (time.Time, bool) {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Done provides a mock function with given fields:
func (_m *Context) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Err provides a mock function with given fields:
func (_m *Context) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Request provides a mock function with given fields:
func (_m *Context) Request() http.Request {
	ret := _m.Called()

	var r0 http.Request
	if rf, ok := ret.Get(0).(func() http.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Request)
		}
	}

	return r0
}

// Response provides a mock function with given fields:
func (_m *Context) Response() http.Response {
	ret := _m.Called()

	var r0 http.Response
	if rf, ok := ret.Get(0).(func() http.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// Value provides a mock function with given fields: key
func (_m *Context) Value(key interface{}) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// WithValue provides a mock function with given fields: key, value
func (_m *Context) WithValue(key string, value interface{}) {
	_m.Called(key, value)
}

type mockConstructorTestingTNewContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewContext creates a new instance of Context. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContext(t mockConstructorTestingTNewContext) *Context {
	mock := &Context{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
