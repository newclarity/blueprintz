package jsonfile

import (
	"blueprintz/global"
)

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

type Plugin struct {
	Component
}
