package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var NilLayout = (*Layout)(nil)
var _ jsonfile.Layouter = NilLayout

type Layout struct {
	workingDir  global.Dir
	CorePath    global.Path
	ProjectPath global.Path
	WebrootPath global.Path
	ContentPath global.Path
	VendorPath  global.Path
}

func NewLayout() *Layout {
	return &Layout{}
}
func ConvertJsonfileLayout(jflo *jsonfile.Layout) *Layout {
	return &Layout{
		CorePath:    jflo.CorePath,
		ProjectPath: jflo.ProjectPath,
		WebrootPath: jflo.WebrootPath,
		ContentPath: jflo.ContentPath,
		VendorPath:  jflo.VendorPath,
	}
}
func (me *Layout) GetCorePath() global.Path {
	return me.CorePath
}
func (me *Layout) GetProjectPath() global.Path {
	return me.ProjectPath
}
func (me *Layout) GetWebrootPath() global.Path {
	return me.WebrootPath
}
func (me *Layout) GetContentPath() global.Path {
	return me.ContentPath
}
func (me *Layout) GetVendorPath() global.Path {
	return me.VendorPath
}

func (me *Layout) GetPluginsPath() global.Path {
	return fmt.Sprintf("%s%cplugins",
		me.GetContentPath(),
		os.PathSeparator,
	)
}
func (me *Layout) GetMuPluginsDir() global.Path {
	return fmt.Sprintf("%s%cmu-plugins",
		me.GetContentPath(),
		os.PathSeparator,
	)
}
func (me *Layout) GetThemesPath() global.Path {
	return fmt.Sprintf("%s%cthemes",
		me.GetContentPath(),
		os.PathSeparator,
	)
}

func (me *Layout) String() string {
	var vp string
	if me.VendorPath != "" {
		vp = fmt.Sprintf("VendorPath: %s\n", me.VendorPath)
	}
	return fmt.Sprintf("%s%s%s%s%s",
		fmt.Sprintf("ProjectPath %s\n", me.ProjectPath),
		fmt.Sprintf("CorePath:   %s\n", me.CorePath),
		fmt.Sprintf("WebrootPath %s\n", me.WebrootPath),
		fmt.Sprintf("ContentPath %s\n", me.ContentPath),
		vp,
	)
}

func (me *Layout) IsComplete() bool {
	return me.CorePath != "" &&
		me.ProjectPath != "" &&
		me.WebrootPath != "" &&
		me.ContentPath != ""

}

var indexRegex *regexp.Regexp
var skipDirs *regexp.Regexp

func init() {
	indexRegex, _ = regexp.Compile(`require\s*\(?(.+)(/wp-blog-header.php)['"]`)
	skipDirs, _ = regexp.Compile(`^(node_modules|.vscode|.vagrant|.idea|.git|.svn|wp-admin|wp-includes)$`)
}

//require( dirname( __FILE__ ) . '/wp-blog-header.php' );
func (me *Layout) isWebRoot(fp string) (is bool) {
	for range only.Once {
		is = false
		b, err := ioutil.ReadFile(fp)
		if err != nil {
			panic(err)
		}
		match := indexRegex.FindStringSubmatch(string(b))
		if match == nil {
			break
		}
		is = match[2] == "/wp-blog-header.php"
	}
	return is

}

func (me *Layout) ScanDir() (sts Status) {
	var err error
	wd := util.GetCurrentDir()
	for range only.Once {
		me.workingDir = wd
		me.ProjectPath = me.getRelativeDir(wd)
		content := 0
		err = util.WalkDirFilesFirst(wd,
			func(fp, bf string, f os.FileInfo, lvl int) (result error) {
				for range only.Once {
					if bf[0] == '.' {
						break
					}
					if !strings.HasSuffix(bf, ".php") {
						break
					}
					if me.WebrootPath == "" && bf == "index.php" && me.isWebRoot(fp) {
						me.WebrootPath = me.getRelativeDir(fp)
					}
					if me.CorePath == "" && bf == "wp-load.php" {
						me.CorePath = me.getRelativeDir(fp)
					}
					if me.IsComplete() {
						result = io.EOF
						break
					}
				}
				return result
			},
			func(fp, bf string, f os.FileInfo, lvl int) (result error) {
				for range only.Once {
					if skipDirs.MatchString(bf) {
						result = filepath.SkipDir
						break
					}
					if bf[0] == '.' {
						break
					}
					if me.ContentPath == "" {
						switch bf {
						case "mu-plugins":
							content++
							result = filepath.SkipDir
						case "plugins":
							content++
							result = filepath.SkipDir
						case "themes":
							content++
							result = filepath.SkipDir
						}
						if content == 2 {
							me.ContentPath = me.getRelativeDir(fp)
						}
					}
					if me.IsComplete() {
						result = io.EOF
						break
					}
				}
				return result
			},
		)
		if err == io.EOF {
			break
		}
	}
	for range only.Once {
		if err == nil {
			break
		}
		if err == io.EOF {
			break
		}
		if err.Error() == "skip this directory" {
			break
		}
		sts = status.Wrap(err).SetMessage("unable to scan dir '%s'", wd)
	}
	return sts
}

func (me *Layout) getRelativeDir(fp string) (dir string) {
	if fp == me.workingDir {
		dir = "./"
	} else {
		dir = "." + filepath.Dir(string([]byte(fp)[len(me.workingDir):]))
	}
	return dir
}
