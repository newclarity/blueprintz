package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilThemeNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilThemeNode

type ThemeNode struct {
	Parent *BpzUi
	Theme  *blueprintz.Theme
	Form   *tview.Form
}

func NewThemeNode(parent *BpzUi, t *blueprintz.Theme) *ThemeNode {
	form := tview.NewForm()
	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	return &ThemeNode{
		Parent: parent,
		Theme:  t,
		Form:   form,
	}
}

func (me *ThemeNode) GetForm() *tview.Form {
	return nil
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

func (me *ThemeNode) GetReference() interface{} {
	return me
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
