package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/rivo/tview"
	"sort"
)

var NilMuPluginsNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilMuPluginsNode

type MuPluginsNode struct {
	*BaseNode
	MuPlugins blueprintz.MuPlugins
}

func NewMuPluginsNode(ui *BrowseUi) *MuPluginsNode {
	mpns := &MuPluginsNode{
		BaseNode:  NewBaseNode(ui, ui.Blueprintz.MuPlugins),
		MuPlugins: ui.Blueprintz.MuPlugins,
	}
	mpns.Embedder = mpns
	return mpns
}

func (me *MuPluginsNode) GetLabel() global.Label {
	return global.MuPluginsNode
}

func (me *MuPluginsNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.MuPlugins))
	for i, tn := range me.MuPlugins {
		tns[i] = NewMuPluginNode(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *MuPluginsNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
