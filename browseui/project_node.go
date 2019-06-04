package browseui

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
	*BaseNode
	Ui       *BrowseUi
	Tree     *tview.TreeView
	Help     *tview.TextView
	children tui.TreeNoders
}

func (me *ProjectNode) SetForm(*tview.Form) {
	panic("implement me")
}

func NewProjectNode(ui *BrowseUi) *ProjectNode {
	pn := &ProjectNode{
		BaseNode: NewBaseNode(ui, ui.Blueprintz),
		Ui:       ui,
		Help:     tview.NewTextView(),
	}
	pn.Embedder = pn
	pn.Help.SetBorder(true).SetTitle("Help")
	pn.AddChild(NewCoreNode(ui))
	pn.AddChild(NewLayoutNode(ui))
	pn.AddChild(NewThemesNode(ui))
	pn.AddChild(NewMuPluginsNode(ui))
	pn.AddChild(NewPluginsNode(ui))
	pn.Tree = pn.makeProjectTree()
	pn.Form = makeNewForm(pn.GetLabel())
	return pn
}

func (me *ProjectNode) AddChild(tn tui.TreeNoder) {
	tn.SetForm(makeNewForm(tn.GetLabel()))
	me.children = append(me.children, tn)
}

func (me *ProjectNode) GetLabel() global.Label {
	return global.ProjectLabel
}

func (me *ProjectNode) GetColor() tui.Color {
	return tcell.ColorAqua
}

func (me *ProjectNode) GetChildren() tui.TreeNoders {
	return me.children
}

func (me *ProjectNode) GetForm() *tview.Form {
	bpt := "Blueprint type:"

	bpz := me.Ui.Blueprintz

	form := me.Form.Clear(true).
		AddInputField("Project name:", bpz.Name, 20, nil, nil).
		AddInputField("Description:", bpz.Desc, 40, nil, nil).
		AddDropDown(bpt, global.AllBlueprintTypes, global.AllBlueprintTypes.Index(bpz.Type), nil).
		AddInputField("Local domain:", bpz.Local, 30, nil, nil)

	form.GetFormItemByLabel(bpt).(*tview.DropDown).SetFieldWidth(10)

	return form

}

func (me *ProjectNode) GetHelp() *tview.TextView {
	return tview.NewTextView()
}

func (me *ProjectNode) makeProjectTree() (tree *tview.TreeView) {
	var sts Status
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
			for range only.Once {
				if node == root || node.GetChildren() == nil {
					ref, ok := node.GetReference().(tui.TreeNoder)
					if !ok {
						break
					}
					me.Ui.FormBox = ref.GetForm()
					if me.Ui.FormBox != nil {
						formatForm(ref.GetForm(), ref.GetLabel())
						//						me.Ui.FormBox.SetTitle(ref.GetLabel())
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
					me.Ui.App.SetFocus(me.Ui.FormBox)
					break
				}
				if len(node.GetChildren()) > 0 {
					node.SetExpanded(!node.IsExpanded())
					break
				}
				if me.Ui.FormBox == nil {
					break
				}
				me.Ui.App.SetFocus(me.Ui.FormBox)
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
	return tree
}
