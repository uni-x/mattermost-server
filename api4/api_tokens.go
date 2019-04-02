// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"net/http"

	"github.com/uni-x/mattermost-server/model"
)

func (api *API) InitApiTokens() {
	api.BaseRoutes.ApiTokens.Handle("/get", api.ApiHandler(getApiToken)).Methods("POST")
}

func getApiToken(c *Context, w http.ResponseWriter, r *http.Request) {
	apiTokenRequest := model.ApiTokenRequestFromJson(r.Body)
	if apiTokenRequest == nil {
		c.SetInvalidParam("token")
		return
	}

	apiToken, err := c.App.RefreshApiToken(apiTokenRequest.UserId, apiTokenRequest.RefreshToken)
	if err != nil {
		c.Err = err
		return
	}
	w.Write([]byte(apiToken))
}
