// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package testlib

import (
	"flag"
	"os"
	"testing"

	"github.com/uni-x/mattermost-server/mlog"
	"github.com/uni-x/mattermost-server/model"
	"github.com/uni-x/mattermost-server/store"
	"github.com/uni-x/mattermost-server/store/sqlstore"
	"github.com/uni-x/mattermost-server/store/storetest"
	"github.com/uni-x/mattermost-server/utils"
)

type MainHelper struct {
	Settings         *model.SqlSettings
	Store            store.Store
	SqlSupplier      *sqlstore.SqlSupplier
	ClusterInterface *FakeClusterInterface

	status int
}

func NewMainHelper() *MainHelper {
	flag.Parse()

	// Setup a global logger to catch tests logging outside of app context
	// The global logger will be stomped by apps initalizing but that's fine for testing.
	// Ideally this won't happen.
	mlog.InitGlobalLogger(mlog.NewLogger(&mlog.LoggerConfiguration{
		EnableConsole: true,
		ConsoleJson:   true,
		ConsoleLevel:  "error",
		EnableFile:    false,
	}))

	utils.TranslationsPreInit()

	settings := storetest.MakeSqlSettings(model.DATABASE_DRIVER_MYSQL)

	clusterInterface := &FakeClusterInterface{}
	sqlSupplier := sqlstore.NewSqlSupplier(*settings, nil)
	testStore := &TestStore{
		store.NewLayeredStore(sqlSupplier, nil, clusterInterface),
	}

	return &MainHelper{
		Settings:         settings,
		Store:            testStore,
		SqlSupplier:      sqlSupplier,
		ClusterInterface: clusterInterface,
	}
}

func (h *MainHelper) Main(m *testing.M) {
	h.status = m.Run()
}

func (h *MainHelper) Close() error {
	storetest.CleanupSqlSettings(h.Settings)

	os.Exit(h.status)

	return nil
}
