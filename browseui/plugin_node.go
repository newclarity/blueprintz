package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilPluginNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilPluginNode

type PluginNode struct {
	*BaseNode
	Plugin *blueprintz.Plugin
}

func NewPluginNode(ui *BrowseUi, p *blueprintz.Plugin) *PluginNode {
	pn := &PluginNode{
		BaseNode: NewBaseNode(ui, p),
		Plugin:   p,
	}
	pn.Embedder = pn
	return pn

}

func (me *PluginNode) GetForm() *tview.Form {
	return me.Ui.AddComponentFormFields(me.Form, me.Plugin)
}

func (me *PluginNode) GetLabel() global.Label {
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

func (me *PluginNode) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *PluginNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
