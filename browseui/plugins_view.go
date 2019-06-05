package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"sort"
)

var NilPluginsView = (*PluginsView)(nil)
var _ tui.Viewer = NilPluginsView

type PluginsView struct {
	*BaseView
	Plugins blueprintz.Plugins
}

func NewPluginsView(ui *BrowseUi) *PluginsView {
	pns := &PluginsView{
		BaseView: NewBaseView(ui),
		Plugins:  ui.Blueprintz.Plugins,
	}
	pns.Embedder = pns
	return pns
}

func (me *PluginsView) GetLabel() global.Label {
	return global.PluginsLabel
}

func (me *PluginsView) GetChildren() tui.Viewers {
	tns := make(tui.Viewers, len(me.Plugins))
	for i, tn := range me.Plugins {
		tns[i] = NewPluginView(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}
func (me *PluginsView) GetHelpId() global.Slug {
	return global.PluginsHelpId
}
