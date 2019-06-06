package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/help"
	"blueprintz/tui"
	"sort"
)

var NilThemesView = (*ThemesView)(nil)
var _ tui.Viewer = NilThemesView

type ThemesView struct {
	*BaseView
	Themes blueprintz.Themes
}

func NewThemesView(ui *BrowseUi) *ThemesView {
	tns := &ThemesView{
		BaseView: NewBaseView(ui),
		Themes:   ui.Blueprintz.Themes,
	}
	tns.Embedder = tns
	return tns
}

func (me *ThemesView) GetLabel() global.Label {
	return global.ThemesLabel
}

func (me *ThemesView) GetChildren() tui.Viewers {
	tns := make(tui.Viewers, len(me.Themes))
	for i, tn := range me.Themes {
		tns[i] = NewThemeView(me.Ui, tn)
	}
	sort.Slice(tns, func(i, j int) bool {
		return tns[i].GetLabel() < tns[j].GetLabel()
	})
	return tns
}

func (me *ThemesView) GetHelpInfo() *help.Info {
	return &help.Info{
		Id:    global.ThemesHelpId,
		Label: global.ThemesLabel,
	}
}
