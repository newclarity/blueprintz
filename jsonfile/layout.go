package jsonfile

import (
	"blueprintz/global"
)

type Layouter interface {
	GetProjectDir() global.Path
	GetWebrootDir() global.Path
	GetContentDir() global.Path
	GetPluginsDir() global.Path
	GetCoreDir() global.Path
}

type Layout struct {
	ProjectDir global.Path `json:"project_dir"`
	WebrootDir global.Path `json:"webroot_dir"`
	ContentDir global.Path `json:"content_dir"`
	CoreDir    global.Path `json:"core_dir"`
}

func NewLayout(layout Layouter) *Layout {
	return &Layout{
		ProjectDir: layout.GetProjectDir(),
		WebrootDir: layout.GetWebrootDir(),
		ContentDir: layout.GetContentDir(),
		CoreDir:    layout.GetCoreDir(),
	}
}
