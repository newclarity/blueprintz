package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"fmt"
	"github.com/Machiel/slugify"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"regexp"
	"strings"
)

var NilBlueprintz = (*Blueprintz)(nil)
var _ jsonfile.Blueprinter = NilBlueprintz

type Blueprintz struct {
	Name          string
	Desc          string
	Type          global.BlueprintType
	Local         global.Domain
	Theme         global.ComponentName
	Layout        *Layout
	Core          *Core
	Themes        Themes
	Plugins       Plugins
	Meta          *Meta
	recognizermap recognize.Map
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

func NewBlueprintzFromJsonfile(jfbp *jsonfile.Blueprintz) *Blueprintz {
	return NewBlueprintz(&Args{
		Name:    jfbp.Name,
		Desc:    jfbp.Desc,
		Type:    jfbp.Type,
		Local:   jfbp.Local,
		Theme:   jfbp.Theme,
		Core:    ConvertJsonfileCore(jfbp.Core),
		Layout:  ConvertJsonfileLayout(jfbp.Layout),
		Themes:  ConvertJsonfileThemes(jfbp.Themes),
		Plugins: ConvertJsonfilePlugns(jfbp.Plugins),
		Meta:    ConvertJsonfileMeta(),
	})
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
	blueprintz.recognizermap = make(recognize.Map, 0)

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
		sts = me.Core.Scandir(me.Layout.GetCorePath())
		if is.Error(sts) {
			break
		}
		sts = me.Themes.Scandir(me.Layout.GetThemesPath())
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

func (me *Blueprintz) GetJsonCore() *jsonfile.Core {
	return jsonfile.NewCoreFromCoreer(me.Core)
}

func (me *Blueprintz) GetJsonTheme() global.ComponentName {
	return me.Theme
}

func (me *Blueprintz) GetJsonLayout() *jsonfile.Layout {
	return jsonfile.NewLayoutFromLayouter(me.Layout)
}

func (me *Blueprintz) GetJsonMeta() *jsonfile.Meta {
	return jsonfile.NewMetaFromMetaer(me.Meta)
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

func (me *Blueprintz) FindRecognizer(args *recognize.Args) (recognizer recognize.Recognizer, sts Status) {
	for range only.Once {
		for n, c := range me.recognizermap {
			if !c.Match(args) {
				continue
			}
			sts = status.Success("found recognizer '%s'", n)
			recognizer = c
			break
		}
	}
	if recognizer == nil {
		sts = status.Fail().
			SetMessage("recognizer not found for '%s'", args.String())
	}
	return recognizer, sts
}

func (me *Blueprintz) GetRecognizer(name global.RecognizerName) recognize.Recognizer {
	c, _ := me.recognizermap[name]
	return c
}

func (me *Blueprintz) RegisterRecognizer(name global.RecognizerName, c recognize.Recognizer) {
	me.recognizermap[name] = c
}

func (me *Blueprintz) GetComponentSourceUrl(comp *Component) (url global.Url, sts Status) {
	for range only.Once {
		c, sts := me.FindRecognizer(&recognize.Args{
			Website: comp.GetWebsite(),
		})
		if is.Error(sts) {
			break
		}
		url = c.GetSourceUrl(comp)
	}
	return url, sts
}
