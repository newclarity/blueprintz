package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilPluginView = (*ProjectTreeView)(nil)
var _ tui.Viewer = NilPluginView

type PluginView struct {
	*BaseView
	Plugin *blueprintz.Plugin
}

func NewPluginView(ui *BrowseUi, p *blueprintz.Plugin) *PluginView {
	pn := &PluginView{
		BaseView: NewBaseView(ui),
		Plugin:   p,
	}
	pn.Embedder = pn
	return pn

}

func (me *PluginView) GetForm() *tview.Form {
	return me.Ui.AddComponentFormFields(me.Form, me.Plugin)
}

func (me *PluginView) GetLabel() global.Label {
	var label global.Label
	for range only.Once {
		plugin := me.Plugin
		if plugin.PluginName != "" {
			label = AddComponentVersion(plugin.Component, plugin.PluginName)
			break
		}
		label = GetComponentLabel(plugin.Component)
	}
	return label
}

func (me *PluginView) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *PluginView) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
