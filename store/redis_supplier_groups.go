// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import (
	"context"

	"github.com/uni-x/mattermost-server/model"
)

func (s *RedisSupplier) GroupCreate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupCreate(ctx, group, hints...)
}

func (s *RedisSupplier) GroupGet(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGet(ctx, groupID, hints...)
}

func (s *RedisSupplier) GroupGetByRemoteID(ctx context.Context, remoteID string, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetByRemoteID(ctx, remoteID, groupSource, hints...)
}

func (s *RedisSupplier) GroupGetAllBySource(ctx context.Context, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetAllBySource(ctx, groupSource, hints...)
}

func (s *RedisSupplier) GroupUpdate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupUpdate(ctx, group, hints...)
}

func (s *RedisSupplier) GroupDelete(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupDelete(ctx, groupID, hints...)
}

func (s *RedisSupplier) GroupGetMemberUsers(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetMemberUsers(ctx, groupID, hints...)
}

func (s *RedisSupplier) GroupGetMemberUsersPage(ctx context.Context, groupID string, offset int, limit int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetMemberUsersPage(ctx, groupID, offset, limit, hints...)
}

func (s *RedisSupplier) GroupGetMemberCount(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetMemberCount(ctx, groupID, hints...)
}

func (s *RedisSupplier) GroupCreateOrRestoreMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupCreateOrRestoreMember(ctx, groupID, userID, hints...)
}

func (s *RedisSupplier) GroupDeleteMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupDeleteMember(ctx, groupID, userID, hints...)
}

func (s *RedisSupplier) GroupCreateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupCreateGroupSyncable(ctx, groupSyncable, hints...)
}

func (s *RedisSupplier) GroupGetGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetGroupSyncable(ctx, groupID, syncableID, syncableType, hints...)
}

func (s *RedisSupplier) GroupGetAllGroupSyncablesByGroup(ctx context.Context, groupID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupGetAllGroupSyncablesByGroup(ctx, groupID, syncableType, hints...)
}

func (s *RedisSupplier) GroupUpdateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupUpdateGroupSyncable(ctx, groupSyncable, hints...)
}

func (s *RedisSupplier) GroupDeleteGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().GroupDeleteGroupSyncable(ctx, groupID, syncableID, syncableType, hints...)
}

func (s *RedisSupplier) PendingAutoAddTeamMembers(ctx context.Context, minGroupMembersCreateAt int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().PendingAutoAddTeamMembers(ctx, minGroupMembersCreateAt, hints...)
}

func (s *RedisSupplier) PendingAutoAddChannelMembers(ctx context.Context, minGroupMembersCreateAt int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	// TODO: Redis caching.
	return s.Next().PendingAutoAddChannelMembers(ctx, minGroupMembersCreateAt, hints...)
}
