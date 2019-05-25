package recognize

import (
	"blueprintz/global"
)

var NilR = (*R)(nil)
var _ Recognizer = NilR

type R struct{}

func (me *R) ValidTypes() (cts global.ComponentTypes) {
	return global.ComponentTypes{}
}

func (me *R) Matches(c Componenter) (match bool) {
	return false
}

func (me *R) GetDownloadUrl(c Componenter) (url global.Url) {
	return ""
}
