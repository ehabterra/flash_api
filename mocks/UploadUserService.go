// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UploadUserService is an autogenerated mock type for the UploadUserService type
type UploadUserService struct {
	mock.Mock
}

// Upload provides a mock function with given fields: id, number, amount
func (_m *UploadUserService) Upload(id string, number string, amount int64) error {
	ret := _m.Called(id, number, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(id, number, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}