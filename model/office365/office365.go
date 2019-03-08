// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthoffice365

import (
	"encoding/json"
	"io"
	//"strconv"
	"strings"

	"fmt"
	"github.com/uni-x/mattermost-server/einterfaces"
	"github.com/uni-x/mattermost-server/model"
)

type Office365Provider struct {
}

type Office365User struct {
	Id string `json:"id"`
	/*	Username string `json:"username"`
		Login    string `json:"login"`
		Email    string `json:"email"`
		Name     string `json:"name"`*/
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	Surname           string `json:"surname"`
	JobTitle          string `json:"jobTitle"`
	UserPrincipalName string `json:"userPrincipalName"`
}

/*
{
	"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#users/$entity",
	"businessPhones":[],
	"displayName":"TBrady Wilson",
	"givenName":"Brady",
	"jobTitle":"STUDENT",
	"mail":null,
	"mobilePhone":null,
	"officeLocation":null,
	"preferredLanguage":null,
	"surname":"Wilson",
	"userPrincipalName":"tbwilson@cgs.ad",
	"id":"f7bc158a-8b95-4aae-9fef-42086fe7996b"}
*/
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
	/*	splitName := strings.Split(glu.Name, " ")
		if len(splitName) == 2 {
			user.FirstName = splitName[0]
			user.LastName = splitName[1]
		} else if len(splitName) >= 2 {
			user.FirstName = splitName[0]
			user.LastName = strings.Join(splitName[1:], " ")
		} else {
			user.FirstName = glu.Name
		}*/
	user.FirstName = glu.GivenName
	user.LastName = glu.Surname
	user.Email = glu.UserPrincipalName
	user.Roles = strings.ToLower(glu.JobTitle)
	userId := glu.Id
	user.AuthData = &userId
	user.AuthService = model.USER_AUTH_SERVICE_OFFICE365

	return user
}

func office365UserFromJson(data io.Reader) *Office365User {
	decoder := json.NewDecoder(data)
	var glu Office365User
	err := decoder.Decode(&glu)
	if err == nil {
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

	if len(glu.UserPrincipalName) == 0 {
		return false
	}

	return true
}

func (glu *Office365User) getAuthData() string {
	return glu.Id
}

func (m *Office365Provider) GetUserFromJson(data io.Reader) *model.User {
	glu := office365UserFromJson(data)
	fmt.Println(glu)
	if glu.IsValid() {
		return userFromOffice365User(glu)
	}

	return &model.User{}
}
