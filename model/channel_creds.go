// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type ChannelCredsSet struct {
	AzureGroups []string
	Users       []string
}

type ChannelCreds struct {
	Owners     *ChannelCredsSet
	Moderators *ChannelCredsSet
	Members    *ChannelCredsSet
	Repliers   *ChannelCredsSet
	Viewers    *ChannelCredsSet
}

type ChannelCredElement struct {
	ChannelId   string
	UserId      string
	AzureGroup  string
	ChannelRole string
}

type ChannelCredsRequest struct {
	UserId      string
	AzureGroups []string
	ChannelRole string
}

func ChannelCredsRequestFromJson(data io.Reader) *ChannelCredsRequest {
	var o *ChannelCredsRequest
	json.NewDecoder(data).Decode(&o)
	return o
}

func ChannelCredsFromJson(data io.Reader) *ChannelCreds {
	var o *ChannelCreds
	err := json.NewDecoder(data).Decode(&o)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	return o
}
