// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthapple

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/uni-x/mattermost-server/einterfaces"
	"github.com/uni-x/mattermost-server/model"
)

type AppleProvider struct {
}

type AppleUser struct {
	Id        string  `json:"id"`
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
}

func init() {
	provider := &AppleProvider{}
	einterfaces.RegisterOauthProvider("apple", provider)
}

func userFromAppleUser(au *AppleUser) *model.User {
	user := &model.User{}
	username := au.FirstName+" "+au.LastName
	if username == "" {
		username = au.Name
		if username == "" {
			username = au.Sub
		}
		splitName := strings.Split(username, " ")
		if len(splitName) == 2 {
			user.FirstName = splitName[0]
			user.LastName = splitName[1]
		} else if len(splitName) >= 2 {
			user.FirstName = splitName[0]
			user.LastName = strings.Join(splitName[1:], " ")
		} else {
			user.FirstName = au.Name
		}
	} else {
		user.FirstName = au.FirstName
		user.LastName = au.LastName
	}
	if len(username) > 22 { username = username[:22] }
	user.Username = model.CleanUsername(username)
	if au.Email != "" {
		user.Email = au.Email
	} else {
		user.Email = au.Sub+"@apple.com"
	}
	if au.Id != "" {
		user.AuthData = &au.Id
	} else {
		user.AuthData = &au.Sub
	}
	user.AuthService = "apple"

	return user
}

func appleUserFromJson(data io.Reader) *AppleUser {
	decoder := json.NewDecoder(data)
	var au AppleUser
	err := decoder.Decode(&au)
	if err == nil {
		return &au
	} else {
		return nil
	}
}

func (au *AppleUser) ToJson() string {
	b, err := json.Marshal(au)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (au *AppleUser) IsValid() bool {
	if au.Sub == "" && au.Id == "" {
		return false
	}

	/*if len(au.Email) == 0 {
		return false
	}*/

	return true
}

func (m *AppleProvider) GetUserFromJson(data io.Reader) *model.User {
	au := appleUserFromJson(data)
	if au.IsValid() {
		return userFromAppleUser(au)
	}

	return &model.User{}
}
