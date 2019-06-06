package help

import (
	"blueprintz/global"
	"fmt"
)

var TextMap = map[global.HelpId]string{

	global.ProjectHelpId: fmt.Sprintf("A %s'Project'%s is a directory containing source code for "+
		"a website implemented using a dialect of WordPress (see %sCore%s for more info on "+
		"%sdialect%s.)\n\nBlueprintz %shighly%s recommends locating your actual website source code "+
		"in a subdirectory one level from the project root, e.g. %s/www%s; see %sLayout%s for more "+
		"info on recommended site directory layout.",
		HelpHighlight, HelpColor,
		HelpEmphasis, HelpColor,
		HelpEmphasis, HelpColor,
		HelpEmphasis, HelpColor,
		HelpEmphasis, HelpColor,
		HelpEmphasis, HelpColor,
	),

	j(global.ProjectHelpId, "project_name"): fmt.Sprintf("Provide a short %s'Project Name'%s here "+
		"to allow Blueprintz to communicate info about this project back to you. This name "+
		"is also helpful to use to communication information about this project among your "+
		"team members.",
		HelpHighlight, HelpColor,
	),

	j(global.ProjectHelpId, "description"): fmt.Sprintf("Provide a longer %s'Description'%s here "+
		"to provide a reader — such as a new team member — to know exactly how this "+
		"WordPress dialect website project is being used.",
		HelpHighlight, HelpColor,
	),
	j(global.ProjectHelpId, "blueprint_type"): fmt.Sprintf("Pick the type of '%sBlueprint%s' you are working with. Currently only '%swebsite%s' is supported.",
		HelpHighlight, HelpColor,
		HelpEmphasis, HelpColor,
	),

	global.CoreHelpId: "",
	j(global.CoreHelpId, "dialect"): fmt.Sprintf("A %s'DialectName'%s of WordPress is a specific %sdistrbution%s of WordPress. "+
		"The default dialect is simply %swordpress%s but certain hosts that offer managed WordPress have their "+
		"own dialects. A %sfork%s of WordPress — such as %sClassicPress%s — is also a considered a dialect by Blueprintz.",
		HelpHighlight, HelpColor,
		HelpEmphasis, HelpColor,
		HelpHighlight, HelpColor,
		HelpEmphasis, HelpColor,
		HelpEmphasis, HelpColor,
	),

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

	global.ComponentsHelpId: "",

	j(global.ComponentHelpId, "name"):     "",
	j(global.ComponentHelpId, "website"):  "",
	j(global.ComponentHelpId, "version"):  "",
	j(global.ComponentHelpId, "subdir"):   "",
	j(global.ComponentHelpId, "mainfile"): "",
	j(global.ComponentHelpId, "download"): "",
	j(global.ComponentHelpId, "external"): "",
}

func j(a, b string) string {
	return fmt.Sprintf("%s:%s", a, b)
}
