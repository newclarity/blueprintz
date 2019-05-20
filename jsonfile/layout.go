package jsonfile

import (
	"blueprintz/global"
)

type Layouter interface {
	GetCorePath() global.Path
	GetProjectPath() global.Path
	GetWebrootPath() global.Path
	GetContentPath() global.Path
	GetPluginsPath() global.Path
	GetVendorPath() global.Path
}

type Layout struct {
	CorePath    global.Path `json:"core_path"`
	ProjectPath global.Path `json:"project_path"`
	WebrootPath global.Path `json:"webroot_path"`
	ContentPath global.Path `json:"content_path"`
	VendorPath  global.Path `json:"vendor_path,omitempty"`
}

func NewLayout() *Layout {
	return &Layout{}
}

func NewLayoutFromLayouter(layout Layouter) *Layout {
	return &Layout{
		ProjectPath: layout.GetProjectPath(),
		WebrootPath: layout.GetWebrootPath(),
		ContentPath: layout.GetContentPath(),
		VendorPath:  layout.GetVendorPath(),
		CorePath:    layout.GetCorePath(),
	}
}
