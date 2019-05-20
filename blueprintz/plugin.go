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

var NilPlugin = (*Plugin)(nil)
var _ jsonfile.Componenter = NilPlugin
var _ recognize.Componenter = NilPlugin

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

func ConvertJsonfilePlugns(jfps jsonfile.Plugins) (ps Plugins) {
	ps = make(Plugins, len(jfps))
	for i, p := range jfps {
		ps[i] = ConvertJsonfilePlugin(p)
	}
	return ps
}

type Plugin struct {
	PluginName global.ComponentName
	PluginURI  global.Url
	*Component
}

func NewPlugin(fh *fileheaders.Plugin) *Plugin {
	return &Plugin{
		PluginName: fh.PluginName,
		PluginURI:  fh.PluginURI,
		Component: &Component{
			Version: fh.Version,
			Subdir:  fh.GetSubdir(),
		},
	}
}

func (me *Plugin) GetName() global.ComponentName {
	return me.PluginName
}

func (me *Plugin) GetWebsite() global.Url {
	return me.PluginURI
}

func (me *Plugins) Scandir(path global.Path) (sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(path)
		files, err := ioutil.ReadDir(dp)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to read directory '%s'", dp)
			break
		}
		for _, f := range files {
			if f.Name()[0] == '.' {
				// Ignore "hidden" plugins
				continue
			}
			fp := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, f.Name())
			ctype := strings.TrimRight(filepath.Base(path), "s")
			c := fileheaders.MakeComponenter(ctype, fp)
			if f.IsDir() {
				c, sts = fileheaders.FindHeaderFile(c, fp, ".php")
			} else {
				c, sts = fileheaders.ReadFileHeaders(c)
			}
			if is.Error(sts) {
				break
			}
			if c == nil {
				continue
			}
			p, ok := c.(*fileheaders.Plugin)
			if !ok {
				sts = status.Fail().SetMessage("Type '%T' does not implement '*fileheaders.Plugin'", c)
				break
			}
			if p == nil {
				continue
			}
			*me = append(*me, NewPlugin(p))
		}
	}
	return sts
}

func ConvertJsonfilePlugin(jfp *jsonfile.Plugin) (ts *Plugin) {
	return &Plugin{
		PluginName: jfp.Name,
		PluginURI:  jfp.Website,
		Component: &Component{
			Version:   jfp.Version,
			Subdir:    jfp.Subdir,
			SourceUrl: jfp.SourceUrl,
			Website:   jfp.Website,
		},
	}
}
