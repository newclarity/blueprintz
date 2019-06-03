package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"github.com/rivo/tview"
	"sort"
)

var NilThemesNode = (*ThemesNode)(nil)
var _ tui.TreeNoder = NilThemesNode

type ThemesNode struct {
	*BaseNode
	Themes blueprintz.Themes
}

func NewThemesNode(ui *BrowseUi) *ThemesNode {
	tns := &ThemesNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz.Themes),
		Themes:   ui.Blueprintz.Themes,
	}
	tns.Embedder = tns
	return tns
}

func (me *ThemesNode) GetLabel() global.Label {
	return global.ThemesNode
}

func (me *ThemesNode) GetChildren() tui.TreeNoders {
	tns := make(tui.TreeNoders, len(me.Themes))
	for i, tn := range me.Themes {
		tns[i] = NewThemeNode(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *ThemesNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}
