// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"net/http"
	"strings"

	"github.com/uni-x/mattermost-server/model"
)

// display name, name. type, team_id
func (a *App) GetApiToken(userId string) (string, *model.AppError) {
	_, err := a.GetUserByAuth(&userId, "office365")
	if err != nil {
		return "", err
	}

	result := <-a.Srv.Store.ApiToken().Get(userId)
	if result.Err != nil {
		return "", result.Err
	}

	apiToken := result.Data.(*model.ApiToken)
	return apiToken.Token, nil
}

func (a *App) RefreshApiToken(userId, refreshToken string) (string, *model.AppError) {
	user, err := a.GetUserByAuth(&userId, "office365")
	if err != nil {
		return "", err
	}

	if !strings.Contains(user.Roles, "system_admin") {
		err := model.NewAppError("GetApiToken", "api.user.get_api_token.user_is_not_a_system_admin.app_error", nil, "", http.StatusNotImplemented)
		return "", err
	}

	if refreshToken != a.Config().ApiSettings.RefreshToken {
		err := model.NewAppError("GetApiToken", "api.user.get_api_token.incorrect_refresh_token.app_error", nil, "", http.StatusNotImplemented)
		return "", err
	}
	result := <-a.Srv.Store.ApiToken().Refresh(userId)
	if result.Err != nil {
		return "", result.Err
	}

	apiToken := result.Data.(*model.ApiToken)
	return apiToken.Token, nil
}
