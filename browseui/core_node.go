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

var dialects = global.Dialects{
	global.WordPressDialect,
	global.ClassicPressDialect,
	global.PantheonWPDialect,
}

func (me *CoreNode) GetForm() *tview.Form {

	core := me.Ui.Blueprintz.Core

	return me.Form.Clear(true).
		AddDropDown("Dialect:", dialects, dialects.Index(core.Dialect), nil).
		AddDropDown("Version:", wordpress.Versions, wordpress.Versions.Index(core.Version), nil)
}

func (me *CoreNode) GetLabel() global.Label {
	return global.CoreLabel
}

func (me *CoreNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
