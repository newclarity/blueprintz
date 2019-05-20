package fileheaders

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilTheme = (*Theme)(nil)
var _ Componenter = NilTheme
var _ jsonfile.Componenter = NilTheme

type Theme struct {
	ThemeName string     `fileheader:"Theme Name"`
	ThemeURI  global.Url `fileheader:"Theme URI"`
	Tags      []string   `fileheader:"Tags"`
	*Component
}

func NewTheme(fp global.Filepath) *Theme {
	return &Theme{
		Component: &Component{
			Filepath: fp,
		},
	}
}

func (me *Theme) GetType() ComponenterType {
	return ThemeComponenter
}

func (me *Theme) GetHeaderValueFieldMap(...Componenter) HeaderValueFieldMap {
	return me.Component.GetHeaderValueFieldMap(me)
}

func (me *Theme) GetName() global.ComponentName {
	return me.ThemeName
}

func (me *Theme) GetWebsite() global.Url {
	return me.ThemeURI
}
