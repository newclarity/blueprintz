package blueprintz

import (
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"sort"
)

var NilMuPlugins = (*MuPlugins)(nil)
var _ tui.TreeNoder = NilMuPlugins

type MuPlugins Plugins

func (me MuPlugins) GetLabel() global.NodeLabel {
	return global.MuPluginsNode
}

func (me MuPlugins) GetReference() interface{} {
	return me
}

func (me MuPlugins) IsSelectable() bool {
	return true
}

func (me MuPlugins) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me MuPlugins) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me))
	for i, tn := range me {
		tns[i] = tn
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *MuPlugins) Scandir(path global.Path, allowHeaderless bool) (sts Status) {
	mp := Plugins(*me)
	sts = (&mp).Scandir(path, allowHeaderless)
	*me = MuPlugins(mp)
	return sts
}
