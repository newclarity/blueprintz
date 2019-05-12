package blueprintz

import (
	"blueprintz/courier"
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilTheme = (*Theme)(nil)
var _ jsonfile.Componenter = NilTheme
var _ courier.Componenter = NilTheme

type ThemeMap map[global.ComponentName]*Theme
type Themes []*Theme

type Theme struct {
	ThemeName global.ComponentName
	ThemeURI  global.Url
	*Component
}
