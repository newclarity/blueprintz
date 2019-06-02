package bpzui

import (
	"blueprintz/blueprintz"
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/tui"
	"github.com/gdamore/tcell"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/is"
	"github.com/gearboxworks/go-status/only"
	"github.com/rivo/tview"
	"path/filepath"
)

var NilProjectNode = (*ProjectNode)(nil)
var _ tui.TreeNoder = NilProjectNode

type ProjectNode struct {
	Ui       *BpzUi
	Tree     *tview.TreeView
	Form     *tview.Form
	Help     *tview.TextView
	children tui.TreeNoders
}

func NewProjectNode(parent *BpzUi) *ProjectNode {

	form := tview.NewForm()
	form.SetBorder(true).
		SetBorderPadding(1, 1, 3, 3).
		SetTitle(global.ProjectNode)

	pn := ProjectNode{
		Ui:   parent,
		Form: form,
		Help: tview.NewTextView(),
	}
	pn.Help.SetBorder(true).SetTitle("Help")
	pn.AddChild(NewCoreNode(parent))
	pn.AddChild(NewLayoutNode(parent))
	pn.AddChild(NewThemesNode(parent))
	pn.AddChild(NewMuPluginsNode(parent))
	pn.AddChild(NewPluginsNode(parent))
	pn.Tree, _ = pn.makeTreeView()
	return &pn
}

func (me *ProjectNode) AddChild(tn tui.TreeNoder) {
	me.children = append(me.children, tn)
}

func (me *ProjectNode) GetLabel() global.Label {
	return global.ProjectNode
}

func (me *ProjectNode) GetReference() interface{} {
	return me
}

func (me *ProjectNode) IsSelectable() bool {
	return true
}

func (me *ProjectNode) GetColor() tui.Color {
	return tcell.ColorAqua
}

func (me *ProjectNode) GetChildren() tui.TreeNoders {
	return me.children
}

func (me *ProjectNode) GetForm() *tview.Form {
	parent := me.Ui
	me.Form.Clear(true)
	me.Form.
		AddInputField("Project name:", "", 20, nil, nil).
		AddInputField("Description:", "", 40, nil, nil).
		AddDropDown("Blueprint type:", coreBlueprintTypes, 0, nil).
		AddInputField("Local domain:", "", 30, nil, nil).
		AddButton("Cancel", func() {
			parent.App.SetFocus(parent.ProjectNode)
		}).
		AddButton("Save", func() {
			parent.App.SetFocus(parent.ProjectNode)
		})
	return me.Form
}

func (me *ProjectNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}

func (me *ProjectNode) makeTreeView() (tree *tview.TreeView, sts Status) {

	for range only.Once {
		bpz := me.Ui.Blueprintz
		sts = bpz.LoadJsonfile()
		bpz, sts = blueprintz.Load()
		if is.Error(sts) {
			break
		}
		if bpz == nil {
			sts = status.Fail().SetMessage("no '%s' file found in current directory",
				filepath.Base(jsonfile.GetFilepath()),
			)
			break
		}

		root := MakeTreeNode(me)
		root.SetExpanded(true)

		tree = tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root)

		//color := tcell.ColorWhite

		tree.SetChangedFunc(func(node *tview.TreeNode) {
			//if color == tcell.ColorAqua {
			//	color = tcell.ColorLime
			//} else {
			//	color = tcell.ColorAqua
			//}
			//	me.NodeView.SetTitleColor(color)
			for range only.Once {
				if node == root || node.GetChildren() == nil {
					ref, ok := node.GetReference().(tui.TreeNoder)
					if !ok {
						break
					}
					me.Ui.NodeView = ref.GetForm()
					if me.Ui.NodeView != nil {
						me.Ui.FullView.RemoveItem(me.Ui.RightHandView)
						me.Ui.RightHandView = me.Ui.NewRightHandView()
						me.Ui.FullView.AddItem(me.Ui.RightHandView, 0, GoldenWide, false)
						//						parent.App.Draw()
					}
					break
				}
				//// Load and show files in this directory.
				//tn, ok := ref.(tui.TreeNoder)
				//if !ok {
				//	changefocus = true
				//	break
				//}
				//c := tn.GetChildren()
				//if c == nil {
				//	changefocus = true
				//	break
				//}
				//for _, cn := range c {
				//	node.AddChild(tui.MakeTreeNode(cn))
				//}
			}
		})

		// If a directory was selected, open it.
		tree.SetSelectedFunc(func(node *tview.TreeNode) {
			for range only.Once {
				if node == root {
					me.Ui.App.SetFocus(me.Ui.NodeView)
					break
				}
				if len(node.GetChildren()) > 0 {
					node.SetExpanded(!node.IsExpanded())
					break
				}
				if me.Ui.NodeView == nil {
					break
				}
				me.Ui.App.SetFocus(me.Ui.NodeView)
			}
		})

		tree.SetBorder(true).
			SetBorderPadding(1, 1, 2, 2).
			SetTitle(global.JsonSchemaCreatedBy).
			SetTitleAlign(tview.AlignCenter)

	}
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}
	return tree, sts
}
