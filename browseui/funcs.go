package browseui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/tui"
	"fmt"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
)

func MakeTreeNode(tn tui.TreeNoder) (tvn *tview.TreeNode) {
	for range only.Once {
		tvn = tview.NewTreeNode(tn.GetLabel()).
			SetReference(tn.GetReference()).
			SetSelectable(tn.IsSelectable()).
			SetColor(tn.GetColor())
		c := tn.GetChildren()
		if c == nil {
			break
		}
		for _, cn := range c {
			tvn.AddChild(MakeTreeNode(cn))
		}
		tvn.SetExpanded(false)
	}
	return tvn
}

func GetComponentLabel(c *blueprintz.Component) global.Label {
	var label global.Label
	for range only.Once {
		if c.Subdir != "" {
			label = c.Subdir
			break
		}
		if c.Basefile != "" {
			label = c.Basefile
			break
		}
	}
	return AddComponentVersion(c, label)
}

func AddComponentVersion(c *blueprintz.Component, label global.Label) global.Label {
	for range only.Once {
		if c.Version == "" {
			break
		}
		label = fmt.Sprintf("%s — %s", label, c.Version)
	}
	return label
}

func formatForm(form *tview.Form, label global.Label) *tview.Form {
	form.SetBorder(true).
		SetTitle(label).
		SetBorderPadding(1, 1, 3, 3)
	return form
}

func makeNewForm(label global.Label) *tview.Form {
	form := tview.NewForm()
	return formatForm(form, label)
}

func noop(i ...interface{}) interface{} { return nil }
