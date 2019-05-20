package cmd

import (
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

var NoCache bool

var RootCmd = &cobra.Command{
	Use:   "blueprintz",
	Short: "Manage and use blueprintz for your WordPress websites, plugins and themes.",
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.BoolVarP(&NoCache, "no-cache", "", false, "Disable caching")
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
