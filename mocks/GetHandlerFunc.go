// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	home "github.com/ehabterra/flash_api/api/restapi/operations/home"
	middleware "github.com/go-openapi/runtime/middleware"

	mock "github.com/stretchr/testify/mock"
)

// GetHandlerFunc is an autogenerated mock type for the GetHandlerFunc type
type GetHandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *GetHandlerFunc) Execute(_a0 home.GetParams) middleware.Responder {
	ret := _m.Called(_a0)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(home.GetParams) middleware.Responder); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
