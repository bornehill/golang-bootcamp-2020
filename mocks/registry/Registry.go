// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	controller "github.com/bornehill/golang-bootcamp-2020/interface/controller"
	mock "github.com/stretchr/testify/mock"
)

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

// NewAppController provides a mock function with given fields:
func (_m *Registry) NewAppController() controller.AppController {
	ret := _m.Called()

	var r0 controller.AppController
	if rf, ok := ret.Get(0).(func() controller.AppController); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(controller.AppController)
	}

	return r0
}
