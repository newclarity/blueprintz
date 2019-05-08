package jsonfile

import (
	"blueprintz/global"
	"blueprintz/only"
	"blueprintz/util"
	"encoding/json"
	"fmt"
	"github.com/Machiel/slugify"
	"github.com/gearboxworks/go-status"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Blueprintz struct {
	Name    string               `json:"name"`
	Desc    string               `json:"desc"`
	Type    global.BlueprintType `json:"type"`
	Local   global.Domain        `json:"local"`
	Theme   global.PartName      `json:"theme"`
	Layout  *Layout              `json:"layout"`
	Themes  Themes               `json:"themes"`
	Plugins Plugins              `json:"plugins"`
	Meta    *Meta                `json:"meta"`
}

type BlueprintzArgs Blueprintz

func NewBlueprintz(args ...*BlueprintzArgs) *Blueprintz {
	var bpz *Blueprintz
	if len(args)==0 {
		bpz = &Blueprintz{}
	} else {
		bpz = (*Blueprintz)(args[0])
	}
	if bpz.Name == "" {
		bpz.Name = "Unnamed"
	}
	re := regexp.MustCompile(`.local$`)
	bpz.Name = re.ReplaceAllLiteralString(bpz.Name,"")
	bpz.Name = strings.Title(bpz.Name)
	if bpz.Desc == "" {
		bpz.Desc = fmt.Sprintf("Description about %s",bpz.Name)
	}
	if bpz.Local == "" {
		bpz.Local = fmt.Sprintf("%s.local",
			slugify.Slugify(bpz.Name),
		)
	}
	if bpz.Type == "" {
		bpz.Type = global.WebsiteBlueprint
	}
	if bpz.Theme == "" {
		bpz.Theme = "default"
	}
	if bpz.Meta == nil {
		bpz.Meta = NewMeta()
	}
	if bpz.Layout == nil {
		bpz.Layout =&Layout{}
				}
	if bpz.Themes == nil {
		bpz.Themes = make(Themes, 0)
	}
	if bpz.Plugins == nil {
		bpz.Plugins = make(Plugins, 0)
	}
	return bpz
}

func (me *Blueprintz) WriteFile() (sts Status) {
	for range only.Once {
		b, err := json.MarshalIndent(me,"","\t")
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot marshal Blueprintz")
			break
		}
		fp := fmt.Sprintf("%s%c%s",
			util.GetCurrentDir(),
			os.PathSeparator,
			global.BlueprintzFile,
		)
		err = ioutil.WriteFile(fp, b, os.ModePerm)
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot write '%s'", fp)
			break
		}
	}
	return sts
}
