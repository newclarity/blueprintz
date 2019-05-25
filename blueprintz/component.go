package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"fmt"
)

var NilComponent = (*Component)(nil)
var _ jsonfile.Componenter = NilComponent
var _ recognize.Componenter = NilComponent

type Component struct {
	Version     global.Version
	Website     global.Url
	Subdir      global.Slug
	HeaderFile  global.Dir
	Source      global.Source
	Maintainer  global.Maintainer
	DownloadUrl global.Url
}

var panicMsg = "Cannot call %s() of blueprintz.Component; use blueprintz.Plugin or blueprintz.Theme instead."

func (me *Component) GetType() global.ComponentType {
	panic(fmt.Sprintf(panicMsg, "GetType"))
}

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

func (me *Component) GetSource() global.Source {
	return me.Source
}

func (me *Component) GetDownloadUrl() global.Url {
	return me.DownloadUrl
}

func (me *Component) GetMaintainer() global.Maintainer {
	return me.Maintainer
}
