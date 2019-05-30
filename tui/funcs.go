package tui

import "github.com/rivo/tview"

func MakeNode(tn TreeNoder) *tview.TreeNode {
	return tview.NewTreeNode(tn.GetLabel()).
		SetReference(tn.GetReference()).
		SetSelectable(tn.IsSelectable()).
		SetColor(tn.GetColor())
}
