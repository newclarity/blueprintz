package blueprintz

import (
	"blueprintz/global"
	"regexp"
)

type CoreVersionFinders []*CoreVersionFinder
type CoreVersionFinder struct {
	CoreType global.CoreType
	Regexp   *regexp.Regexp
}
type (
	vfs = CoreVersionFinders
	vf  = CoreVersionFinder
)

var coreVersionFinders = vfs{
	&vf{
		CoreType: global.ClassicPressCore,
		Regexp:   regexp.MustCompile(`^\$cp_version\s*=\s*['"](.+)['"];\s*$`),
	},
	&vf{
		CoreType: global.WordPressCore,
		Regexp:   regexp.MustCompile(`^\$wp_version\s*=\s*['"](.+)['"];\s*$`),
	},
}
