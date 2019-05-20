package fileheaders

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilPlugin = (*Plugin)(nil)
var _ Componenter = NilPlugin
var _ jsonfile.Componenter = NilPlugin

type Plugin struct {
	PluginName global.ComponentName `fileheader:"Plugin Name"`
	PluginURI  global.Url           `fileheader:"Plugin URI"`
	Network    string               `fileheader:"Network"`
	*Component
}

func NewPlugin(fp global.Filepath) *Plugin {
	return &Plugin{
		Component: &Component{
			Filepath: fp,
		},
	}
}

func (me *Plugin) GetType() ComponenterType {
	return PluginComponenter
}

func (me *Plugin) GetHeaderValueFieldMap(...Componenter) HeaderValueFieldMap {
	return me.Component.GetHeaderValueFieldMap(me)
}

func (me *Plugin) GetName() global.ComponentName {
	return me.PluginName
}

func (me *Plugin) GetWebsite() global.Url {
	return me.PluginURI
}
