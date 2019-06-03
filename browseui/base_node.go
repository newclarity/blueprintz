package browseui

import (
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var NilBaseNode = (*BaseNode)(nil)
var _ tui.TreeNoder = NilBaseNode

type BaseNode struct {
	Ui       *BpzUi
	Embedder tui.TreeNoder
	Data     NodeData
	Form     *tview.Form
}

func NewBaseNode(ui *BpzUi, data NodeData) *BaseNode {
	return &BaseNode{
		Ui:   ui,
		Data: data,
		Form: tview.NewForm(),
	}
}
func (me *BaseNode) GetForm() *tview.Form {
	return nil
}

func (me *BaseNode) SetForm(form *tview.Form) {
	me.Form = form
}

func (me *BaseNode) GetLabel() global.Label {
	panic("Must implement GetLabel() in embedding struct")
}

func (me *BaseNode) GetReference() interface{} {
	return me.Embedder
}

func (me *BaseNode) IsSelectable() bool {
	return true
}

func (me *BaseNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *BaseNode) GetChildren() tui.TreeNoders {
	return nil
}
func (me *BaseNode) GetHelp() *tview.TextView {
	panic("Must implement GetHelp() in embedding struct")
}
