package tui

import (
	"blueprintz/global"
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
	GetHelpId() global.Slug
}
