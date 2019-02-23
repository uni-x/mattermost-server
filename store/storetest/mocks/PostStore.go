// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// PostStore is an autogenerated mock type for the PostStore type
type PostStore struct {
	mock.Mock
}

// AnalyticsPostCount provides a mock function with given fields: teamId, mustHaveFile, mustHaveHashtag
func (_m *PostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) store.StoreChannel {
	ret := _m.Called(teamId, mustHaveFile, mustHaveHashtag)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool, bool) store.StoreChannel); ok {
		r0 = rf(teamId, mustHaveFile, mustHaveHashtag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// AnalyticsPostCountsByDay provides a mock function with given fields: teamId
func (_m *PostStore) AnalyticsPostCountsByDay(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// AnalyticsUserCountsWithPostsByDay provides a mock function with given fields: teamId
func (_m *PostStore) AnalyticsUserCountsWithPostsByDay(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *PostStore) ClearCaches() {
	_m.Called()
}

// Delete provides a mock function with given fields: postId, time, deleteByID
func (_m *PostStore) Delete(postId string, time int64, deleteByID string) store.StoreChannel {
	ret := _m.Called(postId, time, deleteByID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64, string) store.StoreChannel); ok {
		r0 = rf(postId, time, deleteByID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *PostStore) Get(id string) store.StoreChannel {
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

// GetEtag provides a mock function with given fields: channelId, allowFromCache
func (_m *PostStore) GetEtag(channelId string, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(channelId, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(channelId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetFlaggedPosts provides a mock function with given fields: userId, offset, limit
func (_m *PostStore) GetFlaggedPosts(userId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(userId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetFlaggedPostsForChannel provides a mock function with given fields: userId, channelId, offset, limit
func (_m *PostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(userId, channelId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, channelId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetFlaggedPostsForTeam provides a mock function with given fields: userId, teamId, offset, limit
func (_m *PostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(userId, teamId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, teamId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetMaxPostSize provides a mock function with given fields:
func (_m *PostStore) GetMaxPostSize() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetOldest provides a mock function with given fields:
func (_m *PostStore) GetOldest() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetParentsForExportAfter provides a mock function with given fields: limit, afterId
func (_m *PostStore) GetParentsForExportAfter(limit int, afterId string) store.StoreChannel {
	ret := _m.Called(limit, afterId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, string) store.StoreChannel); ok {
		r0 = rf(limit, afterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPosts provides a mock function with given fields: channelId, offset, limit, allowFromCache
func (_m *PostStore) GetPosts(channelId string, offset int, limit int, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(channelId, offset, limit, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int, bool) store.StoreChannel); ok {
		r0 = rf(channelId, offset, limit, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPostsAfter provides a mock function with given fields: channelId, postId, numPosts, offset
func (_m *PostStore) GetPostsAfter(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	ret := _m.Called(channelId, postId, numPosts, offset)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int, int) store.StoreChannel); ok {
		r0 = rf(channelId, postId, numPosts, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPostsBatchForIndexing provides a mock function with given fields: startTime, endTime, limit
func (_m *PostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	ret := _m.Called(startTime, endTime, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int64, int) store.StoreChannel); ok {
		r0 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPostsBefore provides a mock function with given fields: channelId, postId, numPosts, offset
func (_m *PostStore) GetPostsBefore(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	ret := _m.Called(channelId, postId, numPosts, offset)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int, int) store.StoreChannel); ok {
		r0 = rf(channelId, postId, numPosts, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPostsByIds provides a mock function with given fields: postIds
func (_m *PostStore) GetPostsByIds(postIds []string) store.StoreChannel {
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

// GetPostsCreatedAt provides a mock function with given fields: channelId, time
func (_m *PostStore) GetPostsCreatedAt(channelId string, time int64) store.StoreChannel {
	ret := _m.Called(channelId, time)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(channelId, time)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPostsSince provides a mock function with given fields: channelId, time, allowFromCache
func (_m *PostStore) GetPostsSince(channelId string, time int64, allowFromCache bool) store.StoreChannel {
	ret := _m.Called(channelId, time, allowFromCache)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64, bool) store.StoreChannel); ok {
		r0 = rf(channelId, time, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetRepliesForExport provides a mock function with given fields: parentId
func (_m *PostStore) GetRepliesForExport(parentId string) store.StoreChannel {
	ret := _m.Called(parentId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(parentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetSingle provides a mock function with given fields: id
func (_m *PostStore) GetSingle(id string) store.StoreChannel {
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

// InvalidateLastPostTimeCache provides a mock function with given fields: channelId
func (_m *PostStore) InvalidateLastPostTimeCache(channelId string) {
	_m.Called(channelId)
}

// Overwrite provides a mock function with given fields: post
func (_m *PostStore) Overwrite(post *model.Post) store.StoreChannel {
	ret := _m.Called(post)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Post) store.StoreChannel); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *PostStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
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

// PermanentDeleteByChannel provides a mock function with given fields: channelId
func (_m *PostStore) PermanentDeleteByChannel(channelId string) store.StoreChannel {
	ret := _m.Called(channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *PostStore) PermanentDeleteByUser(userId string) store.StoreChannel {
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

// Save provides a mock function with given fields: post
func (_m *PostStore) Save(post *model.Post) store.StoreChannel {
	ret := _m.Called(post)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Post) store.StoreChannel); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Search provides a mock function with given fields: teamId, userId, params
func (_m *PostStore) Search(teamId string, userId string, params *model.SearchParams) store.StoreChannel {
	ret := _m.Called(teamId, userId, params)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, *model.SearchParams) store.StoreChannel); ok {
		r0 = rf(teamId, userId, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: newPost, oldPost
func (_m *PostStore) Update(newPost *model.Post, oldPost *model.Post) store.StoreChannel {
	ret := _m.Called(newPost, oldPost)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Post, *model.Post) store.StoreChannel); ok {
		r0 = rf(newPost, oldPost)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
