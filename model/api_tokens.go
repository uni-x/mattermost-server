// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiTokenRequest struct {
	UserId       string
	RefreshToken string
}

func ApiTokenRequestFromJson(data io.Reader) *ApiTokenRequest {
	var o *ApiTokenRequest
	json.NewDecoder(data).Decode(&o)
	return o
}

type ApiToken struct {
	UserId   string
	Token    string
	CreateAt int64
}

const (
	API_TOKEN_SIZE            = 64
	API_MAX_TOKEN_EXIPRY_TIME = 1000 * 60 * 60 * 24 // 24 hour
)

func NewApiToken(userId string) *ApiToken {
	return &ApiToken{
		UserId:   userId,
		Token:    NewRandomString(API_TOKEN_SIZE),
		CreateAt: GetMillis(),
	}
}

func (t *ApiToken) IsExpired() bool {
	return GetMillis() >= t.CreateAt+API_MAX_TOKEN_EXIPRY_TIME
}

func (t *ApiToken) IsValid() *AppError {
	if len(t.Token) != API_TOKEN_SIZE {
		return NewAppError("ApiToken.IsValid", "model.api_token.is_valid.size", nil, "", http.StatusInternalServerError)
	}

	if t.CreateAt == 0 {
		return NewAppError("ApiToken.IsValid", "model.api_token.is_valid.expiry", nil, "", http.StatusInternalServerError)
	}

	return nil
}
