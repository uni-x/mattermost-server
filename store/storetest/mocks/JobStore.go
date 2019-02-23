// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// JobStore is an autogenerated mock type for the JobStore type
type JobStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *JobStore) Delete(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *JobStore) Get(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllByStatus provides a mock function with given fields: status
func (_m *JobStore) GetAllByStatus(status string) store.StoreChannel {
	ret := _m.Called(status)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllByType provides a mock function with given fields: jobType
func (_m *JobStore) GetAllByType(jobType string) store.StoreChannel {
	ret := _m.Called(jobType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(jobType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllByTypePage provides a mock function with given fields: jobType, offset, limit
func (_m *JobStore) GetAllByTypePage(jobType string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(jobType, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(jobType, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPage provides a mock function with given fields: offset, limit
func (_m *JobStore) GetAllPage(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetCountByStatusAndType provides a mock function with given fields: status, jobType
func (_m *JobStore) GetCountByStatusAndType(status string, jobType string) store.StoreChannel {
	ret := _m.Called(status, jobType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(status, jobType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetNewestJobByStatusAndType provides a mock function with given fields: status, jobType
func (_m *JobStore) GetNewestJobByStatusAndType(status string, jobType string) store.StoreChannel {
	ret := _m.Called(status, jobType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(status, jobType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: job
func (_m *JobStore) Save(job *model.Job) store.StoreChannel {
	ret := _m.Called(job)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Job) store.StoreChannel); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateOptimistically provides a mock function with given fields: job, currentStatus
func (_m *JobStore) UpdateOptimistically(job *model.Job, currentStatus string) store.StoreChannel {
	ret := _m.Called(job, currentStatus)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Job, string) store.StoreChannel); ok {
		r0 = rf(job, currentStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: id, status
func (_m *JobStore) UpdateStatus(id string, status string) store.StoreChannel {
	ret := _m.Called(id, status)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateStatusOptimistically provides a mock function with given fields: id, currentStatus, newStatus
func (_m *JobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) store.StoreChannel {
	ret := _m.Called(id, currentStatus, newStatus)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, string) store.StoreChannel); ok {
		r0 = rf(id, currentStatus, newStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
