package blueprintz

import (
	"blueprintz/global"
)

type Component struct {
	Version    global.Version
	SourceUrl  global.Url
	LocalSlug  global.Slug
	HeaderFile global.Dir
}
