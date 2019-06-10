package browseui

import (
	"blueprintz/global"
	"blueprintz/help"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
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

func (me *LayoutView) GetLabel() global.Label {
	return global.LayoutLabel
}

func (me *LayoutView) GetHelpInfo() *help.Info {
	return &help.Info{
		Id:    global.LayoutHelpId,
		Label: global.LayoutLabel,
	}
}

func (me *LayoutView) GetForm() *tview.Form {
	layout := me.Ui.Blueprintz.Layout

	return me.Form.
		Clear(true).
		AddFormItem(NewInputField(), &InputFieldArgs{
			Label:      "Project Path:",
			Text:       layout.ProjectPath,
			FieldWidth: 15,
			DoneFunc: func(key tcell.Key) {
				return
			},
		}).
		AddFormItem(NewInputField(), &InputFieldArgs{
			Label:      "Webroot Path:",
			Text:       layout.WebrootPath,
			FieldWidth: 25,
			DoneFunc: func(key tcell.Key) {
				return
			},
		}).
		AddFormItem(NewInputField(), &InputFieldArgs{
			Label:      "Core Path:",
			Text:       layout.CorePath,
			FieldWidth: 35,
			DoneFunc: func(key tcell.Key) {
				return
			},
		}).
		AddFormItem(NewInputField(), &InputFieldArgs{
			Label:      "Content Path:",
			Text:       layout.ContentPath,
			FieldWidth: 45,
			DoneFunc: func(key tcell.Key) {
				return
			},
		})

}
