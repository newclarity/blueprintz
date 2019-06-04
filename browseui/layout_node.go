package browseui

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

func NewLayoutNode(ui *BrowseUi) *LayoutNode {
	ln := &LayoutNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz.Core),
	}
	ln.Embedder = ln
	return ln

}

func (me *LayoutNode) GetForm() *tview.Form {
	layout := me.Ui.Blueprintz.Layout
	return me.Form.Clear(true).
		AddInputField("Project Path:", layout.ProjectPath, 15, nil, nil).
		AddInputField("Webroot Path:", layout.WebrootPath, 25, nil, nil).
		AddInputField("Core Path:", layout.CorePath, 35, nil, nil).
		AddInputField("Content Path:", layout.ContentPath, 45, nil, nil)
}

func (me *LayoutNode) GetLabel() global.Label {
	return global.LayoutLabel
}

func (me *LayoutNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
