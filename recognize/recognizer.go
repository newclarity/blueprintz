package recognize

import (
	"blueprintz/global"
	"fmt"
)

type Map map[global.RecognizerName]Recognizer
type List []Recognizer

type Recognizer interface {
	Recognizes() global.ComponentTypes
	Match(*Args) bool
	GetSourceUrl(Componenter) global.Url
}

type Args struct {
	Website global.Url
}

func (me *Args) String() string {
	return fmt.Sprintf("Website=[%s]", me.Website)
}
