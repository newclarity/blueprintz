package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"sort"
)

var NilPluginsNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilPluginsNode

type PluginsNode struct {
	Parent  *BpzUi
	Plugins blueprintz.Plugins
}

func NewPluginsNode(parent *BpzUi) *PluginsNode {
	return &PluginsNode{
		Parent:  parent,
		Plugins: parent.Blueprintz.Plugins,
	}
}

func (me *PluginsNode) GetForm() *tview.Form {
	return nil
}

func (me *PluginsNode) GetLabel() global.Label {
	return global.PluginsNode
}

func (me *PluginsNode) GetReference() interface{} {
	return me
}

func (me *PluginsNode) IsSelectable() bool {
	return true
}

func (me *PluginsNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *PluginsNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.Plugins))
	for i, tn := range me.Plugins {
		tns[i] = NewPluginNode(me.Parent, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}
