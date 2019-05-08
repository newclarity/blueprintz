package blueprintz

import "blueprintz/global"

type ThemeMap map[global.ComponentName]*Theme
type Themes []*Theme

type Theme struct {
	Component
}
