package jsonfile

var NilLegend = (*Legend)(nil)
var _ Legender = NilLegend

type Legender interface {
	GetSources() Sources
}

type Legend struct {
	Sources Sources `json:"sources"`
}

func NewLegend() *Legend {
	return &Legend{
		Sources: make(Sources, 0),
	}
}

func NewLegendFromLegender(l Legender) *Legend {
	return &Legend{
		Sources: l.GetSources(),
	}
}

func (me *Legend) GetSources() Sources {
	return me.Sources
}
