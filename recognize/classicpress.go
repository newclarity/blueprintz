package recognize

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-status/only"
	"strings"
)

var NilClassicPress = (*ClassicPress)(nil)
var _ ComponentRecognizer = NilClassicPress

type ClassicPress struct {
	WordPressOrg
}

const cpGitHubRootUrl = "https://github.com/ClassicPress"
const cpReleaseDownloadUrlTemplate = cpGitHubRootUrl + "/ClassicPress-release/archive/{version}.zip"
const cpPrereleaseDownloadUrlTemplate = cpGitHubRootUrl + "/ClassicPress/archive/{version}.zip"

func (me *ClassicPress) GetCoreDownloadUrl(c Coreer) (url global.Url) {
	ut := cpPrereleaseDownloadUrlTemplate
	v := c.GetVersion()
	for range only.Once {
		if !strings.Contains(v, "+") {
			break
		}
		if !strings.Contains(v, "-") {
			break
		}
		ut = cpReleaseDownloadUrlTemplate
	}
	return strings.ReplaceAll(ut, UriVersionVar, v)
}
