package run

import (
	"blueprintz/jsonfile"
	"fmt"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
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

func Compose() (sts Status) {

	for range only.Once {

		sts = EnsureBlueprintJsonExists()
		if is.Error(sts) {
			break
		}
		jbp := jsonfile.Blueprintz{}
		sts = jbp.LoadFile()
		if is.Error(sts) {
			break
		}

		r, err := git.PlainClone("/tmp/go-status",
			false,
			&git.CloneOptions{
				URL:        "https://github.com/gearboxworks/go-status.git",
				RemoteName: "origin",
				Progress:   &Out{},
			},
		)
		if err != nil {
			panic(err)
		}
		var cos object.CommitIter
		cos, err = r.CommitObjects()
		if err != nil {
			break
		}
		for {
			var co *object.Commit
			co, err = cos.Next()
			if err != nil {
				break
			}
			fmt.Println(co.Message)
		}

		fmt.Printf("WordPress project '%s' composed.",
			jsonfile.GetBasefile(),
		)
	}
	return sts
}

type Out struct{}

func (me *Out) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
