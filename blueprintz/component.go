package blueprintz

import (
	"blueprintz/agent"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"fmt"
)

var NilComponent = (*Component)(nil)
var _ jsonfile.Componenter = NilComponent
var _ agent.Componenter = NilComponent

type Component struct {
	Version    global.Version
	Website    global.Url
	SourceUrl  global.Url
	Subdir     global.Slug
	HeaderFile global.Dir
}

var panicMsg = "Cannot %s() of blueprintz.Component; use blueprintz.Plugin or blueprintz.Theme instead."

func (me *Component) GetName() global.ComponentName {
	panic(fmt.Sprintf(panicMsg, "GetName"))
}

func (me *Component) GetVersion() global.Version {
	return me.Version
}

func (me *Component) GetSubdir() global.Slug {
	return me.Subdir
}

func (me *Component) GetWebsite() global.Url {
	return me.Website
}

func (me *Component) GetSlug() global.Slug {
	return me.Subdir
}
