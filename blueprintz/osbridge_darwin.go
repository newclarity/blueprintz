// +build darwin

package blueprintz

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-osbridge"
)

func GetOsBridge(project global.Name, userdata global.Path) *osbridge.OsBridge {
	return osbridge.NewOsBridge(&osbridge.Args{
		ProjectName:  project,
		UserDataPath: userdata,
		AdminPath:    NixAdminPath,
		ProjectDir:   MacOsProjectDir,
	})
}
