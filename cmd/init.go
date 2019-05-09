package cmd

import (
	"blueprintz/blueprintz"
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

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a blueprintz.json file from an existing WordPress install",
	Run: func(cmd *cobra.Command, args []string) {
		blueprintz.Init()
	},
}

func init() {
	//	fs := InitCmd.Flags()
	//	fs.StringVarP(&global.ListDomain, global.SvnListFlag, "", global.SvnListDomain, "SVN domain to Init from?")
	RootCmd.AddCommand(InitCmd)
}