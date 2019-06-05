package browseui

import (
	"blueprintz/config"
	"blueprintz/global"
	"fmt"
	"github.com/mitchellh/go-homedir"
)

func j(a, b string) string {
	return fmt.Sprintf("%s:%s", a, b)
}

var help = map[global.HelpIdType]string{

	global.ProjectHelpId:                      "",
	j(global.ProjectHelpId, "project_name"):   "",
	j(global.ProjectHelpId, "description"):    "",
	j(global.ProjectHelpId, "local_domain"):   "",
	j(global.ProjectHelpId, "blueprint_type"): "",

	global.CoreHelpId:               "",
	j(global.CoreHelpId, "dialect"): "",
	j(global.CoreHelpId, "version"): "",

	global.LayoutHelpId:                    "",
	j(global.LayoutHelpId, "core_path"):    "",
	j(global.LayoutHelpId, "project_path"): "",
	j(global.LayoutHelpId, "webroot_path"): "",
	j(global.LayoutHelpId, "content_path"): "",

	global.ThemesHelpId:    "",
	global.ThemeHelpId:     "",
	global.PluginsHelpId:   "",
	global.PluginHelpId:    "",
	global.MuPluginsHelpId: "",
	global.MuPluginHelpId:  "",

	j(global.ComponentHelpId, "name"):     "",
	j(global.ComponentHelpId, "website"):  "",
	j(global.ComponentHelpId, "version"):  "",
	j(global.ComponentHelpId, "subdir"):   "",
	j(global.ComponentHelpId, "mainfile"): "",
	j(global.ComponentHelpId, "download"): "",
	j(global.ComponentHelpId, "external"): "",
}

func LoadHelp(c *config.Config) {
	//cdir := c.OsBridge.GetUserConfigDir()

}
