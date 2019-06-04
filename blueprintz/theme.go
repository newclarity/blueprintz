package blueprintz

import (
	"blueprintz/fileheaders"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"sort"
)

var NilTheme = (*Theme)(nil)
var _ jsonfile.Componenter = NilTheme
var _ recognize.Componenter = NilTheme

type ThemeMap map[global.Slug]*Theme
type Themes []*Theme

type Theme struct {
	ThemeName global.ComponentName
	ThemeURI  global.Url
	*Component
}

func NewTheme(fh *fileheaders.Theme) *Theme {
	return &Theme{
		ThemeName: fh.ThemeName,
		ThemeURI:  fh.ThemeURI,
		Component: &Component{
			Version:  fh.Version,
			Subdir:   fh.GetSubdir(),
			Basefile: fh.GetBasefile(),
		},
	}
}

func (me *Theme) Research(bpz *Blueprintz) {
	for _, r := range bpz.GetRecognizerMap() {
		noop(r)
	}
}

func (me *Theme) GetType() global.ComponentType {
	return global.ThemeComponent
}

func (me *Theme) GetName() global.ComponentName {
	return me.ThemeName
}

func (me *Theme) GetWebsite() global.Url {
	return me.ThemeURI
}

func (me Themes) FindTheme(tn global.ComponentName) *Theme {
	var theme *Theme
	for _, t := range me {
		if t.ThemeName != tn {
			continue
		}
		theme = t
	}
	return theme
}

func (me *Themes) Scandir(path global.Path) (sts Status) {
	for range only.Once {
		var cs Componenters
		// Scan dir returning only themes not in GetFileHeadersComponenterMap()
		cs, sts = fileheaders.Scandir(&fileheaders.ScandirArgs{
			ComponenterPath: path,
			FileExtension:   ".css",
			AllowHeaderless: false,
			ComponenterMap:  me.GetFileHeadersComponenterMap(),
		})
		for _, c := range cs {
			t, ok := c.(*fileheaders.Theme)
			if !ok {
				sts = status.OurBad("type '%T' does not implement *fileheaders.Theme", c)
			}
			*me = append(*me, NewTheme(t))
		}
		sort.Slice(*me, func(i, j int) bool {
			return (*me)[i].GetName() < (*me)[j].GetName()
		})
	}
	return sts
}

func (me *Themes) GetFileHeadersComponenterMap() ComponenterMap {
	cm := make(ComponenterMap, 0)
	for _, t := range *me {
		cm[t.Subdir] = fileheaders.NewTheme(t.Subdir)
	}
	return cm
}
func ConvertJsonfileThemes(jfts jsonfile.Themes) (ts Themes) {
	ts = make(Themes, len(jfts))
	for i, t := range jfts {
		ts[i] = ConvertJsonfileTheme(t)
	}
	return ts
}

func ConvertJsonfileTheme(jft *jsonfile.Theme) (ts *Theme) {
	var ex global.YesNo
	if jft.External == "" && jft.DownloadUrl != "" {
		ex = global.YesState
	}
	return &Theme{
		ThemeName: jft.Name,
		ThemeURI:  jft.Website,
		Component: &Component{
			Version:     jft.Version,
			Subdir:      jft.Subdir,
			Basefile:    jft.Basefile,
			DownloadUrl: jft.DownloadUrl,
			Website:     jft.Website,
			External:    ex,
		},
	}
}
