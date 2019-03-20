// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// ReactionStore is an autogenerated mock type for the ReactionStore type
type ReactionStore struct {
	mock.Mock
}

// BulkGetForPosts provides a mock function with given fields: postIds
func (_m *ReactionStore) BulkGetForPosts(postIds []string) store.StoreChannel {
	ret := _m.Called(postIds)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func([]string) store.StoreChannel); ok {
		r0 = rf(postIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: reaction
func (_m *ReactionStore) Delete(reaction *model.Reaction) store.StoreChannel {
	ret := _m.Called(reaction)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Reaction) store.StoreChannel); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteAllWithEmojiName provides a mock function with given fields: emojiName
func (_m *ReactionStore) DeleteAllWithEmojiName(emojiName string) store.StoreChannel {
	ret := _m.Called(emojiName)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(emojiName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetForPost provides a mock function with given fields: postId, allowFromCache
func (_m *ReactionStore) GetForPost(postId string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(postId, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(postId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *ReactionStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	ret := _m.Called(endTime, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int64) store.StoreChannel); ok {
		r0 = rf(endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: reaction
func (_m *ReactionStore) Save(reaction *model.Reaction) store.StoreChannel {
	ret := _m.Called(reaction)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Reaction) store.StoreChannel); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}