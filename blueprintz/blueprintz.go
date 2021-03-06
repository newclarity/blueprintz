package blueprintz

import (
	"blueprintz/config"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-osbridge"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"log"
	"regexp"
	"strings"
)

var Instance *Blueprintz
var NilBlueprintz = (*Blueprintz)(nil)
var _ jsonfile.Blueprinter = NilBlueprintz

type Blueprintz struct {
	Name          string
	Desc          string
	Type          global.BlueprintType
	SaveTo        *CodeLocker
	Theme         global.ComponentName
	Legend        *Legend
	Layout        *Layout
	Core          *Core
	Themes        Themes
	Plugins       Plugins
	MuPlugins     MuPlugins
	Meta          *Meta
	Config        *config.Config
	OsBridge      osbridge.OsBridger
	recognizermap recognize.ComponentRecognizerMap
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
	blueprintz.Renew(&_args)
	return blueprintz
}

func (me *Blueprintz) Research() {

	me.Core.Research()

	for _, p := range me.Plugins {
		p.Research(me)
	}

	for _, t := range me.Themes {
		t.Research(me)
	}

	me.CollectComponentAuthors()

}

func Load(c *config.Config) (bpz *Blueprintz, sts Status) {
	for range only.Once {
		jbp := jsonfile.NewBlueprintz()
		sts := jbp.LoadFile()
		if is.Error(sts) {
			break
		}
		bpz = NewBlueprintzFromJsonfile(c, jbp)
	}
	return bpz, sts
}

func NewBlueprintzFromJsonfile(c *config.Config, jfbp *jsonfile.Blueprintz) *Blueprintz {
	bpz := Blueprintz{}
	bpz.Config = c
	bpz.OsBridge = c.OsBridge
	bpz.RenewFromJsonfile(jfbp)
	return &bpz
}

func (me *Blueprintz) Write() (sts Status) {
	for range only.Once {
		jbp := jsonfile.NewBlueprintzFromBlueprintz(me)
		sts = jbp.WriteFile()
		if is.Error(sts) {
			break
		}
	}
	return sts
}

func (me *Blueprintz) RenewFromJsonfile(jfbp *jsonfile.Blueprintz) {
	me.Renew(&Args{
		Name:          jfbp.Name,
		Desc:          jfbp.Desc,
		Type:          jfbp.Type,
		Theme:         jfbp.Theme,
		Core:          ConvertJsonfileCore(jfbp.Core),
		Layout:        ConvertJsonfileLayout(jfbp.Layout),
		Legend:        ConvertJsonfileLegend(jfbp.Legend),
		Themes:        ConvertJsonfileThemes(jfbp.Themes),
		Plugins:       ConvertJsonfilePlugins(jfbp.Plugins),
		MuPlugins:     ConvertJsonfileMuPlugins(jfbp.MuPlugins),
		Meta:          ConvertJsonfileMeta(),
		OsBridge:      me.OsBridge,
		Config:        me.Config,
		recognizermap: me.GetRecognizerMap(),
	})
}

func (me *Blueprintz) LoadJsonfile(c *config.Config) (sts Status) {
	for range only.Once {
		var jfbp *jsonfile.Blueprintz
		jfbp, sts = jsonfile.LoadJsonFile()
		if is.Error(sts) {
			break
		}
		me.Config = c
		me.OsBridge = c.OsBridge
		me.RenewFromJsonfile(jfbp)
	}
	return sts
}

var localDomainRegex *regexp.Regexp

func init() {
	localDomainRegex = regexp.MustCompile(`.local$`)
}

func (me *Blueprintz) Renew(args ...*Args) *Blueprintz {
	var blueprintz *Blueprintz
	if len(args) == 0 {
		blueprintz = &Blueprintz{}
	} else {
		blueprintz = (*Blueprintz)(args[0])
	}

	if blueprintz.OsBridge == nil && me.OsBridge != nil {
		blueprintz.OsBridge = me.OsBridge
	}
	if blueprintz.Config == nil && me.Config != nil {
		blueprintz.Config = me.Config
	}

	if blueprintz.OsBridge == nil {
		log.Fatalf("Blueprintz much have a valid OsBridge.")
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
	if blueprintz.Type == "" {
		blueprintz.Type = global.WebsiteBlueprint
	}
	if blueprintz.Theme == "" {
		blueprintz.Theme = "default"
	}
	if blueprintz.Core == nil {
		blueprintz.Core = NewCore()
	}
	if blueprintz.Meta == nil {
		blueprintz.Meta = NewMeta()
	}
	if blueprintz.Layout == nil {
		blueprintz.Layout = NewLayout()
	}
	if blueprintz.Legend == nil {
		blueprintz.Legend = NewLegend()
	}
	if blueprintz.Themes == nil {
		blueprintz.Themes = make(Themes, 0)
	}
	if blueprintz.Plugins == nil {
		blueprintz.Plugins = make(Plugins, 0)
	}
	if blueprintz.MuPlugins == nil {
		blueprintz.MuPlugins = make(MuPlugins, 0)
	}
	if blueprintz.recognizermap == nil {
		blueprintz.recognizermap = make(recognize.ComponentRecognizerMap, 0)
	}
	if blueprintz.Config == nil {
		blueprintz.Config = config.NewConfig(blueprintz.OsBridge)
	}
	*me = *blueprintz
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
		sts = me.Plugins.Scandir(me.Layout.GetPluginsPath(), false)
		if is.Error(sts) {
			break
		}
		sts = me.MuPlugins.Scandir(me.Layout.GetMuPluginsDir(), true)
		if is.Error(sts) {
			if !strings.HasPrefix(sts.Message(), "unable to read directory") {
				break
			}
			sts = sts.SetSuccess(true)
		}
	}
	return sts
}

