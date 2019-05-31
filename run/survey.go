package run

import (
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

// Survey and existing project and prepare a working blueprintz.json file
// - Initialize blueprintz.json
// - Survey files
// - Research classifications
// - Solicit input for details
// - Other?

func Survey() (sts Status) {

	for range only.Once {
		sts = Init()
		if is.Error(sts) {
			break
		}
		sts = Scan()
		if is.Error(sts) {
			break
		}
		sts = Research()
		if is.Error(sts) {
			break
		}
	}

	return sts

}
