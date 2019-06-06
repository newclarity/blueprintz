package browseui

import (
	"blueprintz/global"
	"blueprintz/help"
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
		AddDropDown("DialectName:",
			global.AllDialects,
			global.AllDialects.Index(core.Dialect),
			func(option string, optionIndex int) {

			}).
		AddDropDown("Version:", wordpress.Versions, wordpress.Versions.Index(core.Version), nil)
}

func (me *CoreView) GetLabel() global.Label {
	return global.CoreLabel
}

func (me *CoreView) GetHelpInfo() *help.Info {
	return &help.Info{
		Id:    global.CoreHelpId,
		Label: global.CoreLabel,
	}
}

func (me *CoreView) dialectSelected(option string, optionIndex int) {

}
