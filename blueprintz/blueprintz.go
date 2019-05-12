package blueprintz

import (
	"blueprintz/agent"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/only"
	"fmt"
	"github.com/Machiel/slugify"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"regexp"
	"strings"
)

var NilBlueprintz = (*Blueprintz)(nil)
var _ jsonfile.Blueprinter = NilBlueprintz

type Blueprintz struct {
	Name       string
	Desc       string
	Type       global.BlueprintType
	Local      global.Domain
	Theme      global.ComponentName
	Layout     *Layout
	Themes     Themes
	Plugins    Plugins
	Meta       *Meta
	couriermap agent.Map
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
	blueprintz.couriermap = make(agent.Map, 0)

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
		sts = me.Plugins.Scandir(me.Layout.GetPluginsPath())
		if is.Error(sts) {
			break
		}
	}
	return sts
}

func (me *Blueprintz) GetJsonName() string {
	return me.Name
}

func (me *Blueprintz) GetJsonDesc() string {
	return me.Desc
}

func (me *Blueprintz) GetJsonType() global.BlueprintType {
	return me.Type
}

func (me *Blueprintz) GetJsonLocal() global.Domain {
	return me.Local
}

func (me *Blueprintz) GetJsonTheme() global.ComponentName {
	return me.Theme
}

func (me *Blueprintz) GetJsonLayout() *jsonfile.Layout {
	return jsonfile.NewLayout(me.Layout)
}

func (me *Blueprintz) GetJsonMeta() *jsonfile.Meta {
	return jsonfile.NewMeta(me.Meta)
}

func (me *Blueprintz) GetJsonThemes() jsonfile.Themes {
	themes := make(jsonfile.Themes, len(me.Themes))
	for i, p := range me.Themes {
		themes[i] = jsonfile.NewTheme(p)
	}
	return themes
}

func (me *Blueprintz) GetJsonPlugins() jsonfile.Plugins {
	plugins := make(jsonfile.Plugins, len(me.Plugins))
	for i, p := range me.Plugins {
		plugins[i] = jsonfile.NewPlugin(p)
	}
	return plugins
}

func (me *Blueprintz) FindAgent(args *agent.Args) (courier agent.Agenter, sts Status) {
	for range only.Once {
		for n, c := range me.couriermap {
			if !c.Match(args) {
				continue
			}
			sts = status.Success("found agent '%s'", n)
			courier = c
			break
		}
	}
	if courier == nil {
		sts = status.Fail().
			SetMessage("agent not found for '%s'", args.String())
	}
	return courier, sts
}

func (me *Blueprintz) GetAgent(name global.AgentName) agent.Agenter {
	c, _ := me.couriermap[name]
	return c
}

func (me *Blueprintz) RegisterAgent(name global.AgentName, c agent.Agenter) {
	me.couriermap[name] = c
}

func (me *Blueprintz) GetComponentSourceUrl(comp *Component) (url global.Url, sts Status) {
	for range only.Once {
		c, sts := me.FindAgent(&agent.Args{
			Website: comp.GetWebsite(),
		})
		if is.Error(sts) {
			break
		}
		url = c.GetSourceUrl(comp)
	}
	return url, sts
}
