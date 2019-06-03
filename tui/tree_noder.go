package tui

import (
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
	SetForm(*tview.Form)
	GetHelp() *tview.TextView
}
