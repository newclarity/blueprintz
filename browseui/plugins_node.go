package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/rivo/tview"
	"sort"
)

var NilPluginsNode = (*PluginsNode)(nil)
var _ tui.TreeNoder = NilPluginsNode

type PluginsNode struct {
	*BaseNode
	Plugins blueprintz.Plugins
}

func NewPluginsNode(ui *BrowseUi) *PluginsNode {
	pns := &PluginsNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz.Plugins),
		Plugins:  ui.Blueprintz.Plugins,
	}
	pns.Embedder = pns
	return pns
}

func (me *PluginsNode) GetLabel() global.Label {
	return global.PluginsNode
}

func (me *PluginsNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.Plugins))
	for i, tn := range me.Plugins {
		tns[i] = NewPluginNode(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}
func (me *PluginsNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
