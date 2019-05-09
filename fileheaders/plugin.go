package fileheaders

import (
	"blueprintz/global"
)

var NilPlugin = (*Plugin)(nil)
var _ Componenter = NilPlugin

type Plugin struct {
	PluginName string     `fileheader:"Plugin Name"`
	PluginURI  global.Url `fileheader:"Plugin URI"`
	Network    string     `fileheader:"Network"`
	*Component
}

func NewPlugin(fp global.Filepath) *Plugin {
	return &Plugin{
		Component: &Component{
			Filepath: fp,
		},
	}
}

func (me *Plugin) GetHeaderFields(...Componenter) ValueMap {
	return me.Component.GetHeaderFields(me)
}
