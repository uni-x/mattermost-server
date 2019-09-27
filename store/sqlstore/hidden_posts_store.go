// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"net/http"

	"github.com/uni-x/mattermost-server/einterfaces"
	"github.com/uni-x/mattermost-server/model"
	"github.com/uni-x/mattermost-server/store"
)

type SqlHiddenPostsStore struct {
	SqlStore
	metrics           einterfaces.MetricsInterface
}

func NewSqlHiddenPostsStore(sqlStore SqlStore, metrics einterfaces.MetricsInterface) store.HiddenPostsStore {
	s := &SqlHiddenPostsStore{
		SqlStore:          sqlStore,
		metrics:           metrics,
	}

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(HiddenPost{}, "HiddenPosts").SetKeys(false, "UserId", "PostId")
		table.ColMap("PostId").SetMaxSize(26)
		table.ColMap("UserId").SetMaxSize(26)
	}

	return s
}

func (s *SqlHiddenPostsStore) CreateIndexesIfNotExists() {
}

type HiddenPost struct {
	PostId string
	UserId string
}

func (s *SqlHiddenPostsStore) Save(postId, userId string) store.StoreChannel {

	return store.Do(func(result *store.StoreResult) {
		var existing []*HiddenPost
		_, err := s.GetReplica().Select(&existing, "SELECT * FROM HiddenPosts WHERE PostId = :PostId AND UserId = :UserId", map[string]interface{}{"PostId": postId, "UserId": userId})
		if err != nil {
			result.Err = model.NewAppError("SqlPostStore.Save", "store.sql_post.save.existing.app_error", nil, "error="+err.Error(), http.StatusBadRequest)
			return
		}
		if len(existing) == 0 {
			post := &HiddenPost{postId, userId}
			if err := s.GetMaster().Insert(post); err != nil {
				result.Err = model.NewAppError("SqlHiddenPostStore.Save", "store.sql_hidden_post.save.app_error", nil, "error="+err.Error(), http.StatusBadRequest)
			}
		}
	})
}

func (s *SqlHiddenPostsStore) Get(postId, userId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		var existing []*HiddenPost
		_, err := s.GetReplica().Select(&existing, "SELECT * FROM HiddenPosts WHERE PostId = :PostId AND UserId = :UserId", map[string]interface{}{"PostId": postId, "UserId": userId})
		if err != nil {
			result.Err = model.NewAppError("SqlHiddenPostStore.Get", "store.sql_hidden_post.get.app_error", nil, "error="+err.Error(), http.StatusBadRequest)
			return
		}
		result.Data = len(existing) != 0
	})
}

func (s *SqlHiddenPostsStore) Delete(postId, userId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		_, err := s.GetMaster().Exec("DELETE FROM HiddenPosts WHERE PostId = :PostId AND UserId = :UserId", map[string]interface{}{"PostId": postId, "UserId": userId})
		if err != nil {
			result.Err = model.NewAppError("SqlHiddenPostsStore.Delete", "store.sql_post.delete.app_error", nil, "id="+postId+", err="+err.Error(), http.StatusInternalServerError)
		}
	})
}

func (s *SqlHiddenPostsStore) DeletePost(postId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {

		_, err := s.GetMaster().Exec("DELETE FROM HiddenPosts WHERE PostId = :PostId", map[string]interface{}{"PostId": postId})
		if err != nil {
			result.Err = model.NewAppError("SqlHiddenPostsStore.Delete", "store.sql_post.delete.app_error", nil, "id="+postId+", err="+err.Error(), http.StatusInternalServerError)
		}
	})
}
