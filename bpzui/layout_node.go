package bpzui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/rivo/tview"
)

var NilLayoutNode = (*LayoutNode)(nil)
var _ tui.TreeNoder = NilLayoutNode

type LayoutNode struct {
	*BaseNode
}

func NewLayoutNode(ui *BpzUi) *LayoutNode {
	ln := &LayoutNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz.Core),
	}
	ln.Embedder = ln
	return ln

}

func (me *LayoutNode) GetForm() *tview.Form {
	return me.Form.Clear(true).
		AddInputField("Project Path:", "", 15, nil, nil).
		AddInputField("Webroot Path:", "", 25, nil, nil).
		AddInputField("Core Path:", "", 35, nil, nil).
		AddInputField("Content Path:", "", 45, nil, nil)
}

func (me *LayoutNode) GetLabel() global.Label {
	return global.LayoutNode
}

func (me *LayoutNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
