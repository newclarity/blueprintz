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
)

var NilCore = (*Core)(nil)
var _ jsonfile.Coreer = NilCore

type Core struct {
	Version     global.Version
	Dialect     global.DialectName
	DownloadUrl global.Url
}

func NewCore() *Core {
	return &Core{
		Version: global.UnknownVersion,
	}
}

func ConvertJsonfileCore(jfc *jsonfile.Core) *Core {
	return &Core{
		Version:     jfc.Version,
		Dialect:     jfc.Dialect,
		DownloadUrl: jfc.DownloadUrl,
	}
}

func (me *Core) Research() {
	//me.DownloadUrl = ""
	//me.DialectName = ""
	//for _, r := range bpz.GetRecognizerMap() {
	//	if !recognize.IsValidComponentType(me, r) {
	//		continue
	//	}
	//	if r.MatchesComponent(me) {
	//		me.DownloadUrl = r.GetComponentDownloadUrl(me)
	//		me.AuthorType = global.OpenAuthorCode
	//		continue
	//	}
	//}
	//if me.DownloadUrl == "" {
	//	me.matchAuthorType(bpz.Legend.Authors)
	//}
}

func (me *Core) GetVersion() global.Version {
	return me.Version
}

func (me *Core) GetDialect() global.DialectName {
	if me.Dialect == "" {
		me.Dialect = global.WordPressDialect
	}
	return me.Dialect
}

func (me *Core) GetDownloadUrl() global.Url {
	return me.DownloadUrl
}

func fileCloser(f *os.File) {
	_ = f.Close()
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
			if err := scanner.Err(); err != nil {
				sts = status.Warn("unable to read file '%s'", vfp).
					SetCause(err)
				break
			}
			for _, vf := range coreVersionFinders {
				match := vf.Regexp.FindStringSubmatch(line)
				if len(match) == 0 {
					continue
				}
				me.Dialect = vf.CoreType
				if me.Dialect == "" {
					me.Dialect = global.WordPressDialect
				}
				me.Version = match[1]
				break
			}
			if me.Version != global.UnknownVersion {
				break
			}
		}
	}
	return sts
}
