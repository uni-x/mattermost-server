// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package commands

import (
	"github.com/uni-x/mattermost-server/app"
	"github.com/uni-x/mattermost-server/model"
	"github.com/uni-x/mattermost-server/utils"
	"github.com/spf13/cobra"
)

func InitDBCommandContextCobra(command *cobra.Command) (*app.App, error) {
	config, err := command.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	a, err := InitDBCommandContext(config)

	if err != nil {
		// Returning an error just prints the usage message, so actually panic
		panic(err)
	}

	a.InitPlugins(*a.Config().PluginSettings.Directory, *a.Config().PluginSettings.ClientDirectory)
	a.DoAdvancedPermissionsMigration()
	a.DoEmojisPermissionsMigration()

	return a, nil
}

func InitDBCommandContext(configFileLocation string) (*app.App, error) {
	if err := utils.TranslationsPreInit(); err != nil {
		return nil, err
	}
	model.AppErrorInit(utils.T)

	s, err := app.NewServer(app.ConfigFile(configFileLocation))
	if err != nil {
		return nil, err
	}

	a := s.FakeApp()

	if model.BuildEnterpriseReady == "true" {
		a.LoadLicense()
	}

	return a, nil
}
