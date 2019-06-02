package bpzui

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
	Parent *BpzUi
	Plugin *blueprintz.Plugin
	Form   *tview.Form
}

func NewPluginNode(parent *BpzUi, p *blueprintz.Plugin) *PluginNode {
	form := tview.NewForm()
	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	return &PluginNode{
		Parent: parent,
		Plugin: p,
		Form:   form,
	}
}

func (me *PluginNode) GetForm() *tview.Form {
	return nil
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

//func (me *PluginNode) GetLabel() global.Label {
//	return me.Plugin.GetLabel()
//}

func (me *PluginNode) GetReference() interface{} {
	return me
}

func (me *PluginNode) IsSelectable() bool {
	return true
}

func (me *PluginNode) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *PluginNode) GetChildren() tui.TreeNoders {
	return nil
}
