package bpzui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"blueprintz/wordpress"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var NilCoreNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilCoreNode

type CoreNode struct {
	Parent   *BpzUi
	NodeData NodeData
	Form     *tview.Form
}

func NewCoreNode(parent *BpzUi) *CoreNode {
	form := tview.NewForm().
		AddDropDown("WordPress Dialect:", []string{"wordpress", "classicpress"}, 0, nil).
		AddDropDown("Dialect Version:", wordpress.Versions, 0, nil).
		AddButton("Cancel", func() {
			parent.App.SetFocus(parent.ProjectNode)
		}).
		AddButton("Save", func() {
			parent.App.SetFocus(parent.ProjectNode)
		})

	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	return &CoreNode{
		Parent:   parent,
		NodeData: parent.Blueprintz.Core,
		Form:     form,
	}
}
func (me *CoreNode) GetForm() *tview.Form {
	parent := me.Parent
	me.Form.Clear(true)
	me.Form.
		AddDropDown("WordPress Dialect:", []string{"wordpress", "classicpress"}, 0, nil).
		AddDropDown("Dialect Version:", wordpress.Versions, 0, nil).
		AddButton("Cancel", func() {
			parent.App.SetFocus(parent.ProjectNode)
		}).
		AddButton("Save", func() {
			parent.App.SetFocus(parent.ProjectNode)
		})
	return me.Form
}

func (me *CoreNode) GetLabel() global.Label {
	return global.CoreNode
}

func (me *CoreNode) GetReference() interface{} {
	return me
}

func (me *CoreNode) IsSelectable() bool {
	return true
}

func (me *CoreNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *CoreNode) GetChildren() tui.TreeNoders {
	return nil
}
func (me *CoreNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
