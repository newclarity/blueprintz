package jsonfile

import "blueprintz/global"

type ThemeMap map[global.ComponentName]*Theme
type Themes []*Theme

type Theme struct {
	*Component
}

func NewTheme(theme Componenter) *Theme {
	return &Theme{
		Component: NewComponent(theme),
	}
}
