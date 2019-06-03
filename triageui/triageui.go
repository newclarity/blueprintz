package triageui

import (
	"blueprintz/blueprintz"
	"github.com/gearboxworks/go-status"
	"github.com/rivo/tview"
)

type TriageUi struct {
	Blueprintz *blueprintz.Blueprintz
	App        *tview.Application
}

func New(bpz *blueprintz.Blueprintz) *TriageUi {
	app := tview.NewApplication()
	triageui := TriageUi{
		Blueprintz: bpz,
		App:        app,
	}
	return &triageui
}

func (me *TriageUi) Run() (sts Status) {
	err := me.App.Run()
	if err != nil {
		sts = status.Wrap(err)
	}
	return sts
}
