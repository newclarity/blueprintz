package browseui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"blueprintz/wordpress"
	"github.com/rivo/tview"
)

var NilCoreView = (*CoreView)(nil)
var _ tui.Viewer = NilCoreView

type CoreView struct {
	*BaseView
}

func NewCoreView(ui *BrowseUi) *CoreView {
	cn := &CoreView{
		BaseView: NewBaseView(ui),
	}
	cn.Embedder = cn
	return cn
}

func (me *CoreView) GetForm() *tview.Form {

	core := me.Ui.Blueprintz.Core

	return me.Form.Clear(true).
		AddDropDown("Dialect:", global.AllDialects, global.AllDialects.Index(core.Dialect), nil).
		AddDropDown("Version:", wordpress.Versions, wordpress.Versions.Index(core.Version), nil)
}

func (me *CoreView) GetLabel() global.Label {
	return global.CoreLabel
}

func (me *CoreView) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
