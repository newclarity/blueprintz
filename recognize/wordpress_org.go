package recognize

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"regexp"
	"strings"
)

var NilWordPressOrg = (*WordPressOrg)(nil)
var _ Recognizer = NilWordPressOrg

const ComponentUriSlugVar = "{component}"
const ComponentUriTypeVar = "{type}"
const ComponentUriVersionVar = "{version}"
const ComponentDownloadUriTemplate = "https://downloads.wordpress.org/{type}/{component}.{version}.zip"
const ComponentRepoUrlRegex = "^https?://wordpress.org/{type}s/([^/]+)/?"

type WordPressOrg struct {
	R
}

func NewWordPressOrg() *WordPressOrg {
	return &WordPressOrg{}
}

func (me *WordPressOrg) ValidTypes() (cts global.ComponentTypes) {
	return global.ComponentTypes{
		global.PluginComponent,
		global.ThemeComponent,
	}
}

func (me *WordPressOrg) Matches(c Componenter) (match bool) {
	for range only.Once {
		match = true
		regex := strings.ReplaceAll(ComponentRepoUrlRegex, ComponentUriTypeVar, c.GetType())
		re := regexp.MustCompile(regex)
		if re.MatchString(c.GetWebsite()) {
			break
		}
		sts := VerifyUrl(me.GetDownloadUrl(c))
		status.Log(sts)
		if is.Success(sts) {
			break
		}
		match = false
	}
	return match

}

func (me *WordPressOrg) GetDownloadUrl(c Componenter) (url global.Url) {
	url = ComponentDownloadUriTemplate
	url = strings.ReplaceAll(url, ComponentUriTypeVar, c.GetType())
	url = strings.ReplaceAll(url, ComponentUriSlugVar, c.GetSlug())
	url = strings.ReplaceAll(url, ComponentUriVersionVar, c.GetVersion())
	return url
}
