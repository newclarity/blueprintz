package blueprintz

import "blueprintz/jsonfile"

var NilLegend = (*Legend)(nil)
var _ jsonfile.Legender = NilLegend

type Legend struct {
	Sources Sources
}

type LegendArgs Legend

func NewLegend(args ...*LegendArgs) (legend *Legend) {
	if len(args) == 0 {
		legend = &Legend{}
	} else {
		legend = (*Legend)(args[0])
	}
	if legend.Sources == nil {
		legend.Sources = make(Sources, 0)
	}
	return legend

}

func ConvertJsonfileLegend(jfl *jsonfile.Legend) *Legend {
	return &Legend{
		Sources: ConvertJsonfileSources(jfl.Sources),
	}
}

func (me *Legend) GetSources() (jfss jsonfile.Sources) {
	jfss = make(jsonfile.Sources, len(me.Sources))
	for i, s := range me.Sources {
		jfss[i] = jsonfile.NewSourceFromSourcer(s)
	}
	return jfss
}
