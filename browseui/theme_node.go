package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilThemeNode = (*ThemeNode)(nil)
var _ tui.TreeNoder = NilThemeNode

type ThemeNode struct {
	*BaseNode
	Theme *blueprintz.Theme
}

func NewThemeNode(ui *BrowseUi, t *blueprintz.Theme) *ThemeNode {
	tn := &ThemeNode{
		BaseNode: NewBaseNode(ui, t),
		Theme:    t,
	}
	tn.Embedder = tn
	return tn
}

func (me *ThemeNode) GetForm() *tview.Form {
	return me.Ui.AddComponentFormFields(me.Form, me.Theme)
}

func (me *ThemeNode) SetForm(form *tview.Form) {
	me.Form = form
}

func (me *ThemeNode) GetLabel() global.Label {
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

func (me *ThemeNode) IsSelectable() bool {
	return true
}

func (me *ThemeNode) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *ThemeNode) GetChildren() tui.TreeNoders {
	return nil
}
func (me *ThemeNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
