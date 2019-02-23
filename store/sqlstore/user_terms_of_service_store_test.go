package sqlstore

import (
	"testing"

	"github.com/uni-x/mattermost-server/store/storetest"
)

func TestUserTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestUserTermsOfServiceStore)
}
