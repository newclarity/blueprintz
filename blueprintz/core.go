package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/tui"
	"blueprintz/util"
	"bufio"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"os"
)

var NilCore = (*Core)(nil)
var _ jsonfile.Coreer = NilCore
var _ tui.TreeNoder = NilCore

type Core struct {
	Version     global.Version
	CoreType    global.Version
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
		CoreType:    jfc.CoreType,
		DownloadUrl: jfc.DownloadUrl,
	}
}

func (me *Core) GetLabel() global.NodeLabel {
	return global.CoreNode
}

func (me *Core) GetReference() interface{} {
	return me
}

func (me *Core) IsSelectable() bool {
	return true
}

func (me *Core) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me *Core) GetChildren() tui.TreeNoders {
	return nil
}

func (me *Core) Research() {
	//me.DownloadUrl = ""
	//me.CoreType = ""
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

func (me *Core) GetType() global.CoreType {
	return me.CoreType
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
				me.CoreType = vf.CoreType
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
