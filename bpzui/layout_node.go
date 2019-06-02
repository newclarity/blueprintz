package bpzui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var NilLayoutNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilLayoutNode

type LayoutNode struct {
	Parent   *BpzUi
	NodeData NodeData
	Form     *tview.Form
}

func NewLayoutNode(parent *BpzUi) *LayoutNode {
	form := tview.NewForm()
	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	return &LayoutNode{
		Parent:   parent,
		NodeData: parent.Blueprintz.Layout,
	}
}

func (me *LayoutNode) GetForm() *tview.Form {
	return me.Form
}

func (me *LayoutNode) GetLabel() global.Label {
	return global.LayoutNode
}

func (me *LayoutNode) GetReference() interface{} {
	return me
}

func (me *LayoutNode) IsSelectable() bool {
	return true
}

func (me *LayoutNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *LayoutNode) GetChildren() tui.TreeNoders {
	return nil

}

func (me *LayoutNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
