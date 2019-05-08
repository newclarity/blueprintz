package blueprintz

import (
	"blueprintz/jsonfile"
	"blueprintz/only"
	"blueprintz/util"
	"github.com/gearboxworks/go-status/is"
	"log"
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

func Generate() {
	var sts Status
	for range only.Once {
		bpz := NewBlueprintz(&Args{
			Name: filepath.Base(util.GetCurrentDir()),
		})
		sts = bpz.Layout.ScanDir()
		if is.Error(sts) {
			break
		}
		jsbpz := jsonfile.NewBlueprintz(bpz)
		sts = jsbpz.WriteFile()
		if is.Error(sts) {
			break
		}
	}
	if is.Error(sts) {
		log.Fatal(sts.FullError())
	}
}
