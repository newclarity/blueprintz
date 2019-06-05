package browseui

import (
	"blueprintz/global"
	"blueprintz/helptext"
	"fmt"
	"github.com/gearboxworks/go-status/only"
	"strings"
)

func GetHelp(helpid global.HelpId) (help string) {
	for range only.Once {
		helpmap := helptext.TextMap
		helpid = strings.ToLower(helpid)
		parts := strings.Split(helpid+":", ":")

		var fh global.HelpId
		if helpid != parts[0] {
			fh, _ = helpmap[helpid]
		}
		nh, _ := helpmap[parts[0]]
		var chid global.HelpId
		switch parts[0] {
		case global.ThemesHelpId, global.PluginsHelpId, global.MuPluginsHelpId:
			chid = global.ComponentsHelpId
		case global.ThemeHelpId, global.PluginHelpId, global.MuPluginHelpId:
			chid = global.ComponentHelpId
		}
		ch, _ := helpmap[chid]
		fh = joinhelp(fh, ch)
		if fh != "" {
			help = fh
			break
		}
		if nh != "" {
			help = nh
			break
		}
		//help = joinhelp(fh, nh)
		//if help != "" {
		//	break
		//}
		help = fmt.Sprintf(
			"No helptext specified yet for this Help ID: '%s'",
			helpid,
		)
	}
	return help
}

func joinhelp(specific global.HelpId, general global.HelpId) (help string) {
	switch {
	case specific != "" && general != "":
		help = fmt.Sprintf("%s\n\n%s", specific, general)

	case general == "":
		help = specific

	default:
		help = general

	}
	return help
}

//func LoadHelp(c *config.Config) {
//	//cdir := c.OsBridge.GetUserConfigDir()
//
//}
