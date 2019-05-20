package run

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
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

func Scan() (sts Status) {
	for range only.Once {
		if !util.FileExists(jsonfile.GetFilepath()) {
			sts = status.YourBad("The file '%s' does not exist; run `bpz init` to create one.",
				jsonfile.GetBasefile(),
			)
			break
		}
		jbp := jsonfile.NewBlueprintz()
		sts = jbp.LoadFile()
		if is.Error(sts) {
			break
		}
		bpz := blueprintz.NewBlueprintzFromJsonfile(jbp)
		bpz.RegisterRecognizer(
			global.WordPressOrgRecognizer,
			recognize.NewWordPressOrg(),
		)
		sts = bpz.Scandir()
		if is.Error(sts) {
			break
		}

		jbp = jsonfile.NewBlueprintzFromBlueprintz(bpz)
		sts = jbp.WriteFile()
		if is.Error(sts) {
			break
		}
		fmt.Printf("The file '%s' was scanned.", jsonfile.GetBasefile())
	}
	return sts
}
