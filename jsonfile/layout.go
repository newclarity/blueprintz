package jsonfile

import (
	"blueprintz/global"
)

type Layouter interface {
	GetProjectDir() global.RelativeDir
	GetWebrootDir() global.RelativeDir
	GetContentDir() global.RelativeDir
	GetPluginsDir() global.RelativeDir
	GetCoreDir() global.RelativeDir
}

type Layout struct {
	ProjectDir global.RelativeDir `json:"project_dir"`
	WebrootDir global.RelativeDir `json:"webroot_dir"`
	ContentDir global.RelativeDir `json:"content_dir"`
	CoreDir    global.RelativeDir `json:"core_dir"`
}

func NewLayout(layout Layouter) *Layout {
	return &Layout{
		ProjectDir: layout.GetProjectDir(),
		WebrootDir: layout.GetWebrootDir(),
		ContentDir: layout.GetContentDir(),
		CoreDir:    layout.GetCoreDir(),
	}
}
