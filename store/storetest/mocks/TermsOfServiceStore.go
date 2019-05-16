// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// TermsOfServiceStore is an autogenerated mock type for the TermsOfServiceStore type
type TermsOfServiceStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: id, allowFromCache
func (_m *TermsOfServiceStore) Get(id string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(id, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(id, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetLatest provides a mock function with given fields: allowFromCache
func (_m *TermsOfServiceStore) GetLatest(allowFromCache bool) store.StoreChannel {
	ret := _m.Called(allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(bool) store.StoreChannel); ok {
		r0 = rf(allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: termsOfService
func (_m *TermsOfServiceStore) Save(termsOfService *model.TermsOfService) store.StoreChannel {
	ret := _m.Called(termsOfService)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.TermsOfService) store.StoreChannel); ok {
		r0 = rf(termsOfService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
