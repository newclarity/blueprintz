package courier

import (
	"blueprintz/global"
	"fmt"
)

type Map map[global.CourierName]Courier
type List []Courier

type Courier interface {
	Match(*Args) bool
	GetSourceUrl(Componenter) global.Url
}

type Args struct {
	Website global.Url
}

func (me *Args) String() string {
	return fmt.Sprintf("Website=[%s]", me.Website)
}
