// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/uni-x/mattermost-server/store/storetest"
)

func TestChannelMemberHistoryStore(t *testing.T) {
	StoreTest(t, storetest.TestChannelMemberHistoryStore)
}
