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
var _ ComponentRecognizer = NilWordPressOrg

type WordPressOrg struct {
	R
}

func NewWordPressOrg() *WordPressOrg {
	return &WordPressOrg{}
}

func (me *WordPressOrg) ValidComponentTypes() (cts global.ComponentTypes) {
	return global.ComponentTypes{
		global.PluginComponent,
		global.ThemeComponent,
	}
}

func (me *WordPressOrg) MatchesComponent(c Componenter) (match bool) {
	for range only.Once {
		match = true
		regex := strings.ReplaceAll(WpComponentRepoUrlRegex, UriTypeVar, c.GetType())
		re := regexp.MustCompile(regex)
		if re.MatchString(c.GetWebsite()) {
			break
		}
		sts := VerifyUrl(me.GetComponentDownloadUrl(c))
		status.Log(sts)
		if is.Success(sts) {
			break
		}
		match = false
	}
	return match

}

const wpDownloadUrlTemplate = "https://wordpress.org/wordpress-{version}.zip"

func (me *WordPressOrg) GetCoreDownloadUrl(c Coreer) (url global.Url) {
	return strings.ReplaceAll(
		wpDownloadUrlTemplate,
		UriVersionVar,
		c.GetVersion(),
	)
}

func (me *WordPressOrg) GetComponentDownloadUrl(c Componenter) (url global.Url) {
	url = WpComponentDownloadUriTemplate
	url = strings.ReplaceAll(url, UriTypeVar, c.GetType())
	url = strings.ReplaceAll(url, UriComponentVar, c.GetSlug())
	url = strings.ReplaceAll(url, UriVersionVar, c.GetVersion())
	return url
}
