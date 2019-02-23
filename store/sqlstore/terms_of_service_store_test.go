package sqlstore

import (
	"testing"

	"github.com/uni-x/mattermost-server/store/storetest"
)

func TestTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestTermsOfServiceStore)
}
