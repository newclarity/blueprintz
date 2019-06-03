package cmd

import (
	"blueprintz/global"
	"blueprintz/run"
	"fmt"
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

var TriageCmd = &cobra.Command{
	Use: "triage",
	Short: fmt.Sprintf("Run Triage UI to fill in missing info in '%s'",
		global.BlueprintzFile,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		return run.Triage()
	},
}

func init() {
	//	fs := BuildCmd.Flags()
	//	fs.StringVarP(&global.ListDomain, global.SvnListFlag, "", global.SvnListDomain, "SVN domain to Build from?")
	RootCmd.AddCommand(TriageCmd)
}
