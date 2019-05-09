package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"fmt"
)

var NilComponent = (*Component)(nil)
var _ jsonfile.Componenter = NilComponent

type Component struct {
	Version    global.Version
	SourceUrl  global.Url
	LocalSlug  global.Slug
	HeaderFile global.Dir
}

var panicMsg = "Cannot %s() of blueprintz.Component; use blueprintz.Plugin or blueprintz.Theme instead."

func (me *Component) GetName() global.ComponentName {
	panic(fmt.Sprintf(panicMsg, "GetName"))
}

func (me *Component) GetVersion() global.Version {
	return me.Version
}

func (me *Component) GetLocalDir() global.Slug {
	return me.LocalSlug
}

func (me *Component) GetWebsite() global.Url {
	panic(fmt.Sprintf(panicMsg, "GetWebsite"))
}