func (me *Blueprintz) CollectComponentAuthors() {
	ss := me.Legend.Authors
	sm := make(global.UrlBoolMap, 0)
	for _, s := range ss {
		sm[s.Website] = true
	}
	for _, t := range me.Themes {
		ss, sm = collectComponentAuthor(t.Component, ss, sm)
	}
	for _, p := range me.Plugins {
		ss, sm = collectComponentAuthor(p.Component, ss, sm)
	}
	for _, mp := range me.MuPlugins {
		ss, sm = collectComponentAuthor(mp.Component, ss, sm)
	}
	me.Legend.Authors = ss
	return
}

var ommittableDomainsRegexp = regexp.MustCompile("^(wordpress.org|github.com|bitbucket.org)$")

func collectComponentAuthor(c *Component, ss Authors, sm global.UrlBoolMap) (Authors, global.UrlBoolMap) {
	for range only.Once {
		if c.DownloadUrl != "" {
			break
		}
		ws := c.GetWebsite()
		if ws == "" {
			break
		}
		d := util.ExtractDomain(ws)
		if d == "" {
			break
		}
		if ommittableDomainsRegexp.MatchString(d) {
			break
		}
		if _, ok := sm[d]; ok {
			break
		}
		ss = append(ss, NewAuthor(d))
		sm[d] = true
	}
	return ss, sm
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

func (me *Blueprintz) GetJsonCore() *jsonfile.Core {
	return jsonfile.NewCoreFromCoreer(me.Core)
}

func (me *Blueprintz) GetJsonTheme() global.ComponentName {
	return me.Theme
}

func (me *Blueprintz) GetJsonLegend() *jsonfile.Legend {
	return jsonfile.NewLegendFromLegender(me.Legend)
}

func (me *Blueprintz) GetJsonLayout() *jsonfile.Layout {
	return jsonfile.NewLayoutFromLayouter(me.Layout)
}

func (me *Blueprintz) GetJsonMeta() *jsonfile.Meta {
	return jsonfile.NewMetaFromMetaer(me.Meta)
}

func (me *Blueprintz) GetJsonThemes() jsonfile.Themes {
	themes := make(jsonfile.Themes, len(me.Themes))
	for i, t := range me.Themes {
		themes[i] = jsonfile.NewTheme(t)
	}
	return themes
}

func (me *Blueprintz) getJsonPlugins(ps Plugins) jsonfile.Plugins {
	plugins := make(jsonfile.Plugins, len(ps))
	for i, p := range ps {
		plugins[i] = jsonfile.NewPlugin(p)
	}
	return plugins
}

func (me *Blueprintz) getJsonMuPlugins(ps MuPlugins) jsonfile.Plugins {
	plugins := make(jsonfile.Plugins, len(ps))
	for i, p := range ps {
		plugins[i] = jsonfile.NewPlugin(p)
	}
	return plugins
}

func (me *Blueprintz) GetJsonPlugins() jsonfile.Plugins {
	return me.getJsonPlugins(me.Plugins)
}

func (me *Blueprintz) GetJsonMuPlugins() jsonfile.Plugins {
	return me.getJsonMuPlugins(me.MuPlugins)
}

func (me *Blueprintz) FindRecognizer(c recognize.Componenter) (recognizer recognize.ComponentRecognizer, sts Status) {
	for n, r := range me.recognizermap {
		if !r.MatchesComponent(c) {
			continue
		}
		sts = status.Success("found recognizer '%s'", n)
		recognizer = r
		break
	}
	if recognizer == nil {
		sts = status.Fail().
			SetMessage("recognizer not found for '%s'", c.GetWebsite())
	}
	return recognizer, sts
}

func (me *Blueprintz) GetRecognizer(name global.RecognizerName) recognize.ComponentRecognizer {
	c, _ := me.recognizermap[name]
	return c
}

func (me *Blueprintz) GetRecognizerMap() recognize.ComponentRecognizerMap {
	return me.recognizermap
}

func (me *Blueprintz) RegisterRecognizer(name global.RecognizerName, c recognize.ComponentRecognizer) {
	me.recognizermap[name] = c
}
