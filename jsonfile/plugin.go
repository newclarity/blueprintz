package jsonfile

import (
	"blueprintz/global"
)

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

type Plugin struct {
	*Component
}

func NewPlugins() Plugins {
	return make(Plugins, 0)
}

func NewPlugin(plugin Componenter) *Plugin {
	return &Plugin{
		Component: NewComponent(plugin.GetName(), plugin),
	}
}

func (me *Plugins) Dedup() {
	pm := me.GetMap()
	*me = (*me)[0:0]
	for _, p := range pm {
		*me = append(*me, p)
	}
}

func (me *Plugins) GetMap() PluginMap {
	pm := make(PluginMap, 0)
	for _, p := range *me {
		pm[p.Subdir] = p
	}
	return pm
}
