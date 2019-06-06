package browseui

import (
	"blueprintz/global"
	"blueprintz/help"
	"blueprintz/tui"
	"github.com/rivo/tview"
)

var NilLayoutView = (*LayoutView)(nil)
var _ tui.Viewer = NilLayoutView

type LayoutView struct {
	*BaseView
}

func NewLayoutView(ui *BrowseUi) *LayoutView {
	ln := &LayoutView{
		BaseView: NewBaseView(ui),
	}
	ln.Embedder = ln
	return ln

}

func (me *LayoutView) GetForm() *tview.Form {
	layout := me.Ui.Blueprintz.Layout
	return me.Form.Clear(true).
		AddInputField("Project Path:", layout.ProjectPath, 15, nil, persistProjectPath).
		AddInputField("Webroot Path:", layout.WebrootPath, 25, nil, nil).
		AddInputField("Core Path:", layout.CorePath, 35, nil, nil).
		AddInputField("Content Path:", layout.ContentPath, 45, nil, nil)
}

func persistProjectPath(text string) {
}

func (me *LayoutView) GetLabel() global.Label {
	return global.LayoutLabel
}

func (me *LayoutView) GetHelpInfo() *help.Info {
	return &help.Info{
		Id:    global.LayoutHelpId,
		Label: global.LayoutLabel,
	}
}
