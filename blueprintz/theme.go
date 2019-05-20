package blueprintz

import (
	"blueprintz/fileheaders"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var NilTheme = (*Theme)(nil)
var _ jsonfile.Componenter = NilTheme
var _ recognize.Componenter = NilTheme

type ThemeMap map[global.ComponentName]*Theme
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
			Version: fh.Version,
			Subdir:  fh.GetSubdir(),
		},
	}
}

func ConvertJsonfileThemes(jfts jsonfile.Themes) (ts Themes) {
	ts = make(Themes, len(jfts))
	for i, t := range jfts {
		ts[i] = ConvertJsonfileTheme(t)
	}
	return ts
}

func ConvertJsonfileTheme(jft *jsonfile.Theme) (ts *Theme) {
	return &Theme{
		ThemeName: jft.Name,
		ThemeURI:  jft.Website,
		Component: &Component{
			Version:   jft.Version,
			Subdir:    jft.Subdir,
			SourceUrl: jft.SourceUrl,
			Website:   jft.Website,
		},
	}
}

func (me *Theme) GetName() global.ComponentName {
	return me.ThemeName
}

func (me *Theme) GetWebsite() global.Url {
	return me.ThemeURI
}

func (me *Themes) Scandir(path global.Path) (sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(path)
		files, err := ioutil.ReadDir(dp)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to read directory '%s'", dp)
			break
		}
		for _, f := range files {
			if f.Name()[0] == '.' {
				// Ignore "hidden" themes
				continue
			}
			fp := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, f.Name())
			ctype := strings.TrimRight(filepath.Base(path), "s")
			c := fileheaders.MakeComponenter(ctype, fp)
			if f.IsDir() {
				c, sts = fileheaders.FindHeaderFile(c, fp, ".css")
			} else {
				c, sts = fileheaders.ReadFileHeaders(c)
			}
			if is.Error(sts) {
				break
			}
			if c == nil {
				continue
			}
			t, ok := c.(*fileheaders.Theme)
			if !ok {
				sts = status.Fail().SetMessage("Type '%T' does not implement '*fileheaders.Theme'", c)
				break
			}
			if t == nil {
				continue
			}
			*me = append(*me, NewTheme(t))
		}
	}
	return sts
}
