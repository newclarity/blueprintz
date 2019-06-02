package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"blueprintz/tui"
	"fmt"
	"github.com/gdamore/tcell"
)

var NilComponent = (*Component)(nil)
var _ jsonfile.Componenter = NilComponent
var _ recognize.Componenter = NilComponent

type Component struct {
	Version     global.Version
	Website     global.Url
	Subdir      global.Slug
	Basefile    global.Basefile
	HeaderFile  global.Dir
	DownloadUrl global.Url
	KeepCopy    global.YesNo
	External    global.YesNo
	Purchased   global.YesNo
	Develop     global.YesNo
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

func (me *Component) GetBasefile() global.Basefile {
	return me.Basefile
}

func (me *Component) GetDownloadUrl() global.Url {
	return me.DownloadUrl
}

func (me *Component) GetExternal() (ex global.YesNo) {
	ex = me.External
	if ex == "" && me.DownloadUrl != "" {
		ex = global.YesState
	}
	return ex
}

func (me *Component) IsSelectable() bool {
	return true
}

func (me *Component) GetColor() tui.Color {
	return tcell.ColorWhite
}

func (me *Component) GetChildren() tui.TreeNoders {
	return nil
}

func (me *Component) GetReference() interface{} {
	panic(fmt.Sprintf(panicMsg, "GetReference"))
}
