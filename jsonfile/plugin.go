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
