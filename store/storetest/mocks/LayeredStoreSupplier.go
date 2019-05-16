// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/uni-x/mattermost-server/model"
import store "github.com/uni-x/mattermost-server/store"

// LayeredStoreSupplier is an autogenerated mock type for the LayeredStoreSupplier type
type LayeredStoreSupplier struct {
	mock.Mock
}

// ChannelMembersToAdd provides a mock function with given fields: ctx, since, hints
func (_m *LayeredStoreSupplier) ChannelMembersToAdd(ctx context.Context, since int64, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, since)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, since, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ChannelMembersToRemove provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) ChannelMembersToRemove(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// CountGroupsByTeam provides a mock function with given fields: ctx, teamId, opts, hints
func (_m *LayeredStoreSupplier) CountGroupsByTeam(ctx context.Context, teamId string, opts model.GroupSearchOpts, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, teamId, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GroupSearchOpts, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, teamId, opts, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GetGroups provides a mock function with given fields: ctx, page, perPage, opts, hints
func (_m *LayeredStoreSupplier) GetGroups(ctx context.Context, page int, perPage int, opts model.GroupSearchOpts, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, page, perPage, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int, int, model.GroupSearchOpts, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, page, perPage, opts, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GetGroupsByChannel provides a mock function with given fields: ctx, channelId, page, perPage, hints
func (_m *LayeredStoreSupplier) GetGroupsByChannel(ctx context.Context, channelId string, page int, perPage int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, channelId, page, perPage)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, channelId, page, perPage, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GetGroupsByTeam provides a mock function with given fields: ctx, teamId, opts, hints
func (_m *LayeredStoreSupplier) GetGroupsByTeam(ctx context.Context, teamId string, opts model.GroupSearchOpts, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, teamId, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GroupSearchOpts, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, teamId, opts, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupCreate provides a mock function with given fields: ctx, group, hints
func (_m *LayeredStoreSupplier) GroupCreate(ctx context.Context, group *model.Group, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, group)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Group, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, group, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupCreateGroupSyncable provides a mock function with given fields: ctx, groupSyncable, hints
func (_m *LayeredStoreSupplier) GroupCreateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupSyncable)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.GroupSyncable, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupSyncable, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupCreateOrRestoreMember provides a mock function with given fields: ctx, groupID, userID, hints
func (_m *LayeredStoreSupplier) GroupCreateOrRestoreMember(ctx context.Context, groupID string, userID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, userID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDelete provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreSupplier) GroupDelete(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDeleteGroupSyncable provides a mock function with given fields: ctx, groupID, syncableID, syncableType, hints
func (_m *LayeredStoreSupplier) GroupDeleteGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableID, syncableType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.GroupSyncableType, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableID, syncableType, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDeleteMember provides a mock function with given fields: ctx, groupID, userID, hints
func (_m *LayeredStoreSupplier) GroupDeleteMember(ctx context.Context, groupID string, userID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, userID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGet provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreSupplier) GroupGet(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetAllBySource provides a mock function with given fields: ctx, groupSource, hints
func (_m *LayeredStoreSupplier) GroupGetAllBySource(ctx context.Context, groupSource model.GroupSource, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupSource)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, model.GroupSource, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupSource, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetAllGroupSyncablesByGroup provides a mock function with given fields: ctx, groupID, syncableType, hints
func (_m *LayeredStoreSupplier) GroupGetAllGroupSyncablesByGroup(ctx context.Context, groupID string, syncableType model.GroupSyncableType, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GroupSyncableType, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableType, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetByRemoteID provides a mock function with given fields: ctx, remoteID, groupSource, hints
func (_m *LayeredStoreSupplier) GroupGetByRemoteID(ctx context.Context, remoteID string, groupSource model.GroupSource, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, remoteID, groupSource)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GroupSource, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, remoteID, groupSource, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetGroupSyncable provides a mock function with given fields: ctx, groupID, syncableID, syncableType, hints
func (_m *LayeredStoreSupplier) GroupGetGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableID, syncableType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.GroupSyncableType, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableID, syncableType, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetMemberCount provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreSupplier) GroupGetMemberCount(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetMemberUsers provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreSupplier) GroupGetMemberUsers(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetMemberUsersPage provides a mock function with given fields: ctx, groupID, offset, limit, hints
func (_m *LayeredStoreSupplier) GroupGetMemberUsersPage(ctx context.Context, groupID string, offset int, limit int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, offset, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupUpdate provides a mock function with given fields: ctx, group, hints
func (_m *LayeredStoreSupplier) GroupUpdate(ctx context.Context, group *model.Group, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, group)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Group, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, group, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupUpdateGroupSyncable provides a mock function with given fields: ctx, groupSyncable, hints
func (_m *LayeredStoreSupplier) GroupUpdateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupSyncable)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.GroupSyncable, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupSyncable, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *LayeredStoreSupplier) Next() store.LayeredStoreSupplier {
	ret := _m.Called()

	var r0 store.LayeredStoreSupplier
	if rf, ok := ret.Get(0).(func() store.LayeredStoreSupplier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LayeredStoreSupplier)
		}
	}

	return r0
}

// ReactionDelete provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreSupplier) ReactionDelete(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) (*model.Reaction, *model.AppError) {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Reaction
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *model.Reaction); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *model.AppError); ok {
		r1 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ReactionDeleteAllWithEmojiName provides a mock function with given fields: ctx, emojiName, hints
func (_m *LayeredStoreSupplier) ReactionDeleteAllWithEmojiName(ctx context.Context, emojiName string, hints ...store.LayeredStoreHint) *model.AppError {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, emojiName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *model.AppError); ok {
		r0 = rf(ctx, emojiName, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// ReactionGetForPost provides a mock function with given fields: ctx, postId, hints
func (_m *LayeredStoreSupplier) ReactionGetForPost(ctx context.Context, postId string, hints ...store.LayeredStoreHint) ([]*model.Reaction, *model.AppError) {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, postId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*model.Reaction
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) []*model.Reaction); ok {
		r0 = rf(ctx, postId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(context.Context, string, ...store.LayeredStoreHint) *model.AppError); ok {
		r1 = rf(ctx, postId, hints...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ReactionPermanentDeleteBatch provides a mock function with given fields: ctx, endTime, limit, hints
func (_m *LayeredStoreSupplier) ReactionPermanentDeleteBatch(ctx context.Context, endTime int64, limit int64, hints ...store.LayeredStoreHint) (int64, *model.AppError) {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, endTime, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, ...store.LayeredStoreHint) int64); ok {
		r0 = rf(ctx, endTime, limit, hints...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, ...store.LayeredStoreHint) *model.AppError); ok {
		r1 = rf(ctx, endTime, limit, hints...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ReactionSave provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreSupplier) ReactionSave(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) (*model.Reaction, *model.AppError) {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.Reaction
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *model.Reaction); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *model.AppError); ok {
		r1 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ReactionsBulkGetForPosts provides a mock function with given fields: ctx, postIds, hints
func (_m *LayeredStoreSupplier) ReactionsBulkGetForPosts(ctx context.Context, postIds []string, hints ...store.LayeredStoreHint) ([]*model.Reaction, *model.AppError) {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, postIds)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*model.Reaction
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...store.LayeredStoreHint) []*model.Reaction); ok {
		r0 = rf(ctx, postIds, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(context.Context, []string, ...store.LayeredStoreHint) *model.AppError); ok {
		r1 = rf(ctx, postIds, hints...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// RoleDelete provides a mock function with given fields: ctx, roldId, hints
func (_m *LayeredStoreSupplier) RoleDelete(ctx context.Context, roldId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, roldId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, roldId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGet provides a mock function with given fields: ctx, roleId, hints
func (_m *LayeredStoreSupplier) RoleGet(ctx context.Context, roleId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, roleId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, roleId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) RoleGetAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByName provides a mock function with given fields: ctx, name, hints
func (_m *LayeredStoreSupplier) RoleGetByName(ctx context.Context, name string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, name, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByNames provides a mock function with given fields: ctx, names, hints
func (_m *LayeredStoreSupplier) RoleGetByNames(ctx context.Context, names []string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, names)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, names, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RolePermanentDeleteAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) RolePermanentDeleteAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleSave provides a mock function with given fields: ctx, role, hints
func (_m *LayeredStoreSupplier) RoleSave(ctx context.Context, role *model.Role, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Role, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, role, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeDelete provides a mock function with given fields: ctx, schemeId, hints
func (_m *LayeredStoreSupplier) SchemeDelete(ctx context.Context, schemeId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGet provides a mock function with given fields: ctx, schemeId, hints
func (_m *LayeredStoreSupplier) SchemeGet(ctx context.Context, schemeId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGetAllPage provides a mock function with given fields: ctx, scope, offset, limit, hints
func (_m *LayeredStoreSupplier) SchemeGetAllPage(ctx context.Context, scope string, offset int, limit int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, scope, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, scope, offset, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGetByName provides a mock function with given fields: ctx, schemeName, hints
func (_m *LayeredStoreSupplier) SchemeGetByName(ctx context.Context, schemeName string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeName, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemePermanentDeleteAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) SchemePermanentDeleteAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeSave provides a mock function with given fields: ctx, scheme, hints
func (_m *LayeredStoreSupplier) SchemeSave(ctx context.Context, scheme *model.Scheme, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, scheme)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Scheme, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, scheme, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SetChainNext provides a mock function with given fields: _a0
func (_m *LayeredStoreSupplier) SetChainNext(_a0 store.LayeredStoreSupplier) {
	_m.Called(_a0)
}

// TeamMembersToAdd provides a mock function with given fields: ctx, since, hints
func (_m *LayeredStoreSupplier) TeamMembersToAdd(ctx context.Context, since int64, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, since)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, since, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// TeamMembersToRemove provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreSupplier) TeamMembersToRemove(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}
