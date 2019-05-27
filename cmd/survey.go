package cmd

import (
	"blueprintz/run"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/spf13/cobra"
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

// Survey and existing project and prepare a working blueprintz.json file
// - Initialize blueprintz.json
// - Survey files
// - Research classifications
// - Solicit input for details
// - Other?

var SurveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "Survey a project for core, plugins, themes, etc.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var sts Status
		for range only.Once {
			sts = run.Init()
			if is.Error(sts) {
				break
			}
			sts = run.Scan()
			if is.Error(sts) {
				break
			}
			sts = run.Research()
			if is.Error(sts) {
				break
			}
		}
		return sts
	},
}

func init() {
	//	fs := SurveyCmd.Flags()
	//	fs.StringVarP(&global.ListDomain, global.SvnListFlag, "", global.SvnListDomain, "SVN domain to Survey from?")
	RootCmd.AddCommand(SurveyCmd)
}
