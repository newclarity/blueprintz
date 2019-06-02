package blueprintz

import (
	"blueprintz/global"
	"blueprintz/tui"
)

var NilMuPlugins = (*MuPlugins)(nil)
var _ tui.Labeler = NilMuPlugins

type MuPlugins Plugins

func (me MuPlugins) GetLabel() global.Label {
	return global.MuPluginsNode
}

func (me *MuPlugins) Scandir(path global.Path, allowHeaderless bool) (sts Status) {
	mp := Plugins(*me)
	sts = (&mp).Scandir(path, allowHeaderless)
	*me = MuPlugins(mp)
	return sts
}
