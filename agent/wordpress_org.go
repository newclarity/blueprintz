package agent

import (
	"blueprintz/global"
	"regexp"
	"strings"
)

const PluginUriPluginVar = "{plugin}"
const PluginUriVersionVar = "{version}"
const PluginDownloadUriTemplate = "https://downloads.wordpress.org/plugin/{plugin}.{version}.zip"
const PluginRepoUrlRegex = "^https?://wordpress.org/plugins/([^/]+)/?"

var NilWordPressOrg = (*WordPressOrg)(nil)
var _ Agenter = NilWordPressOrg

type WordPressOrg struct {
}

func NewWordPressOrg() *WordPressOrg {
	return &WordPressOrg{}
}
func (me *WordPressOrg) GetSourceUrl(c Componenter) (url global.Url) {
	url = strings.Replace(PluginDownloadUriTemplate, PluginUriPluginVar, c.GetSlug(), -1)
	url = strings.Replace(url, PluginUriVersionVar, c.GetVersion(), -1)
	return url
}

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(PluginRepoUrlRegex)
}

func (me *WordPressOrg) Match(args *Args) (match bool) {
	return re.MatchString(args.Website)
}
