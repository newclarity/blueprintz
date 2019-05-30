package tui

import (
	"blueprintz/global"
	"github.com/gdamore/tcell"
)

type TreeNoders []TreeNoder
type TreeNoder interface {
	GetLabel() global.NodeLabel
	GetReference() interface{}
	IsSelectable() bool
	GetColor() tcell.Color
	GetChildren() TreeNoders
}

var _ TreeNoder = (*NilTreeNoder)(nil)

type NilTreeNoder struct{}

func (me NilTreeNoder) GetLabel() global.NodeLabel {
	return global.UnknownNode
}

func (me NilTreeNoder) GetReference() interface{} {
	return me
}

func (me NilTreeNoder) IsSelectable() bool {
	return false
}

func (me NilTreeNoder) GetColor() tcell.Color {
	return tcell.ColorOrangeRed
}

func (me NilTreeNoder) GetChildren() TreeNoders {
	return nil
}
