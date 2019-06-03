package bpzui

import (
	"blueprintz/blueprintz"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

type BpzUi struct {
	Blueprintz    *blueprintz.Blueprintz
	App           *tview.Application
	FullView      *tview.Flex
	ProjectBox    *tview.TreeView
	RightHandView *tview.Flex
	FormBox       *tview.Form
	HelpBox       *tview.TextView
}

func New(bpz *blueprintz.Blueprintz) *BpzUi {
	app := tview.NewApplication()
	bpzui := BpzUi{
		Blueprintz: bpz,
		App:        app,
		HelpBox:    tview.NewTextView(),
	}

	sts := bpz.LoadJsonfile()
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}

	pn := NewProjectNode(&bpzui)

	bpzui.ProjectBox = pn.Tree
	bpzui.FormBox = pn.GetForm()
	bpzui.HelpBox = pn.Help

	bpzui.RightHandView = bpzui.NewRightHandView()

	bpzui.FullView = tview.NewFlex().
		AddItem(bpzui.ProjectBox, 0, GoldenNarrow, true).
		AddItem(bpzui.RightHandView, 0, GoldenWide, false)

	app.SetRoot(bpzui.FullView, true)

	var exitingForm bool
	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for range only.Once {
			if event == nil {
				break
			}

			if event.Key() != tcell.KeyEsc {
				break
			}

			_, ok := app.GetFocus().(tview.FormItem)
			if ok {
				exitingForm = true
				event = nil
				app.SetFocus(bpzui.ProjectBox)
				break
			}

			switch app.GetFocus() {
			case bpzui.FormBox:
				app.SetFocus(bpzui.ProjectBox)
				break

			case bpzui.ProjectBox:
				if exitingForm {
					exitingForm = false
					break
				}
				app.Stop()
				break

			}
		}
		return event
	})

	return &bpzui
}

func (me *BpzUi) NewRightHandView() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(me.FormBox, 0, GoldenWide, false).
		AddItem(me.HelpBox, 0, GoldenNarrow, false)
}

func (me *BpzUi) Run() (sts Status) {
	err := me.App.Run()
	if err != nil {
		sts = status.Wrap(err)
	}
	return sts
}

func (me *BpzUi) MakeNodeView() (form *tview.Box, sts Status) {
	return tview.NewBox().SetBorder(true).SetTitle("Node"), nil
}
