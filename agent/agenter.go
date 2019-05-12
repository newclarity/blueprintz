package agent

import (
	"blueprintz/global"
	"fmt"
)

type Map map[global.AgentName]Agenter
type List []Agenter

type Agenter interface {
	Match(*Args) bool
	GetSourceUrl(Componenter) global.Url
}

type Args struct {
	Website global.Url
}

func (me *Args) String() string {
	return fmt.Sprintf("Website=[%s]", me.Website)
}
