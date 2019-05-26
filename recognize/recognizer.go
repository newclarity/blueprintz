package recognize

import (
	"blueprintz/global"
)

var NilR = (*R)(nil)
var _ ComponentRecognizer = NilR

type R struct{}

func (me *R) GetCoreDownloadUrl(Coreer) string {
	return ""
}

func (me *R) ValidComponentTypes() (cts global.ComponentTypes) {
	return global.ComponentTypes{}
}

func (me *R) MatchesComponent(c Componenter) (match bool) {
	return false
}

func (me *R) GetComponentDownloadUrl(c Componenter) (url global.Url) {
	return ""
}
