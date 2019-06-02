package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"sort"
)

var NilThemesNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilThemesNode

type ThemesNode struct {
	Parent *BpzUi
	Themes blueprintz.Themes
}

func NewThemesNode(parent *BpzUi) *ThemesNode {
	cn := ThemesNode{
		Parent: parent,
		Themes: parent.Blueprintz.Themes,
	}
	return &cn
}

func (me *ThemesNode) GetForm() *tview.Form {
	return nil
}

func (me *ThemesNode) GetLabel() global.Label {
	return global.ThemesNode
}

func (me *ThemesNode) GetReference() interface{} {
	return me
}

func (me *ThemesNode) IsSelectable() bool {
	return true
}

func (me *ThemesNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *ThemesNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.Themes))
	for i, tn := range me.Themes {
		tns[i] = NewThemeNode(me.Parent, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}
