package fileheaders

import "blueprintz/global"

var NilTheme = (*Theme)(nil)
var _ Componenter = NilTheme

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

func (me *Theme) GetHeaderFields(...Componenter) ValueMap {
	return me.Component.GetHeaderFields(me)
}
