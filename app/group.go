// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"github.com/uni-x/mattermost-server/model"
)

func (a *App) GetGroup(id string) (*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().Get(id)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.Group), nil
}

func (a *App) GetGroupByRemoteID(remoteID string, groupSource model.GroupSource) (*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().GetByRemoteID(remoteID, groupSource)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.Group), nil
}

func (a *App) GetGroupsBySource(groupSource model.GroupSource) ([]*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().GetAllBySource(groupSource)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.([]*model.Group), nil
}

func (a *App) CreateGroup(group *model.Group) (*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().Create(group)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.Group), nil
}

func (a *App) UpdateGroup(group *model.Group) (*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().Update(group)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.Group), nil
}

func (a *App) DeleteGroup(groupID string) (*model.Group, *model.AppError) {
	result := <-a.Srv.Store.Group().Delete(groupID)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.Group), nil
}

func (a *App) GetGroupMemberUsers(groupID string) ([]*model.User, *model.AppError) {
	result := <-a.Srv.Store.Group().GetMemberUsers(groupID)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.([]*model.User), nil
}

func (a *App) GetGroupMemberUsersPage(groupID string, page int, perPage int) ([]*model.User, int, *model.AppError) {
	result := <-a.Srv.Store.Group().GetMemberUsersPage(groupID, page, perPage)
	if result.Err != nil {
		return nil, 0, result.Err
	}
	members := result.Data.([]*model.User)
	result = <-a.Srv.Store.Group().GetMemberCount(groupID)
	if result.Err != nil {
		return nil, 0, result.Err
	}
	count := int(result.Data.(int64))
	return members, count, nil
}

func (a *App) CreateOrRestoreGroupMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	result := <-a.Srv.Store.Group().CreateOrRestoreMember(groupID, userID)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupMember), nil
}

func (a *App) DeleteGroupMember(groupID string, userID string) (*model.GroupMember, *model.AppError) {
	result := <-a.Srv.Store.Group().DeleteMember(groupID, userID)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupMember), nil
}

func (a *App) CreateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	result := <-a.Srv.Store.Group().CreateGroupSyncable(groupSyncable)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupSyncable), nil
}

func (a *App) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	result := <-a.Srv.Store.Group().GetGroupSyncable(groupID, syncableID, syncableType)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupSyncable), nil
}

func (a *App) GetGroupSyncables(groupID string, syncableType model.GroupSyncableType) ([]*model.GroupSyncable, *model.AppError) {
	result := <-a.Srv.Store.Group().GetAllGroupSyncablesByGroupId(groupID, syncableType)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.([]*model.GroupSyncable), nil
}

func (a *App) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) (*model.GroupSyncable, *model.AppError) {
	result := <-a.Srv.Store.Group().UpdateGroupSyncable(groupSyncable)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupSyncable), nil
}

func (a *App) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) (*model.GroupSyncable, *model.AppError) {
	result := <-a.Srv.Store.Group().DeleteGroupSyncable(groupID, syncableID, syncableType)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.(*model.GroupSyncable), nil
}

func (a *App) PendingAutoAddTeamMembers(minGroupMembersCreateAt int64) ([]*model.UserTeamIDPair, *model.AppError) {
	result := <-a.Srv.Store.Group().PendingAutoAddTeamMembers(minGroupMembersCreateAt)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.([]*model.UserTeamIDPair), nil
}

func (a *App) PendingAutoAddChannelMembers(minGroupMembersCreateAt int64) ([]*model.UserChannelIDPair, *model.AppError) {
	result := <-a.Srv.Store.Group().PendingAutoAddChannelMembers(minGroupMembersCreateAt)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Data.([]*model.UserChannelIDPair), nil
}
