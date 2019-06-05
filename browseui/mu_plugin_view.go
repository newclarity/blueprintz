package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilMuPluginView = (*MuPluginView)(nil)
var _ tui.Viewer = NilMuPluginView

type MuPluginView struct {
	*BaseView
	Plugin *blueprintz.Plugin
}

func NewMuPluginView(ui *BrowseUi, p *blueprintz.Plugin) *MuPluginView {
	mpn := &MuPluginView{
		BaseView: NewBaseView(ui),
		Plugin:   p,
	}
	mpn.Embedder = mpn
	return mpn

}

func (me *MuPluginView) GetForm() *tview.Form {
	return me.Ui.AddComponentFormFields(me.Form, me.Plugin)
}

func (me *MuPluginView) GetLabel() global.Label {
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

func (me *MuPluginView) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *MuPluginView) GetHelpId() global.Slug {
	return global.MuPluginHelpId
}
