package cmd

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/log"
	"blueprintz/recognize"
	"blueprintz/util"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/spf13/cobra"
	"path/filepath"
)

//
// Name: Blueprintz for WordPress
//
// Copyright (C) 2019 NewClarity Consulting LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

var RootCmd = &cobra.Command{
	Use:   "blueprintz",
	Short: "Manage and use blueprints for your WordPress websites, plugins and themes.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		for range only.Once {
			blueprintz.Instance = blueprintz.NewBlueprintz(&blueprintz.Args{
				Name:     filepath.Base(util.GetProjectDir()),
				OsBridge: blueprintz.GetOsBridge(global.AppName, global.UserDataPath),
			})

			sts := blueprintz.Instance.Config.Initialize()
			if is.Error(sts) {
				err = sts.Cause()
				break
			}

			blueprintz.Instance.RegisterRecognizer(
				global.WordPressOrgRecognizer,
				recognize.NewWordPressOrg(),
			)

			status.Logger = log.NewLogger()
		}
		return err
	},
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.BoolVarP(&global.NoCache, "no-cache", "", false, "Disable caching")
	pf.StringVarP(&global.ProjectDir, "project-dir", "", util.GetCurrentDir(), "Project directory")
}

/*
 * Possible command names
 *
 *	- survey - Document core, plugins and theme
 *	- research - Discover plug download URLs
 *	- extract
 *	- build
 *	- assemble
 *	- verify
 *	- validate
 *	- update
 *	- upgrade
 *	- generate
 *	- create
 *	- new
 *	- init
 *	- make
 *	- run
 *	- download
 *	- upload
 *	- package
 *	- bundle
 *	- ???
 *	- ???
 *	- ???
 *	- ???
 *	- ???
 *
 */
