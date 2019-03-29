// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthoffice365

import (
	"encoding/json"
	"io"

	"github.com/uni-x/mattermost-server/einterfaces"
	"github.com/uni-x/mattermost-server/model"
)

type Office365Provider struct {
}

type Office365User struct {
	Id                string `json:"id"`
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	Surname           string `json:"surname"`
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
	Groups            []string
}

func init() {
	provider := &Office365Provider{}
	einterfaces.RegisterOauthProvider(model.USER_AUTH_SERVICE_OFFICE365, provider)
}

func userFromOffice365User(glu *Office365User) *model.User {
	user := &model.User{}
	username := glu.DisplayName
	if username == "" {
		username = glu.GivenName + "_" + glu.Surname
	}
	user.Username = model.CleanUsername(username)
	user.FirstName = glu.GivenName
	user.LastName = glu.Surname
	user.Email = glu.Mail
	if glu.Mail != "" {
		user.Email = glu.Mail
	} else {
		user.Email = glu.UserPrincipalName
	}
	azureGroups := "ALL"
	for _, group := range glu.Groups {
		mmGroup, exists := model.GroupsMapping[group]
		if exists {
			azureGroups = azureGroups + " " + mmGroup
		}
	}
	user.AzureGroups = azureGroups
	userId := glu.Id
	user.AuthData = &userId
	user.AuthService = model.USER_AUTH_SERVICE_OFFICE365

	return user
}

func office365UserFromJson(data io.Reader, groups []string) *Office365User {
	decoder := json.NewDecoder(data)
	var glu Office365User
	err := decoder.Decode(&glu)
	if err == nil {
		glu.Groups = groups
		return &glu
	} else {
		return nil
	}
}

func (glu *Office365User) ToJson() string {
	b, err := json.Marshal(glu)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (glu *Office365User) IsValid() bool {
	if glu.Id == "" {
		return false
	}
	if len(glu.Mail) == 0 && len(glu.UserPrincipalName) == 0 {
		return false
	}
	return true
}

func (glu *Office365User) getAuthData() string {
	return glu.Id
}

func (m *Office365Provider) GetUserFromJson(data io.Reader, groups []string) *model.User {
	glu := office365UserFromJson(data, groups)
	if glu.IsValid() {
		return userFromOffice365User(glu)
	}
	return &model.User{}
}

func (m *Office365Provider) GetUsersFromJson(data io.Reader) ([]*model.User, error) {
	decoder := json.NewDecoder(data)
	var glus []Office365User
	var users []*model.User
	err := decoder.Decode(&glus)
	if err != nil {
		return nil, err
	} else {
		for _, glu := range glus {
			user := userFromOffice365User(&glu)
			if user != nil {
				users = append(users, user)
			}
		}
	}
	return users, nil
}
