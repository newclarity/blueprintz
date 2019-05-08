package jsonfile

import (
	"blueprintz/global"
)

type Layout struct {
	ProjectDir global.RelativeDir `json:"project_dir"`
	WebrootDir global.RelativeDir `json:"webroot_dir"`
	ContentDir global.RelativeDir `json:"content_dir"`
	CoreDir    global.RelativeDir `json:"core_dir"`
}

func NewLayout(layout global.Layouter) *Layout {
	return &Layout{
		ProjectDir: layout.GetProjectDir(),
		WebrootDir: layout.GetWebrootDir(),
		ContentDir: layout.GetContentDir(),
		CoreDir:    layout.GetCoreDir(),
	}
}
