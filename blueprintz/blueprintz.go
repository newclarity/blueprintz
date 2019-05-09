package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/only"
	"fmt"
	"github.com/Machiel/slugify"
	"github.com/gearboxworks/go-status/is"
	"regexp"
	"strings"
)

var NilBlueprintz = (*Blueprintz)(nil)
var _ jsonfile.Blueprinter = NilBlueprintz

type Blueprintz struct {
	Name    string
	Desc    string
	Type    global.BlueprintType
	Local   global.Domain
	Theme   global.ComponentName
	Layout  *Layout
	Themes  Themes
	Plugins Plugins
	Meta    *Meta
}

type Args Blueprintz

func NewBlueprintz(args ...*Args) *Blueprintz {
	var _args Args
	if len(args) == 0 {
		_args = Args{}
	} else {
		_args = *args[0]
	}
	blueprintz := &Blueprintz{}
	return blueprintz.Renew(&_args)
}

var localDomainRegex *regexp.Regexp

func init() {
	localDomainRegex = regexp.MustCompile(`.local$`)
}
func (me *Blueprintz) Renew(args ...*Args) *Blueprintz {
	*me = Blueprintz{}
	var blueprintz *Blueprintz
	if len(args) == 0 {
		blueprintz = &Blueprintz{}
	} else {
		blueprintz = (*Blueprintz)(args[0])
	}
	if blueprintz.Name == "" {
		blueprintz.Name = "Unnamed"
	}
	blueprintz.Name = strings.Title(
		localDomainRegex.ReplaceAllLiteralString(blueprintz.Name, ""),
	)
	if blueprintz.Desc == "" {
		blueprintz.Desc = fmt.Sprintf("Description about %s", blueprintz.Name)
	}
	if blueprintz.Local == "" {
		blueprintz.Local = fmt.Sprintf("%s.local",
			slugify.Slugify(blueprintz.Name),
		)
	}
	if blueprintz.Type == "" {
		blueprintz.Type = global.WebsiteBlueprint
	}
	if blueprintz.Theme == "" {
		blueprintz.Theme = "default"
	}
	if blueprintz.Meta == nil {
		blueprintz.Meta = NewMeta()
	}
	if blueprintz.Layout == nil {
		blueprintz.Layout = NewLayout()
	}
	if blueprintz.Themes == nil {
		blueprintz.Themes = make(Themes, 0)
	}
	if blueprintz.Plugins == nil {
		blueprintz.Plugins = make(Plugins, 0)
	}
	return blueprintz
}

func (me *Blueprintz) Scandir() (sts Status) {
	for range only.Once {
		sts = me.Layout.ScanDir()
		if is.Error(sts) {
			break
		}
		sts = me.Plugins.Scandir(me.Layout)
		if is.Error(sts) {
			break
		}
	}
	return sts
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
	themes := make(jsonfile.Themes, len(me.Themes))
	for i, p := range me.Themes {
		themes[i] = jsonfile.NewTheme(p)
	}
	return themes
}

func (me *Blueprintz) GetPlugins() jsonfile.Plugins {
	plugins := make(jsonfile.Plugins, len(me.Plugins))
	for i, p := range me.Plugins {
		plugins[i] = jsonfile.NewPlugin(p)
	}
	return plugins
}
