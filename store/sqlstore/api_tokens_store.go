// Copyright (c) 2017 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"database/sql"
	"net/http"

	"github.com/uni-x/mattermost-server/mlog"
	"github.com/uni-x/mattermost-server/model"
	"github.com/uni-x/mattermost-server/store"
)

type SqlApiTokenStore struct {
	SqlStore
}

func NewSqlApiTokenStore(sqlStore SqlStore) store.ApiTokenStore {
	s := &SqlApiTokenStore{sqlStore}

	for _, db := range sqlStore.GetAllConns() {
		table := db.AddTableWithName(model.ApiToken{}, "ApiTokens").SetKeys(false, "UserId")
		table.ColMap("UserId").SetMaxSize(64)
		table.ColMap("Token").SetMaxSize(64)
	}

	return s
}

func (s SqlApiTokenStore) CreateIndexesIfNotExists() {
}

func (s SqlApiTokenStore) Delete(userId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		if _, err := s.GetMaster().Exec("DELETE FROM ApiTokens WHERE UserId = :UserId", map[string]interface{}{"UserId": userId}); err != nil {
			result.Err = model.NewAppError("SqlApiTokenStore.Delete", "store.sql_recover.delete.app_error", nil, "", http.StatusInternalServerError)
		}
	})
}

func (s SqlApiTokenStore) Get(userId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		token := &model.ApiToken{}

		err := s.GetReplica().SelectOne(&token, "SELECT * FROM ApiTokens WHERE UserId = :UserId", map[string]interface{}{"UserId": userId})
		if err != nil {
			result.Err = model.NewAppError("SqlApiTokenStore.Get", "store.sql_recover.get.app_error", nil, err.Error(), http.StatusInternalServerError)
			return
		} else if token.IsExpired() {
			result.Err = model.NewAppError("SqlApiTokenStore.Get", "store.sql_recover.expired.app_error", nil, "token expired", http.StatusInternalServerError)
			return
		}

		result.Data = token
	})
}

func (s SqlApiTokenStore) Refresh(userId string) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		token := &model.ApiToken{}

		err := s.GetReplica().SelectOne(&token, "SELECT * FROM ApiTokens WHERE UserId = :UserId", map[string]interface{}{"UserId": userId})
		if err != nil {
			if err == sql.ErrNoRows {
				token = model.NewApiToken(userId)
				transaction, err := s.GetMaster().Begin()
				if err != nil {
					result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
					return
				}

				if _, err := transaction.Exec(`
				INSERT INTO ApiTokens (UserId, Token, CreateAt) VALUES(:UserId, :Token, :CreateAt)
				`, map[string]interface{}{"UserId": userId, "Token": token.Token, "CreateAt": model.GetMillis()}); err != nil {
					result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
					return
				}
				if err := transaction.Commit(); err != nil {
					result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				result.Err = model.NewAppError("SqlApiTokenStore.Get", "store.sql_recover.refresh.app_error", nil, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if token.IsExpired() {
			token = model.NewApiToken(userId)
			transaction, err := s.GetMaster().Begin()
			if err != nil {
				result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err := transaction.Exec(`
			UPDATE ApiTokens SET Token = :Token, CreateAt = :CreateAt WHERE UserId = :UserId
			`, map[string]interface{}{"UserId": userId, "Token": token.Token, "CreateAt": model.GetMillis()}); err != nil {
				result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := transaction.Commit(); err != nil {
				result.Err = model.NewAppError("SqlTokenStore.Save", "store.sql_recover.save.app_error", nil, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		result.Data = token
	})
}

func (s SqlApiTokenStore) Cleanup() {
	mlog.Debug("Cleaning up token store.")
	deltime := model.GetMillis() - model.MAX_TOKEN_EXIPRY_TIME
	if _, err := s.GetMaster().Exec("DELETE FROM ApiTokens WHERE CreateAt < :DelTime", map[string]interface{}{"DelTime": deltime}); err != nil {
		mlog.Error("Unable to cleanup token store.")
	}
}
