package blueprintz

import (
	"blueprintz/global"
	"regexp"
)

type CoreVersionFinders []*CoreVersionFinder
type CoreVersionFinder struct {
	CoreType global.Dialect
	Regexp   *regexp.Regexp
}
type (
	vfs = CoreVersionFinders
	vf  = CoreVersionFinder
)

var coreVersionFinders = vfs{
	&vf{
		CoreType: global.ClassicPressDialect,
		Regexp:   regexp.MustCompile(`^\$cp_version\s*=\s*['"](.+)['"];\s*$`),
	},
	&vf{
		CoreType: global.WordPressDialect,
		Regexp:   regexp.MustCompile(`^\$wp_version\s*=\s*['"](.+)['"];\s*$`),
	},
}
