package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilMuPluginNode = (*MuPluginNode)(nil)
var _ tui.TreeNoder = NilMuPluginNode

type MuPluginNode struct {
	*BaseNode
	Plugin *blueprintz.Plugin
}

func NewMuPluginNode(ui *BrowseUi, p *blueprintz.Plugin) *MuPluginNode {
	mpn := &MuPluginNode{
		BaseNode: NewBaseNode(ui, p),
		Plugin:   p,
	}
	mpn.Embedder = mpn
	return mpn

}

func (me *MuPluginNode) GetForm() *tview.Form {
	return addComponentFormFields(me.Form, "Plugin")
}

func (me *MuPluginNode) GetLabel() global.Label {
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

func (me *MuPluginNode) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *MuPluginNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
