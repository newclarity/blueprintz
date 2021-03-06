package fileheaders

import (
	"blueprintz/global"
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

func noop(i ...interface{}) {}

func MakeComponenter(typ string, fp global.Filepath) (c Componenter, sts Status) {
	switch typ {
	case global.PluginComponent, global.MuPluginComponent:
		c = NewPlugin(fp)
	case global.ThemeComponent:
		c = NewTheme(fp)
	default:
		sts = status.Fail().SetMessage("Invalid Componenter type '%s'", typ)
	}
	return c, sts
}

func ReadFileHeaders(component Componenter) (c Componenter, sts util.Status) {
	for range only.Once {
		sts = component.ReadHeader(component)
		if is.Error(sts) || is.Warn(sts) {
			break
		}
		c = component
	}
	return c, sts
}

func FindHeaderFile(component Componenter, dp global.Dir, ext string) (c Componenter, sts util.Status) {
	var files []os.FileInfo
	var i int
	var fn global.Basefile
	fn = fmt.Sprintf("%s%s", filepath.Base(dp), ext)
	for {
		tryfile := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, fn)
		if util.FileExists(tryfile) {
			component.SetFilepath(tryfile)
			c, sts = ReadFileHeaders(component)
			if is.Error(sts) {
				break
			}
			if c != nil {
				c = component
				sts = status.Success("header file '%s' found", c.GetFilepath())
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
			if strings.HasSuffix(strings.ToLower(fn), ext) {
				break
			}
			if i >= len(files) {
				break
			}
		}
	}
	return c, sts
}

type ScandirArgs struct {
	ComponenterPath global.Path
	FileExtension   string
	AllowHeaderless bool
	ComponenterMap  ComponenterMap
}

func Scandir(args *ScandirArgs) (cs Componenters, sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(args.ComponenterPath)
		files, err := ioutil.ReadDir(dp)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to read directory '%s'", dp)
			break
		}
		cs = make(Componenters, 0)
		for _, f := range files {
			n := f.Name()
			if n[0] == '.' {
				// Ignore "hidden" plugins and themes
				continue
			}
			if _, ok := args.ComponenterMap[n]; ok {
				// Tell them we already got a one
				continue
			}
			fp := fmt.Sprintf("%s%c%s", dp, os.PathSeparator, n)
			ctype := strings.TrimRight(filepath.Base(args.ComponenterPath), "s")
			c, sts := MakeComponenter(ctype, fp)
			if is.Error(sts) {
				break
			}
			c.SetAllowHeaderless(args.AllowHeaderless)
			if f.IsDir() {
				c, sts = FindHeaderFile(c, fp, args.FileExtension)
			} else {
				c.SetIsRootFile(true)
				c, sts = ReadFileHeaders(c)
			}
			if is.Error(sts) {
				break
			}
			if c == nil {
				// File was not a valid Componenter
				continue
			}
			cs = append(cs, c)
		}
	}
	return cs, sts
}
