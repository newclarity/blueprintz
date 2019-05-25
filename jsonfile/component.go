package jsonfile

import (
	"blueprintz/global"
)

type Componenter interface {
	GetName() global.ComponentName
	GetVersion() global.Version
	GetSubdir() global.Slug
	GetWebsite() global.Url
	GetSource() global.Source
	GetMaintainer() global.Maintainer
	GetDownloadUrl() global.Url
}

type Component struct {
	Name        global.ComponentName `json:"name"`
	Website     global.Url           `json:"website"`
	Version     global.Version       `json:"version"`
	Subdir      global.Slug          `json:"subdir"`
	DownloadUrl global.Url           `json:"download,omitempty"`
	Maintainer  global.Maintainer    `json:"maintainer,omitempty"`
	Source      global.Source        `json:"author_type,omitempty"`
}

func NewComponent(n global.ComponentName, c Componenter) *Component {
	return &Component{
		Name:        n,
		Version:     c.GetVersion(),
		Subdir:      c.GetSubdir(),
		Website:     c.GetWebsite(),
		Source:      c.GetSource(),
		Maintainer:  c.GetMaintainer(),
		DownloadUrl: c.GetDownloadUrl(),
	}
}
