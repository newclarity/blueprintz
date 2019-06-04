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

var NilProjectTreeView = (*ProjectTreeView)(nil)
var _ tui.Viewer = NilProjectTreeView

type ProjectTreeView struct {
	*BaseView
	Ui       *BrowseUi
	Tree     *tview.TreeView
	Help     *tview.TextView
	children tui.Viewers
}

func (me *ProjectTreeView) SetForm(*tview.Form) {
	panic("implement me")
}

func NewProjectTreeView(ui *BrowseUi) *ProjectTreeView {
	ptv := &ProjectTreeView{
		BaseView: NewBaseView(ui),
		Ui:       ui,
		Help:     tview.NewTextView(),
	}
	ptv.Embedder = ptv
	ptv.Help.SetBorder(true).SetTitle("Help")
	ptv.AddChild(NewCoreView(ui))
	ptv.AddChild(NewLayoutView(ui))
	ptv.AddChild(NewThemesView(ui))
	ptv.AddChild(NewMuPluginsView(ui))
	ptv.AddChild(NewPluginsView(ui))
	ptv.Tree = ptv.makeProjectTreeView()
	ptv.Form = makeNewForm(ptv.GetLabel())
	return ptv
}

func (me *ProjectTreeView) AddChild(tn tui.Viewer) {
	tn.SetForm(makeNewForm(tn.GetLabel()))
	me.children = append(me.children, tn)
}

func (me *ProjectTreeView) GetLabel() global.Label {
	return global.ProjectLabel
}

func (me *ProjectTreeView) GetColor() tui.Color {
	return tcell.ColorAqua
}

func (me *ProjectTreeView) GetChildren() tui.Viewers {
	return me.children
}

func (me *ProjectTreeView) GetForm() *tview.Form {
	bpt := "Blueprint type:"

	bpz := me.Ui.Blueprintz

	form := me.Form.Clear(true).
		AddInputField("Project name:", bpz.Name, 20, nil, nil).
		AddInputField("Description:", bpz.Desc, 40, nil, nil).
		AddInputField("Local domain:", bpz.Local, 30, nil, nil).
		AddDropDown(bpt, global.AllBlueprintTypes, global.AllBlueprintTypes.Index(bpz.Type), nil)

	form.GetFormItemByLabel(bpt).(*tview.DropDown).SetFieldWidth(10)

	return form

}

func (me *ProjectTreeView) GetHelp() *tview.TextView {
	return tview.NewTextView()
}

func (me *ProjectTreeView) makeProjectTreeView() (tree *tview.TreeView) {
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
				ui := me.Ui
				ref, ok := node.GetReference().(tui.Viewer)
				if !ok {
					break
				}

				ui.FormBox = ref.GetForm()
				if ui.FormBox == nil {
					break
				}
				formatBox(ui.FormBox.Box, ref.GetLabel())
				ui.FullView.RemoveItem(ui.RightHandView)
				ui.RightHandView = ui.NewRightHandView()
				ui.FullView.AddItem(ui.RightHandView, 0, GoldenWide, false)

				break
				//// Load and show files in this directory.
				//tn, ok := ref.(tui.Viewer)
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
					me.Ui.Mode = FormMode
					break
				}
				if len(node.GetChildren()) > 0 {
					node.SetExpanded(!node.IsExpanded())
					break
				}
				if me.Ui.FormBox == nil {
					break
				}
				me.Ui.Mode = FormMode
				me.Ui.App.SetFocus(me.Ui.FormBox)
			}
		})

		tree.SetBorder(true).
			SetBorderPadding(1, 1, 2, 2).
			SetTitle("Project View").
			SetTitleAlign(tview.AlignCenter)

	}
	if is.Error(sts) {
		sts.SetLogAs(status.FatalLog).Log()
	}
	return tree
}
