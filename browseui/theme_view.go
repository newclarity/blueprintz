package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilThemeView = (*ThemeView)(nil)
var _ tui.Viewer = NilThemeView

type ThemeView struct {
	*BaseView
	Theme *blueprintz.Theme
}

func NewThemeView(ui *BrowseUi, t *blueprintz.Theme) *ThemeView {
	tn := &ThemeView{
		BaseView: NewBaseView(ui),
		Theme:    t,
	}
	tn.Embedder = tn
	return tn
}

func (me *ThemeView) GetForm() *tview.Form {
	return me.Ui.AddComponentFormFields(me.Form, me.Theme)
}

func (me *ThemeView) SetForm(form *tview.Form) {
	me.Form = form
}

func (me *ThemeView) GetLabel() global.Label {
	var label global.Label
	for range only.Once {
		theme := me.Theme
		if theme.ThemeName != "" {
			label = AddComponentVersion(theme.Component, theme.ThemeName)
			break
		}
		label = GetComponentLabel(theme.Component)
	}
	return label
}

func (me *ThemeView) IsSelectable() bool {
	return true
}

func (me *ThemeView) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *ThemeView) GetChildren() tui.Viewers {
	return nil
}
func (me *ThemeView) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
