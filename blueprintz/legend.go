package blueprintz

import "blueprintz/jsonfile"

var NilLegend = (*Legend)(nil)
var _ jsonfile.Legender = NilLegend

type Legend struct {
	Authors     Authors
	CodeLockers CodeLockers
}

type LegendArgs Legend

func NewLegend(args ...*LegendArgs) (legend *Legend) {
	if len(args) == 0 {
		legend = &Legend{}
	} else {
		legend = (*Legend)(args[0])
	}
	//if legend.Authors == nil {
	//	legend.Authors = make(Authors, 0)
	//}
	return legend

}

func ConvertJsonfileLegend(jfl *jsonfile.Legend) *Legend {
	return &Legend{
		//Authors: ConvertJsonfileAuthors(jfl.Authors),
	}
}

//func (me *Legend) GetAuthors() (jfss jsonfile.Authors) {
//	jfss = make(jsonfile.Authors, len(me.Authors))
//	for i, s := range me.Authors {
//		jfss[i] = jsonfile.NewAuthorFromAuthorer(s)
//	}
//	return jfss
//}
