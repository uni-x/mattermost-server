// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package einterfaces

import (
	"io"

	"fmt"
	"github.com/uni-x/mattermost-server/model"
)

type OauthProvider interface {
	GetUserFromJson(data io.Reader, groups []string) *model.User
	GetUsersFromJson(data io.Reader) ([]*model.User, error)
}

var oauthProviders = make(map[string]OauthProvider)

func RegisterOauthProvider(name string, newProvider OauthProvider) {
	fmt.Println("Registering ", name)
	oauthProviders[name] = newProvider
}

func GetOauthProvider(name string) OauthProvider {
	provider, ok := oauthProviders[name]
	if ok {
		return provider
	}
	return nil
}
