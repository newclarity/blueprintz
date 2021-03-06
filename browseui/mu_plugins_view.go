package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/help"
	"blueprintz/tui"
	"sort"
)

var NilMuPluginsView = (*ProjectTreeView)(nil)
var _ tui.Viewer = NilMuPluginsView

type MuPluginsView struct {
	*BaseView
	MuPlugins blueprintz.MuPlugins
}

func NewMuPluginsView(ui *BrowseUi) *MuPluginsView {
	mpns := &MuPluginsView{
		BaseView:  NewBaseView(ui),
		MuPlugins: ui.Blueprintz.MuPlugins,
	}
	mpns.Embedder = mpns
	return mpns
}

func (me *MuPluginsView) GetLabel() global.Label {
	return global.MuPluginsLabel
}

func (me *MuPluginsView) GetChildren() tui.Viewers {
	tns := make(tui.Viewers, len(me.MuPlugins))
	for i, tn := range me.MuPlugins {
		tns[i] = NewMuPluginView(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *MuPluginsView) GetHelpInfo() *help.Info {
	return &help.Info{
		Id:    global.MuPluginsHelpId,
		Label: global.MuPluginsLabel,
	}
}
