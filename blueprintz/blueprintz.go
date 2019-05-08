package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"fmt"
	"github.com/Machiel/slugify"
	"regexp"
	"strings"
)

var NilBlueprintz = (*Blueprintz)(nil)
var _ jsonfile.Blueprinter = NilBlueprintz

type Blueprintz struct {
	Name    string               `json:"name"`
	Desc    string               `json:"desc"`
	Type    global.BlueprintType `json:"type"`
	Local   global.Domain        `json:"local"`
	Theme   global.ComponentName `json:"theme"`
	Layout  *Layout              `json:"layout"`
	Themes  Themes               `json:"themes"`
	Plugins Plugins              `json:"plugins"`
	Meta    *Meta                `json:"meta"`
}

type Args Blueprintz

func NewBlueprintz(args ...*Args) *Blueprintz {
	var bpz *Blueprintz
	if len(args) == 0 {
		bpz = &Blueprintz{}
	} else {
		bpz = (*Blueprintz)(args[0])
	}
	if bpz.Name == "" {
		bpz.Name = "Unnamed"
	}
	re := regexp.MustCompile(`.local$`)
	bpz.Name = re.ReplaceAllLiteralString(bpz.Name, "")
	bpz.Name = strings.Title(bpz.Name)
	if bpz.Desc == "" {
		bpz.Desc = fmt.Sprintf("Description about %s", bpz.Name)
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
		bpz.Layout = NewLayout()
	}
	if bpz.Themes == nil {
		bpz.Themes = make(Themes, 0)
	}
	if bpz.Plugins == nil {
		bpz.Plugins = make(Plugins, 0)
	}
	return bpz
}

func (me *Blueprintz) GetName() string {
	return me.Name
}

func (me *Blueprintz) GetDesc() string {
	return me.Desc
}

func (me *Blueprintz) GetType() global.BlueprintType {
	return me.Type
}

func (me *Blueprintz) GetLocal() global.Domain {
	return me.Local
}

func (me *Blueprintz) GetTheme() global.ComponentName {
	return me.Theme
}

func (me *Blueprintz) GetLayout() *jsonfile.Layout {
	return jsonfile.NewLayout(me.Layout)
}

func (me *Blueprintz) GetMeta() *jsonfile.Meta {
	return jsonfile.NewMeta(me.Meta)
}

func (me *Blueprintz) GetThemes() jsonfile.Themes {
	return jsonfile.Themes{}
}

func (me *Blueprintz) GetPlugins() jsonfile.Plugins {
	return jsonfile.Plugins{}
}
