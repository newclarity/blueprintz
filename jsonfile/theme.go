package jsonfile

import "blueprintz/global"

type ThemeMap map[global.ComponentName]*Theme
type Themes []*Theme

type Theme struct {
	*Component
}

func NewThemes() Themes {
	return make(Themes, 0)
}

func NewTheme(theme Componenter) *Theme {
	return &Theme{
		Component: NewComponent(theme.GetName(), theme),
	}
}

func (me *Themes) Dedup() {
	tm := me.GetMap()
	*me = (*me)[0:0]
	for _, t := range tm {
		*me = append(*me, t)
	}
}

func (me *Themes) GetMap() ThemeMap {
	tm := make(ThemeMap, 0)
	for _, t := range *me {
		tm[t.Subdir] = t
	}
	return tm
}
