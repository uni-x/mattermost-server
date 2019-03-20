// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ChannelCreds struct {
	ChannelId string
	Owners    struct {
		Groups []string
		Users  []string
	}
	Moderators struct {
		Groups []string
		Users  []string
	}
	Members struct {
		Groups []string
		Users  []string
	}
	Repliers struct {
		Groups []string
		Users  []string
	}
	Viewers struct {
		Groups []string
		Users  []string
	}
}

type ChannelCredElement struct {
	ChannelId   string
	UserId      string
	AzureRole   string
	ChannelRole string
}

type ChannelCredsRequest struct {
	ChannelId   string
	UserId      string
	AzureRoles  []string
	ChannelRole string
}

func ChannelCredsRequestFromJson(data io.Reader) *ChannelCredsRequest {
	var o *ChannelCredsRequest
	json.NewDecoder(data).Decode(&o)
	return o
}

func ChannelCredsFromJson(data io.Reader) *ChannelCreds {
	var o *ChannelCreds
	json.NewDecoder(data).Decode(&o)
	return o
}
