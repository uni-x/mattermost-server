// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package commands

import (
	"flag"
	"os"
	"testing"

	"github.com/uni-x/mattermost-server/api4"
	"github.com/uni-x/mattermost-server/testlib"
)

func TestMain(m *testing.M) {
	// Command tests are run by re-invoking the test binary in question, so avoid creating
	// another container when we detect same.
	flag.Parse()
	if filter := flag.Lookup("test.run").Value.String(); filter == "ExecCommand" {
		status := m.Run()
		os.Exit(status)
		return
	}

	mainHelper = testlib.NewMainHelper()
	defer mainHelper.Close()
	api4.UseTestStore(mainHelper.Store)

	mainHelper.Main(m)
}
