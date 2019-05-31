package main

import (
	"blueprintz/cmd"
)

//
// Name: Blueprintz for WordPress
// Description: Manage and use blueprintz for your WordPress websites, plugins and themes
// Version 1.0
// Author: Mike Schinkel <mike@newclarity.net>
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

func main() {
	_ = cmd.RootCmd.Execute()
	//sts,ok := err.(status.Status)
	//if !ok {
	//	log.Fatal(err)
	//}
	//if is.Error(sts) {
	//	log.Fatal(sts.LongMessage())
	//}
	//if is.Success(sts) && sts != nil {
	//	log.Println(sts.Message())
	//}
}
