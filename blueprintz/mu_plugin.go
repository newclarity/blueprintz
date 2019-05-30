package blueprintz

import (
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
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

func (me MuPlugins) GetColor() tcell.Color {
	return tcell.ColorLime
}

func (me MuPlugins) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me))
	for i, tn := range me {
		tns[i] = tn
	}
	return tns
}

func (me *MuPlugins) Scandir(path global.Path, allowHeaderless bool) (sts Status) {
	ps := make(Plugins, len(*me))
	for i, mup := range *me {
		ps[i] = mup
	}
	return ps.Scandir(path, allowHeaderless)
}
