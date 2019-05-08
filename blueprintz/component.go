package blueprintz

import "blueprintz/global"

type Component struct {
	Name      global.ComponentName
	Version   global.Version
	SourceUrl global.Url
	LocalSlug global.Slug
}
