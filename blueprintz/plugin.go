package blueprintz

import (
	"blueprintz/fileheaders"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"sort"
	"strings"
)

var NilPlugin = (*Plugin)(nil)
var _ jsonfile.Componenter = NilPlugin
var _ recognize.Componenter = NilPlugin

type PluginMap map[global.Slug]*Plugin
type Plugins []*Plugin

type Plugin struct {
	PluginName global.ComponentName
	PluginURI  global.Url
	*Component
}

func NewPlugin(fh *fileheaders.Plugin) *Plugin {
	return &Plugin{
		PluginName: fh.PluginName,
		PluginURI:  fh.PluginURI,
		Component: &Component{
			Version:  fh.Version,
			Subdir:   fh.GetSubdir(),
			Basefile: fh.GetBasefile(),
		},
	}
}

func (me *Plugin) Research(bpz *Blueprintz) {
	me.DownloadUrl = ""
	me.SourceType = "custom"
	for _, r := range bpz.GetRecognizerMap() {
		if !recognize.IsValidComponentType(me, r) {
			continue
		}
		if r.MatchesComponent(me) {
			me.DownloadUrl = r.GetComponentDownloadUrl(me)
			me.SourceType = global.OpenSourceCode
			continue
		}
	}
	if me.DownloadUrl == "" {
		me.matchSourceType(bpz.Legend.Sources)
	}
}

func normalizeUrl(url global.Url) global.Url {
	return strings.ReplaceAll(url, "https:", "http:")
}

func (me *Plugin) matchSourceType(sources Sources) (matched bool) {
	for range only.Once {
		ws := normalizeUrl(me.GetWebsite())
		if ws == "" {
			break
		}
		for _, s := range sources {
			if !strings.HasPrefix(ws, string(normalizeUrl(s.Website))) {
				continue
			}
			me.SourceType = s.SourceType
			matched = true
			break
		}
		if matched {
			break
		}
	}
	return matched
}

func (me *Plugin) GetType() global.ComponentType {
	return global.PluginComponent
}
func (me *Plugin) GetName() global.ComponentName {
	return me.PluginName
}

func (me *Plugin) GetWebsite() global.Url {
	return me.PluginURI
}

func (me *Plugins) Scandir(path global.Path, allowHeaderless bool) (sts Status) {
	for range only.Once {
		var cs Componenters
		// Scan dir returning only plugins not in GetFileHeadersComponenterMap()
		cs, sts = fileheaders.Scandir(&fileheaders.ScandirArgs{
			ComponenterPath: path,
			FileExtension:   ".php",
			AllowHeaderless: allowHeaderless,
			ComponenterMap:  me.GetFileHeadersComponenterMap(),
		})
		if is.Error(sts) {
			break
		}
		for _, c := range cs {
			p, ok := c.(*fileheaders.Plugin)
			if !ok {
				sts = status.OurBad("type '%T' does not implement *fileheaders.Theme", c)
			}
			*me = append(*me, NewPlugin(p))
		}
		sort.Slice(*me, func(i, j int) bool {
			return (*me)[i].GetName() < (*me)[j].GetName()
		})
	}
	return sts
}

func (me *Plugins) GetFileHeadersComponenterMap() ComponenterMap {
	cm := make(ComponenterMap, 0)
	for _, t := range *me {
		cm[t.Subdir] = fileheaders.NewPlugin(t.Subdir)
	}
	return cm
}

func ConvertJsonfilePlugins(jfps jsonfile.Plugins) (ps Plugins) {
	ps = make(Plugins, len(jfps))
	for i, p := range jfps {
		ps[i] = ConvertJsonfilePlugin(p)
	}
	return ps
}

func ConvertJsonfilePlugin(jfp *jsonfile.Plugin) *Plugin {
	return &Plugin{
		PluginName: jfp.Name,
		PluginURI:  jfp.Website,
		Component: &Component{
			Version:     jfp.Version,
			Subdir:      jfp.Subdir,
			DownloadUrl: jfp.DownloadUrl,
			Website:     jfp.Website,
		},
	}
}
