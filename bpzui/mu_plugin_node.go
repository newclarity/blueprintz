package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

var NilMuPluginNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilMuPluginNode

type MuPluginNode struct {
	Parent *BpzUi
	Plugin *blueprintz.Plugin
	Form   *tview.Form
}

func NewMuPluginNode(parent *BpzUi, p *blueprintz.Plugin) *MuPluginNode {
	form := tview.NewForm()
	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	return &MuPluginNode{
		Parent: parent,
		Plugin: p,
		Form:   form,
	}
}

func (me *MuPluginNode) GetForm() *tview.Form {
	return nil
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

func (me *MuPluginNode) GetReference() interface{} {
	return me
}

func (me *MuPluginNode) IsSelectable() bool {
	return true
}

func (me *MuPluginNode) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *MuPluginNode) GetChildren() tui.TreeNoders {
	return nil
}
