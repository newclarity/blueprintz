package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilSource = (*Source)(nil)
var _ jsonfile.Sourcer = NilSource

type Sources []*Source
type Source struct {
	Website    global.Url
	SourceType global.SourceType
}

func NewSource(ws global.Url) *Source {
	return &Source{
		Website: ws,
	}
}

func (me *Source) GetWebsite() global.Url {
	return me.Website
}

func (me *Source) GetSourceType() global.SourceType {
	return me.SourceType
}

func ConvertJsonfileSources(jfss jsonfile.Sources) Sources {
	bpzss := make(Sources, len(jfss))
	for i, s := range jfss {
		bpzss[i] = ConvertJsonfileSource(s)
	}
	return bpzss
}

func ConvertJsonfileSource(jfs *jsonfile.Source) *Source {
	return &Source{
		Website:    jfs.GetWebsite(),
		SourceType: jfs.GetSourceType(),
	}
}
