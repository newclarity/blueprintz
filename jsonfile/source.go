package jsonfile

import "blueprintz/global"

var NilSource = (*Source)(nil)
var _ Sourcer = NilSource

type Sourcers []Sourcer
type Sourcer interface {
	GetWebsite() global.Url
	GetSourceType() global.SourceType
}

type Sources []*Source
type Source struct {
	Website    global.Url        `json:"website"`
	SourceType global.SourceType `json:"type"`
}

func NewSource(s Sourcer) *Source {
	return &Source{
		Website:    s.GetWebsite(),
		SourceType: s.GetSourceType(),
	}
}

func (me *Source) GetWebsite() global.Url {
	return me.Website
}

func (me *Source) GetSourceType() global.SourceType {
	return me.SourceType
}
