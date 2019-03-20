// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// PreferenceStore is an autogenerated mock type for the PreferenceStore type
type PreferenceStore struct {
	mock.Mock
}

// CleanupFlagsBatch provides a mock function with given fields: limit
func (_m *PreferenceStore) CleanupFlagsBatch(limit int64) store.StoreChannel {
	ret := _m.Called(limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64) store.StoreChannel); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: userId, category, name
func (_m *PreferenceStore) Delete(userId string, category string, name string) store.StoreChannel {
	ret := _m.Called(userId, category, name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, string) store.StoreChannel); ok {
		r0 = rf(userId, category, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteCategory provides a mock function with given fields: userId, category
func (_m *PreferenceStore) DeleteCategory(userId string, category string) store.StoreChannel {
	ret := _m.Called(userId, category)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteCategoryAndName provides a mock function with given fields: category, name
func (_m *PreferenceStore) DeleteCategoryAndName(category string, name string) store.StoreChannel {
	ret := _m.Called(category, name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(category, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: userId, category, name
func (_m *PreferenceStore) Get(userId string, category string, name string) store.StoreChannel {
	ret := _m.Called(userId, category, name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, string) store.StoreChannel); ok {
		r0 = rf(userId, category, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: userId
func (_m *PreferenceStore) GetAll(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetCategory provides a mock function with given fields: userId, category
func (_m *PreferenceStore) GetCategory(userId string, category string) store.StoreChannel {
	ret := _m.Called(userId, category)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// IsFeatureEnabled provides a mock function with given fields: feature, userId
func (_m *PreferenceStore) IsFeatureEnabled(feature string, userId string) store.StoreChannel {
	ret := _m.Called(feature, userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(feature, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *PreferenceStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: preferences
func (_m *PreferenceStore) Save(preferences *model.Preferences) store.StoreChannel {
	ret := _m.Called(preferences)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Preferences) store.StoreChannel); ok {
		r0 = rf(preferences)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}