// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthoffice365

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"

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
	ThumbnailPhoto    string `json:"thumbnailphoto"`
	Groups            []string
}

func init() {
	provider := &Office365Provider{}
	einterfaces.RegisterOauthProvider(model.USER_AUTH_SERVICE_OFFICE365, provider)
}

func userFromOffice365User(office365User *Office365User) (*model.User, io.Reader) {
	user := &model.User{}
	username := office365User.DisplayName
	if username == "" {
		username = office365User.GivenName + "_" + office365User.Surname
	}
	user.Username = model.CleanUsername(username)
	user.FirstName = office365User.GivenName
	user.LastName = office365User.Surname
	user.Email = office365User.Mail
	if office365User.Mail != "" {
		user.Email = office365User.Mail
	} else {
		user.Email = office365User.UserPrincipalName
	}
	azureGroups := "ALL"
	for _, group := range office365User.Groups {
		mmGroup, exists := model.GroupsMapping[group]
		if exists {
			azureGroups = azureGroups + " " + mmGroup
		}
	}
	user.AzureGroups = azureGroups
	userId := office365User.Id
	user.AuthData = &userId
	user.AuthService = model.USER_AUTH_SERVICE_OFFICE365
	fmt.Println("IMAGE:", len(office365User.ThumbnailPhoto))
	if office365User.ThumbnailPhoto == "" {
		office365User.ThumbnailPhoto = DEFAULT_PROFILE_IMAGE
	}

	fileBytes, err := getFileFromBlob(office365User.ThumbnailPhoto)
	if err != nil {
		fileBytes = []byte{}
	}
	return user, bytes.NewReader(fileBytes)
}

func getFileFromBlob(blob string) ([]byte, error) {
	blobParts := strings.Split(blob, ";base64,")
	if len(blobParts) > 1 {
		blob = blobParts[1]
	}
	decoded, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		if err != nil {
			return []byte{}, err
		}
	}
	return decoded, nil
}

func office365UserFromJson(data io.Reader, groups []string) *Office365User {
	decoder := json.NewDecoder(data)
	var office365User Office365User
	err := decoder.Decode(&office365User)
	if err == nil {
		office365User.Groups = groups
		return &office365User
	} else {
		return nil
	}
}

func (office365User *Office365User) ToJson() string {
	b, err := json.Marshal(office365User)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (office365User *Office365User) IsValid() bool {
	if office365User.Id == "" {
		return false
	}
	if len(office365User.Mail) == 0 && len(office365User.UserPrincipalName) == 0 {
		return false
	}
	return true
}

func (office365User *Office365User) getAuthData() string {
	return office365User.Id
}

func (m *Office365Provider) GetUserFromJson(data io.Reader, groups []string) (*model.User, io.Reader) {
	office365User := office365UserFromJson(data, groups)
	if office365User.IsValid() {
		return userFromOffice365User(office365User)
	}
	return &model.User{}, nil
}

func (m *Office365Provider) GetUsersFromJson(data io.Reader) ([]*model.User, error) {
	decoder := json.NewDecoder(data)
	var office365Users []Office365User
	var users []*model.User
	err := decoder.Decode(&office365Users)
	if err != nil {
		return nil, err
	} else {
		for _, office365User := range office365Users {
			user, _ := userFromOffice365User(&office365User)
			if user != nil {
				users = append(users, user)
			}
		}
	}
	return users, nil
}
