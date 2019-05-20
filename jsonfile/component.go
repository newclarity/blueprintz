package jsonfile

import (
	"blueprintz/global"
)

type Componenter interface {
	GetName() global.ComponentName
	GetVersion() global.Version
	GetSubdir() global.Slug
	GetWebsite() global.Url
}

type Component struct {
	Name      global.ComponentName `json:"name"`
	Website   global.Url           `json:"website"`
	Version   global.Version       `json:"version"`
	Subdir    global.Slug          `json:"subdir"`
	SourceUrl global.Url           `json:"source,omitempty"`
}

func NewComponent(n global.ComponentName, c Componenter) *Component {
	return &Component{
		Name:    n,
		Version: c.GetVersion(),
		Subdir:  c.GetSubdir(),
		Website: c.GetWebsite(),
	}
}
