package blueprintz

import (
	"blueprintz/fileheaders"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/recognize"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"sort"
	"strings"
)

var NilPlugin = (*Plugin)(nil)
var _ jsonfile.Componenter = NilPlugin
var _ recognize.Componenter = NilPlugin
var _ tui.TreeNoder = NilPlugin

var NilPlugins = (*Plugins)(nil)
var _ tui.TreeNoder = NilPlugins

type PluginMap map[global.Slug]*Plugin
type Plugins []*Plugin

func (me Plugins) GetLabel() global.NodeLabel {
	return global.PluginsNode
}

func (me Plugins) GetReference() interface{} {
	return me
}

func (me Plugins) IsSelectable() bool {
	return true
}

func (me Plugins) GetColor() tui.Color {
	return tcell.ColorLime
}

func (me Plugins) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me))
	for i, tn := range me {
		tns[i] = tn
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

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

func (me *Plugin) GetLabel() global.NodeLabel {
	var label global.NodeLabel
	for range only.Once {
		if me.PluginName != "" {
			label = me.AddVersion(me.PluginName)
			break
		}
		label = me.Component.GetLabel()
	}
	return label
}

func (me *Plugin) GetReference() interface{} {
	return me
}

func (me *Plugin) Research(bpz *Blueprintz) {
	me.DownloadUrl = ""
	for _, r := range bpz.GetRecognizerMap() {
		if !recognize.IsValidComponentType(me, r) {
			continue
		}
		if r.MatchesComponent(me) {
			me.DownloadUrl = r.GetComponentDownloadUrl(me)
			me.External = global.YesState
			continue
		}
	}
	if me.DownloadUrl == "" {
		//me.matchAuthorType(bpz.Legend.Authors)
	}
}

func normalizeUrl(url global.Url) global.Url {
	return strings.ReplaceAll(url, "https:", "http:")
}

//func (me *Plugin) matchAuthorType(Authors Authors) (matched bool) {
//	for range only.Once {
//		ws := normalizeUrl(me.GetWebsite())
//		if ws == "" {
//			break
//		}
//		for _, s := range Authors {
//			if !strings.HasPrefix(ws, string(normalizeUrl(s.Website))) {
//				continue
//			}
//			me.AuthorType = s.AuthorType
//			matched = true
//			break
//		}
//		if matched {
//			break
//		}
//	}
//	return matched
//}

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

func ConvertJsonfileMuPlugins(jfps jsonfile.Plugins) (mups MuPlugins) {
	mups = make(MuPlugins, len(jfps))
	for i, mup := range jfps {
		mups[i] = ConvertJsonfilePlugin(mup)
	}
	return mups
}

func ConvertJsonfilePlugin(jfp *jsonfile.Plugin) *Plugin {
	var ex global.YesNo
	if jfp.External == "" && jfp.DownloadUrl != "" {
		ex = global.YesState
	}
	return &Plugin{
		PluginName: jfp.Name,
		PluginURI:  jfp.Website,
		Component: &Component{
			Version:     jfp.Version,
			Subdir:      jfp.Subdir,
			Basefile:    jfp.Basefile,
			DownloadUrl: jfp.DownloadUrl,
			Website:     jfp.Website,
			External:    ex,
		},
	}
}
