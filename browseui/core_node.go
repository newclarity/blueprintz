package browseui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"blueprintz/wordpress"
	"github.com/rivo/tview"
)

var NilCoreNode = (*CoreNode)(nil)
var _ tui.TreeNoder = NilCoreNode

type CoreNode struct {
	*BaseNode
}

func NewCoreNode(ui *BrowseUi) *CoreNode {
	cn := &CoreNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz.Core),
	}
	cn.Embedder = cn
	return cn
}

func (me *CoreNode) GetForm() *tview.Form {
	return me.Form.Clear(true).
		AddDropDown("WordPress Dialect:", []string{"wordpress", "classicpress"}, 0, nil).
		AddDropDown("Dialect Version:", wordpress.Versions, 0, nil)
}

func (me *CoreNode) GetLabel() global.Label {
	return global.CoreNode
}

func (me *CoreNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
