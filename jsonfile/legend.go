package jsonfile

var NilLegend = (*Legend)(nil)
var _ Legender = NilLegend

type Legender interface {
	//	GetAuthors() Authors
}

type Legend struct {
	//	Authors Authors `json:"Authors"`
}

func NewLegend() *Legend {
	return &Legend{
		//		Authors: make(Authors, 0),
	}
}

func NewLegendFromLegender(l Legender) *Legend {
	return &Legend{
		//		Authors: l.GetAuthors(),
	}
}

//func (me *Legend) GetAuthors() Authors {
//	return me.Authors
//}
