package run

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
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

func Browse() (sts Status) {

	for range only.Once {
		var bpz *blueprintz.Blueprintz
		bpz, sts = blueprintz.Load()
		if is.Error(sts) {
			break
		}
		if bpz == nil {
			sts = status.Fail().SetMessage("no '%s' file found in current directory",
				filepath.Base(jsonfile.GetFilepath()),
			)
			break
		}
		println(bpz.Name)

		root := tview.NewTreeNode(global.ProjectNode).
			SetColor(tcell.ColorAqua)

		tree := tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root)

		m := make(global.TreeNodeMap, len(global.FirstLevelNodeLabels))
		for _, nl := range global.FirstLevelNodeLabels {
			ref := bpz.GetTreeNoder(nl)
			node := tui.MakeNode(ref)
			root.AddChild(node)
			m[nl] = node
		}

		// If a directory was selected, open it.
		tree.SetSelectedFunc(func(node *tview.TreeNode) {
			for range only.Once {
				ref := node.GetReference()
				if ref == nil {
					break // Selecting the root node does nothing.
				}
				children := node.GetChildren()
				if len(children) > 0 {
					// Collapse if visible, expand if collapsed.
					node.SetExpanded(!node.IsExpanded())
					break
				}
				// Load and show files in this directory.
				tn, ok := ref.(tui.TreeNoder)
				if !ok {
					break
				}
				c := tn.GetChildren()
				if c == nil {
					break
				}
				for _, cn := range c {
					tui.MakeNode(cn)
					node.AddChild(tui.MakeNode(cn))
				}
			}
		})

		//form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
		ui := tview.NewApplication()
		err := ui.SetRoot(tree, true).Run()

		if err != nil {
			sts = status.Wrap(err)
			break
		}
	}
	return sts
}
