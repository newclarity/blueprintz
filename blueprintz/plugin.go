package blueprintz

import (
	"blueprintz/fileheaders"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/only"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"io/ioutil"
	"os"
	"path/filepath"
)

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

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
			Version:   fh.Version,
			LocalSlug: fh.GetSlug(),
		},
	}
}

//
// Given an assumed known header file, read the plugin headers
// Return nil if no headers found
//
func (me Plugins) ReadHeaders(file global.Filepath) (fh *fileheaders.Plugin, sts Status) {
	for range only.Once {
		fh = fileheaders.NewPlugin(file)
		sts = fh.Read(fh)
		if is.Error(sts) {
			break
		}
	}
	if fh.Filepath == "" {
		fh = nil
	}
	return fh, sts
}

func (me Plugins) FindHeaderFile(dp global.Dir) (fp global.Filepath, fh *fileheaders.Plugin, sts Status) {
	var files []os.FileInfo
	var i int
	var fn global.Basefile
	fn = fmt.Sprintf("%s.php", filepath.Base(dp))
	for {
		tryfile := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, fn)
		fh, sts = me.ReadHeaders(tryfile)
		if is.Error(sts) {
			break
		}
		if fh != nil {
			fp = tryfile
			sts = status.Success("header file for plugin '%s' found", fh.PluginName).
				SetDetail("plugin header file is '%s'", dp).
				SetData(dp)
			break
		}
		if files == nil {
			var err error
			files, err = ioutil.ReadDir(dp)
			if err != nil {
				sts = status.Wrap(err).SetMessage("unable to read directory '%s'", dp)
				break
			}
		}
		if i >= len(files) {
			sts = status.Warn("'%s' is not a plugin directory", dp)
			break
		}
		fn = files[i].Name()
		i++
	}
	return fp, fh, sts
}

func (me *Plugins) Scandir(layouter jsonfile.Layouter) (sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(layouter.GetPluginsDir())
		files, err := ioutil.ReadDir(dp)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to read directory '%s'", dp)
			break
		}
		var fh *fileheaders.Plugin
		for _, f := range files {
			if f.Name()[0] == '.' {
				// Ignore "hidden" plugins
				continue
			}
			fp := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, f.Name())
			_, fh, sts = me.FindHeaderFile(fp)
			*me = append(*me, NewPlugin(fh))
		}
	}
	return sts
}
