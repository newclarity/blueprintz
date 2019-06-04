package jsonfile

import (
	"blueprintz/global"
)

type Componenter interface {
	GetName() global.ComponentName
	GetType() global.ComponentType
	GetVersion() global.Version
	GetSubdir() global.Slug
	GetBasefile() global.Basefile
	GetWebsite() global.Url
	GetExternal() global.YesNo
	GetDownloadUrl() global.Url
}

type Component struct {
	Name        global.ComponentName `json:"name,omitempty"`
	Website     global.Url           `json:"website,omitempty"`
	Version     global.Version       `json:"version,omitempty"`
	Subdir      global.Slug          `json:"subdir,omitempty"`
	Basefile    global.Basefile      `json:"mainfile"`
	DownloadUrl global.Url           `json:"download,omitempty"`
	KeepCopy    global.YesNo         `json:"keepcopy,omitempty"`
	External    global.YesNo         `json:"external,omitempty"`
	Purchased   global.YesNo         `json:"purchased,omitempty"`
	Develop     global.YesNo         `json:"develop,omitempty"`
}

func NewComponent(n global.ComponentName, c Componenter) *Component {
	return &Component{
		Name:        n,
		Version:     c.GetVersion(),
		Subdir:      c.GetSubdir(),
		Basefile:    c.GetBasefile(),
		Website:     c.GetWebsite(),
		External:    c.GetExternal(),
		DownloadUrl: c.GetDownloadUrl(),
	}
}
