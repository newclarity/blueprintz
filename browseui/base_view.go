package browseui

import (
	"blueprintz/global"
	"blueprintz/help"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var NilBaseView = (*BaseView)(nil)
var _ tui.Viewer = NilBaseView

type BaseView struct {
	Ui       *BrowseUi
	Embedder tui.Viewer
	Form     *tview.Form
}

func NewBaseView(ui *BrowseUi) *BaseView {
	return &BaseView{
		Ui:   ui,
		Form: tview.NewForm(),
	}
}
func (me *BaseView) GetForm() *tview.Form {
	return nil
}

func (me *BaseView) SetForm(form *tview.Form) {
	me.Form = form
}

func (me *BaseView) GetLabel() global.Label {
	panic("Must implement GetLabel() in embedding struct")
}

func (me *BaseView) GetReference() interface{} {
	return me.Embedder
}

func (me *BaseView) IsSelectable() bool {
	return true
}

func (me *BaseView) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *BaseView) GetChildren() tui.Viewers {
	return nil
}
func (me *BaseView) GetHelpInfo() *help.Info {
	panic("Must implement GetHelpInfo() in embedding struct")
}
