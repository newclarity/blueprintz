package jsonfile

import (
	"blueprintz/global"
)

type Layouter interface {
	GetProjectPath() global.Path
	GetWebrootPath() global.Path
	GetContentPath() global.Path
	GetPluginsPath() global.Path
	GetCorePath() global.Path
}

type Layout struct {
	ProjectPath global.Path `json:"project_path"`
	WebrootPath global.Path `json:"webroot_path"`
	ContentPath global.Path `json:"content_path"`
	CorePath    global.Path `json:"core_path"`
}

func NewLayout(layout Layouter) *Layout {
	return &Layout{
		ProjectPath: layout.GetProjectPath(),
		WebrootPath: layout.GetWebrootPath(),
		ContentPath: layout.GetContentPath(),
		CorePath:    layout.GetCorePath(),
	}
}
