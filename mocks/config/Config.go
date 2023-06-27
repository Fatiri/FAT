// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	models "github.com/Fatiri/fat/models"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// InitConfig provides a mock function with given fields:
func (_m *Config) InitConfig() (*models.Config, error) {
	ret := _m.Called()

	var r0 *models.Config
	if rf, ok := ret.Get(0).(func() *models.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Config)
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
