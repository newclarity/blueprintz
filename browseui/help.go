package browseui

import (
	"blueprintz/global"
	"blueprintz/help"
	"fmt"
	"github.com/gearboxworks/go-status/only"
	"strings"
)

func GetHelp(helpid global.HelpId) (h string) {
	for range only.Once {
		helpmap := help.TextMap
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
			h = fh
			break
		}
		if nh != "" {
			h = nh
			break
		}
		//h = joinhelp(fh, nh)
		//if h != "" {
		//	break
		//}
		h = fmt.Sprintf(
			"No help specified yet for this Help ID: '%s'",
			helpid,
		)
	}
	return h
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
