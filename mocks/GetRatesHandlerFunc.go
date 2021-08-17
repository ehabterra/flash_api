// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	middleware "github.com/go-openapi/runtime/middleware"
	mock "github.com/stretchr/testify/mock"

	models "github.com/ehabterra/flash_api/api/models"

	rates "github.com/ehabterra/flash_api/api/restapi/operations/rates"
)

// GetRatesHandlerFunc is an autogenerated mock type for the GetRatesHandlerFunc type
type GetRatesHandlerFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *GetRatesHandlerFunc) Execute(_a0 rates.GetRatesParams, _a1 *models.Principle) middleware.Responder {
	ret := _m.Called(_a0, _a1)

	var r0 middleware.Responder
	if rf, ok := ret.Get(0).(func(rates.GetRatesParams, *models.Principle) middleware.Responder); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(middleware.Responder)
		}
	}

	return r0
}
