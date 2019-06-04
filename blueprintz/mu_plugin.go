package blueprintz

import (
	"blueprintz/global"
	"blueprintz/tui"
)

var NilMuPlugins = (*MuPlugins)(nil)
var _ tui.Labeler = NilMuPlugins

type MuPlugins Plugins

func (me MuPlugins) GetLabel() global.Label {
	return global.MuPluginsLabel
}

func (me *MuPlugins) Scandir(path global.Path, allowHeaderless bool) (sts Status) {
	mp := Plugins(*me)
	sts = (&mp).Scandir(path, allowHeaderless)
	*me = MuPlugins(mp)
	return sts
}

func (me MuPlugins) FindPlugin(pn global.ComponentName) *Plugin {
	var plugin *Plugin
	for _, p := range me {
		if p.PluginName != pn {
			continue
		}
		plugin = p
	}
	return plugin
}
