// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	crypto "github.com/FAT/common/crypto"
	mock "github.com/stretchr/testify/mock"

	models "github.com/FAT/models"
)

// IndodaxWebhook is an autogenerated mock type for the IndodaxWebhook type
type IndodaxWebhook struct {
	mock.Mock
}

// Order provides a mock function with given fields: payload
func (_m *IndodaxWebhook) Order(payload *crypto.PayloadSHA512) (*models.ResponseOrderIndodax, error) {
	ret := _m.Called(payload)

	var r0 *models.ResponseOrderIndodax
	if rf, ok := ret.Get(0).(func(*crypto.PayloadSHA512) *models.ResponseOrderIndodax); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseOrderIndodax)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*crypto.PayloadSHA512) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
