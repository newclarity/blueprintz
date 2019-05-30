package cmd

import (
	"blueprintz/jsonfile"
	"blueprintz/run"
	"fmt"
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

var BrowseCmd = &cobra.Command{
	Use: "browse",
	Short: fmt.Sprintf("Browse and edit a '%s' file in a tree view",
		filepath.Base(jsonfile.GetFilepath()),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		return run.Browse()
	},
}

func init() {
	//	fs := BuildCmd.Flags()
	//	fs.StringVarP(&global.ListDomain, global.SvnListFlag, "", global.SvnListDomain, "SVN domain to Build from?")
	RootCmd.AddCommand(BrowseCmd)
}