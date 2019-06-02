package tui

import (
	"blueprintz/global"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
)

type (
	Status = status.Status
	Color  = tcell.Color
)

type Labeler interface {
	GetLabel() global.Label
}
