package tui

import (
	"blueprintz/help"
	"github.com/rivo/tview"
)

type Viewers []Viewer
type Viewer interface {
	Labeler
	GetReference() interface{}
	IsSelectable() bool
	GetColor() Color
	GetChildren() Viewers
	GetForm() *tview.Form
	SetForm(*tview.Form)
	GetHelpInfo() *help.Info
}
