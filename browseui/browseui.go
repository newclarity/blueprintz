package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/tui"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
	"regexp"
	"strings"
)

type BrowseMode = string

const (
	TreeMode BrowseMode = "tree"
	FormMode BrowseMode = "form"
	ExitMode BrowseMode = "exit"
)

type BrowseUi struct {
	Blueprintz    *blueprintz.Blueprintz
	App           *tview.Application
	FullView      *tview.Flex
	FrameBox      *tview.Frame
	ProjectBox    *tview.TreeView
	RightHandView *tview.Flex
	FormBox       *tview.Form
	HelpBox       *tview.TextView
	Mode          BrowseMode
}

func New(bpz *blueprintz.Blueprintz) *BrowseUi {
	app := tview.NewApplication()
	browseui := BrowseUi{
		Blueprintz: bpz,
		App:        app,
		HelpBox:    tview.NewTextView(),
		Mode:       TreeMode,
	}

	pn := NewProjectTreeView(&browseui)

	browseui.ProjectBox = pn.Tree
	browseui.FormBox = pn.GetForm()

	browseui.HelpBox = pn.Help

	browseui.RightHandView = browseui.NewRightHandView()

	browseui.FullView = tview.NewFlex().
		AddItem(browseui.ProjectBox, 0, GoldenNarrow, true).
		AddItem(browseui.RightHandView, 0, GoldenWide, false)

	browseui.FrameBox = tview.NewFrame(browseui.FullView).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(global.AppName, true, tview.AlignCenter, tcell.ColorWhite).
		AddText(global.BrowseUiNavMenu, false, tview.AlignCenter, tcell.ColorLightSteelBlue)

	app.SetRoot(browseui.FrameBox, true)

	browseui.ShowHelp("")

	var exitingForm bool
	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for range only.Once {

			if event == nil {
				break
			}
			switch event.Key() {
			case tcell.KeyEsc:

				_, ok := app.GetFocus().(tview.FormItem)
				if ok {
					exitingForm = true
					event = nil
					browseui.ShowHelp("")
					app.SetFocus(browseui.ProjectBox)
					break
				}

				switch app.GetFocus() {
				case browseui.FormBox:
					app.SetFocus(browseui.ProjectBox)
					browseui.Mode = TreeMode
					break

				case browseui.ProjectBox:
					if exitingForm {
						exitingForm = false
						break
					}
					browseui.Mode = ExitMode
					app.Stop()
					break

				}
			case tcell.KeyEnter:
				browseui.Mode = FormMode
				break

			default:
				if browseui.Mode == FormMode {
					break
				}

				// Clear the form; It will be redrawn the
				// next node selected in the tree has a form
				if browseui.FormBox == nil {
					browseui.FormBox = tview.NewForm()
				}
				formatBox(browseui.FormBox.Box, "Form View")
				browseui.FormBox.Clear(true)

			}
		}
		return event
	})

	return &browseui
}

func (me *BrowseUi) MakeNewForm(label global.Label) *tview.Form {
	form := tview.NewForm()
	return me.FormatForm(form, label)
}

func (me *BrowseUi) FormatForm(form *tview.Form, label string) *tview.Form {
	formatBox(form.Box, label)
	form.SetChangedFunc(func(item tview.FormItem) {
		me.ShowHelp(item.GetLabel())
	})
	return form
}

func (me *BrowseUi) CurrentNodeLabel() global.Label {
	return me.ProjectBox.GetCurrentNode().GetText()
}

var spaceRegexp = regexp.MustCompile(`\s+`)
var punctRegexp = regexp.MustCompile(`([:?])`)

func (me *BrowseUi) ShowHelp(field global.Label) {

	var helptext string

	node := me.ProjectBox.GetCurrentNode()

	ref, ok := node.GetReference().(tui.Viewer)
	if ok {
		helptext = ref.GetHelpId()
	} else {
		helptext = node.GetText()
	}

	if field != "" {
		// Remove spaces, colons and question marks
		field = spaceRegexp.ReplaceAllString(field, "_")
		field = punctRegexp.ReplaceAllString(field, "")
		helptext = fmt.Sprintf("%s:%s", helptext, field)
	}

	formatBox(me.HelpBox.Box, "Help")

	me.HelpBox.SetText(strings.ToLower(helptext))
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

func (me *BrowseUi) MakeFormView() (form *tview.Box, sts Status) {
	return tview.NewBox().SetBorder(true).SetTitle("Form"), nil
}

var externalOptions = global.YesNos{
	global.UnsetState,
	global.YesState,
	global.NoState,
}

func (me *BrowseUi) AddComponentFormFields(form *tview.Form, c jsonfile.Componenter) *tview.Form {
	name := fmt.Sprintf("%s Name", strings.Title(c.GetType()))
	return form.Clear(true).
		AddInputField(name, c.GetName(), 65, nil, nil).
		AddInputField("Version:", c.GetVersion(), 16, nil, nil).
		AddInputField("Subdir/Slug:", c.GetSubdir(), 30, nil, nil).
		AddInputField("Main file:", c.GetBasefile(), 30, nil, nil).
		AddInputField("Website:", c.GetWebsite(), 60, nil, nil).
		AddInputField("Download URL:", c.GetDownloadUrl(), 80, nil, nil).
		AddDropDown("External?", externalOptions, externalOptions.Index(c.GetExternal()), nil)

}
