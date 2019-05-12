package blueprintz

import (
	"blueprintz/courier"
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
	"strings"
)

var NilPlugin = (*Plugin)(nil)
var _ jsonfile.Componenter = NilPlugin
var _ courier.Componenter = NilPlugin

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

type Plugin struct {
	PluginName global.ComponentName
	PluginURI  global.Url
	*Component
}

func (me *Plugin) GetName() global.ComponentName {
	return me.PluginName
}

func (me *Plugin) GetWebsite() global.Url {
	return me.PluginURI
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

//
// Given an assumed known header file, read the plugin headers
// Return nil if no headers found
//
func (me Plugins) ReadHeaders(file global.Filepath) (fh *fileheaders.Plugin, sts Status) {
	for range only.Once {
		fh = fileheaders.NewPlugin(file)
		sts = fh.Read(fh)
		if is.Error(sts) || is.Warn(sts) {
			break
		}
	}
	if is.Error(sts) || is.Warn(sts) {
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
		if util.FileExists(tryfile) {
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
		for {
			fn = files[i].Name()
			i++
			if strings.HasSuffix(strings.ToLower(fn), ".php") {
				break
			}
			if i >= len(files) {
				break
			}
		}
	}
	return fp, fh, sts
}

func (me *Plugins) Scandir(path global.Path) (sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(path)
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
			if f.IsDir() {
				_, fh, sts = me.FindHeaderFile(fp)
			} else {
				fh, sts = me.ReadHeaders(fp)
			}
			if is.Error(sts) {
				break
			}
			if fh == nil {
				continue
			}
			*me = append(*me, NewPlugin(fh))
		}
	}
	return sts
}
