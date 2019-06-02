package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"sort"
)

var NilMuPluginsNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilMuPluginsNode

type MuPluginsNode struct {
	Parent    *BpzUi
	MuPlugins blueprintz.MuPlugins
}

func NewMuPluginsNode(parent *BpzUi) *MuPluginsNode {
	return &MuPluginsNode{
		Parent:    parent,
		MuPlugins: parent.Blueprintz.MuPlugins,
	}

}

func (me *MuPluginsNode) GetForm() *tview.Form {
	return nil
}

func (me *MuPluginsNode) GetLabel() global.Label {
	return global.MuPluginsNode
}

func (me *MuPluginsNode) GetReference() interface{} {
	return me
}

func (me *MuPluginsNode) IsSelectable() bool {
	return true
}

func (me *MuPluginsNode) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *MuPluginsNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.MuPlugins))
	for i, tn := range me.MuPlugins {
		tns[i] = NewMuPluginNode(me.Parent, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *MuPluginsNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
