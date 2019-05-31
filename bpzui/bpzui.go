package bpzui

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

type BpzUi struct {
	App           *tview.Application
	ProjectView   *tview.TreeView
	RightHandView *tview.Flex
	NodeView      *tview.Box
	HelpView      *tview.Box
	FullView      *tview.Flex
}

func New() *BpzUi {
	bpzui := BpzUi{
		App:      tview.NewApplication(),
		NodeView: tview.NewBox().SetBorder(true).SetTitle("Node"),
		HelpView: tview.NewBox().SetBorder(true).SetTitle("Help"),
	}
	pv, sts := bpzui.MakeProjectView()
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}
	bpzui.ProjectView = pv
	bpzui.RightHandView = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(bpzui.NodeView, 0, 618, false).
		AddItem(bpzui.HelpView, 0, 382, false)

	bpzui.FullView = tview.NewFlex().
		AddItem(bpzui.ProjectView, 0, 382, true).
		AddItem(bpzui.RightHandView, 0, 618, false)

	bpzui.App.SetRoot(bpzui.FullView, true)

	// Shortcuts to navigate the slides.
	bpzui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			bpzui.App.SetFocus(bpzui.ProjectView)
		}
		return event
	})

	return &bpzui
}

func (me *BpzUi) Run() (sts Status) {
	err := me.App.Run()
	if err != nil {
		sts = status.Wrap(err)
	}
	return sts
}

func (me *BpzUi) MakeProjectView() (tree *tview.TreeView, sts Status) {

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

		root := tview.NewTreeNode(global.ProjectNode).
			SetColor(tcell.ColorAqua)

		tree = tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root)

		me.ProjectView = tree

		m := make(global.TreeNodeMap, len(global.FirstLevelNodeLabels))
		for _, nl := range global.FirstLevelNodeLabels {
			ref := bpz.GetTreeNoder(nl)
			node := tui.MakeNode(ref)
			root.AddChild(node)
			m[nl] = node
		}

		color := tcell.ColorWhite
		me.ProjectView.SetChangedFunc(func(node *tview.TreeNode) {
			if color == tcell.ColorAqua {
				color = tcell.ColorLime
			} else {
				color = tcell.ColorAqua
			}
			me.NodeView.SetTitleColor(color)
		})

		// If a directory was selected, open it.
		me.ProjectView.SetSelectedFunc(func(node *tview.TreeNode) {
			changefocus := false
			children := node.GetChildren()
			for range only.Once {
				ref := node.GetReference()
				if ref == nil {
					break // Selecting the root node does nothing.
				}
				if len(children) != 0 {
					node.SetExpanded(!node.IsExpanded())
					break
				}
				// Load and show files in this directory.
				tn, ok := ref.(tui.TreeNoder)
				if !ok {
					changefocus = true
					break
				}
				c := tn.GetChildren()
				if c == nil {
					changefocus = true
					break
				}
				for _, cn := range c {
					tui.MakeNode(cn)
					node.AddChild(tui.MakeNode(cn))
				}
			}
			if changefocus {
				me.App.SetFocus(me.NodeView)
			}
		})

		me.ProjectView.SetBorder(true).
			SetBorderPadding(1, 1, 2, 2).
			SetTitle(global.JsonSchemaCreatedBy).
			SetTitleAlign(tview.AlignCenter)

	}
	return tree, sts
}
