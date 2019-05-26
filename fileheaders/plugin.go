package fileheaders

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"os"
	"strings"
)

var NilPlugin = (*Plugin)(nil)
var _ Componenter = NilPlugin
var _ jsonfile.Componenter = NilPlugin

type Plugins []*Plugin
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

func (me *Plugin) GetSlug() global.Slug {
	return me.GetSubdir()
}

func (me *Plugin) GetSubdir() global.Slug {
	subdir := me.Component.GetSubdir()
	if subdir == me.GetType()+"s" {
		subdir = ""
	}
	return subdir
}

func (me *Plugin) GetType() (ct global.ComponentType) {
	parts := strings.Split(me.Filepath, string(os.PathSeparator))
	for _, i := range []int{3, 2} {
		if len(parts) <= 2 {
			ct = global.UnknownComponent
		}
		try := strings.TrimRight(parts[len(parts)-i], "s")
		for _, _ct := range []global.ComponentType{global.PluginComponent, global.MuPluginComponent} {
			if try != _ct {
				continue
			}
			ct = _ct
			break
		}
	}
	return ct
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
