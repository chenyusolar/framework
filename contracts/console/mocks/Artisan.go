// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	console "github.com/chenyusolar/framework/contracts/console"
	mock "github.com/stretchr/testify/mock"
)

// Artisan is an autogenerated mock type for the Artisan type
type Artisan struct {
	mock.Mock
}

// Call provides a mock function with given fields: command
func (_m *Artisan) Call(command string) {
	_m.Called(command)
}

// CallAndExit provides a mock function with given fields: command
func (_m *Artisan) CallAndExit(command string) {
	_m.Called(command)
}

// Register provides a mock function with given fields: commands
func (_m *Artisan) Register(commands []console.Command) {
	_m.Called(commands)
}

// Run provides a mock function with given fields: args, exitIfArtisan
func (_m *Artisan) Run(args []string, exitIfArtisan bool) {
	_m.Called(args, exitIfArtisan)
}

type mockConstructorTestingTNewArtisan interface {
	mock.TestingT
	Cleanup(func())
}

// NewArtisan creates a new instance of Artisan. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArtisan(t mockConstructorTestingTNewArtisan) *Artisan {
	mock := &Artisan{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
