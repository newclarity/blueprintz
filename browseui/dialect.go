package browseui

import "blueprintz/global"

type Dialect struct {
	Name     global.DialectName
	Releases DialectReleases
}

type DialectReleases []*DialectRelease
type DialectRelease struct {
	Version global.Version
}

func (me *Dialect) GetReleases() DialectReleases {
	return DialectReleases{}
}
