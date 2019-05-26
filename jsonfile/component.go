package jsonfile

import (
	"blueprintz/global"
)

type Componenter interface {
	GetName() global.ComponentName
	GetVersion() global.Version
	GetSubdir() global.Slug
	GetBasefile() global.Basefile
	GetWebsite() global.Url
	GetSourceType() global.SourceType
	GetDownloadUrl() global.Url
}

type Component struct {
	Name        global.ComponentName `json:"name,omitempty"`
	Website     global.Url           `json:"website,omitempty"`
	Version     global.Version       `json:"version,omitempty"`
	Subdir      global.Slug          `json:"subdir,omitempty"`
	Basefile    global.Basefile      `json:"file"`
	DownloadUrl global.Url           `json:"download,omitempty"`
	SourceType  global.SourceType    `json:"type,omitempty"`
}

func NewComponent(n global.ComponentName, c Componenter) *Component {
	return &Component{
		Name:        n,
		Version:     c.GetVersion(),
		Subdir:      c.GetSubdir(),
		Basefile:    c.GetBasefile(),
		Website:     c.GetWebsite(),
		SourceType:  c.GetSourceType(),
		DownloadUrl: c.GetDownloadUrl(),
	}
}
