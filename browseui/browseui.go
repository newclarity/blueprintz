package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
	"strings"
)

type BrowseUi struct {
	Blueprintz    *blueprintz.Blueprintz
	App           *tview.Application
	FullView      *tview.Flex
	ProjectBox    *tview.TreeView
	RightHandView *tview.Flex
	FormBox       *tview.Form
	HelpBox       *tview.TextView
}

func New(bpz *blueprintz.Blueprintz) *BrowseUi {
	app := tview.NewApplication()
	browseui := BrowseUi{
		Blueprintz: bpz,
		App:        app,
		HelpBox:    tview.NewTextView(),
	}

	sts := bpz.LoadJsonfile()
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}

	pn := NewProjectNode(&browseui)

	browseui.ProjectBox = pn.Tree
	browseui.FormBox = pn.GetForm()
	browseui.HelpBox = pn.Help

	browseui.RightHandView = browseui.NewRightHandView()

	browseui.FullView = tview.NewFlex().
		AddItem(browseui.ProjectBox, 0, GoldenNarrow, true).
		AddItem(browseui.RightHandView, 0, GoldenWide, false)

	app.SetRoot(browseui.FullView, true)

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
				app.SetFocus(browseui.ProjectBox)
				break
			}

			switch app.GetFocus() {
			case browseui.FormBox:
				app.SetFocus(browseui.ProjectBox)
				break

			case browseui.ProjectBox:
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

	return &browseui
}

func (me *BrowseUi) NewRightHandView() *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(me.FormBox, 0, GoldenWide, false).
		AddItem(me.HelpBox, 0, GoldenNarrow, false)
}

func (me *BrowseUi) Run() (sts Status) {
	err := me.App.Run()
	if err != nil {
		sts = status.Wrap(err)
	}
	return sts
}

func (me *BrowseUi) MakeNodeView() (form *tview.Box, sts Status) {
	return tview.NewBox().SetBorder(true).SetTitle("Node"), nil
}

var externalOptions = global.YesNos{
	global.UnsetState,
	global.YesState,
	global.NoState,
}

func (me *BrowseUi) AddComponentFormFields(form *tview.Form, c jsonfile.Componenter) *tview.Form {
	name := fmt.Sprintf("%s Name", strings.Title(c.GetType()))
	return form.Clear(true).
		AddInputField(name, c.GetName(), 50, nil, nil).
		AddInputField("Version:", c.GetVersion(), 16, nil, nil).
		AddInputField("Subdir/Slug:", c.GetSubdir(), 30, nil, nil).
		AddInputField("Main file:", c.GetBasefile(), 30, nil, nil).
		AddInputField("Website:", c.GetWebsite(), 60, nil, nil).
		AddInputField("Download URL:", c.GetDownloadUrl(), 80, nil, nil).
		AddDropDown("External?", externalOptions, externalOptions.Index(c.GetExternal()), nil)

}
