package tui

import (
	"blueprintz/global"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type TreeNoders []TreeNoder
type TreeNoder interface {
	Labeler
	GetReference() interface{}
	IsSelectable() bool
	GetColor() Color
	GetChildren() TreeNoders
	GetForm() *tview.Form
}

var _ TreeNoder = (*NilTreeNoder)(nil)

type NilTreeNoder struct{}

func (me NilTreeNoder) GetLabel() global.Label {
	return global.UnknownNode
}

func (me NilTreeNoder) GetReference() interface{} {
	return me
}

func (me NilTreeNoder) IsSelectable() bool {
	return false
}

func (me NilTreeNoder) GetColor() Color {
	return tcell.ColorOrangeRed
}

func (me NilTreeNoder) GetChildren() TreeNoders {
	return nil
}

func (me NilTreeNoder) GetForm() *tview.Form {
	return nil
}
