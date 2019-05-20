package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/util"
	"bufio"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"os"
	"regexp"
)

var NilCore = (*Core)(nil)
var _ jsonfile.Coreer = NilCore

type Core struct {
	Version global.Version
}

func NewCore() *Core {
	return &Core{
		Version: global.UnknownVersion,
	}
}

func ConvertJsonfileCore(jfc *jsonfile.Core) *Core {
	return &Core{
		Version: jfc.Version,
	}
}

func (me *Core) GetVersion() global.Version {
	return me.Version
}
func fileCloser(f *os.File) {
	_ = f.Close()
}

var versionFinder *regexp.Regexp

func init() {
	versionFinder = regexp.MustCompile(`^\$wp_version\s*=\s*['"](.+)['"];\s*$`)
}
func (me *Core) Scandir(path global.Path) (sts Status) {
	for range only.Once {
		dp := util.ToAbsoluteDir(path)
		vfp := fmt.Sprintf("%s/wp-includes/version.php", dp)
		if !util.FileExists(vfp) {
			sts = status.Warn("version file '%s' does not exist", vfp)
			break
		}
		f, err := os.Open(vfp)
		if err != nil {
			sts = status.Warn("unable to open file '%s'", vfp).
				SetCause(err)
			break
		}
		defer fileCloser(f)
		scanner := bufio.NewScanner(f) // Splits on newlines by default.
		for scanner.Scan() {
			line := scanner.Text()
			match := versionFinder.FindStringSubmatch(line)
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			if len(match) == 0 {
				continue
			}
			me.Version = match[1]
			break
		}
		if err := scanner.Err(); err != nil {
			sts = status.Warn("unable to read file '%s'", vfp).
				SetCause(err)
			break
		}
	}
	return sts
}
