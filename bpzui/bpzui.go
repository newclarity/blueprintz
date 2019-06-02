package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/rivo/tview"
)

type BpzUi struct {
	Blueprintz    *blueprintz.Blueprintz
	App           *tview.Application
	FullView      *tview.Flex
	ProjectNode   *tview.TreeView
	RightHandView *tview.Flex
	NodeView      *tview.Form
	HelpView      *tview.TextView
}

func New(bpz *blueprintz.Blueprintz) *BpzUi {
	bpzui := BpzUi{
		Blueprintz: bpz,
		App:        tview.NewApplication(),
		HelpView:   tview.NewTextView(),
	}

	sts := bpz.LoadJsonfile()
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}

	pn := NewProjectNode(&bpzui)

	bpzui.ProjectNode = pn.Tree
	bpzui.HelpView.SetBorder(true).SetTitle("Help")
	bpzui.NodeView = pn.Form

	bpzui.RightHandView = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(bpzui.NodeView, 0, 618, false).
		AddItem(bpzui.HelpView, 0, 382, false)

	bpzui.FullView = tview.NewFlex().
		AddItem(bpzui.ProjectNode, 0, 382, true).
		AddItem(bpzui.RightHandView, 0, 618, false)

	bpzui.App.SetRoot(bpzui.FullView, true)

	// Shortcuts to navigate the slides.
	bpzui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			bpzui.App.SetFocus(bpzui.ProjectNode)
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

var coreBlueprintTypes = blueprintDropdownTypes()

func blueprintDropdownTypes() global.BlueprintTypes {
	ts := global.AllBlueprintTypes
	ts = append(ts, "")
	copy(ts[1:], ts[0:])
	ts[0] = "Select a Blueprint Type"
	return ts
}

func (me *BpzUi) MakeNodeView() (form *tview.Box, sts Status) {
	return tview.NewBox().SetBorder(true).SetTitle("Node"), nil
}