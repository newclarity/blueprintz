package jsonfile

import "blueprintz/global"

var NilAuthor = (*Author)(nil)
var _ Authorer = NilAuthor

type Authorers []Authorer
type Authorer interface {
	GetWebsite() global.Url
}

type Authors []*Author
type Author struct {
	Website global.Url `json:"website"`
}

func NewAuthorFromAuthorer(s Authorer) *Author {
	return &Author{
		Website: s.GetWebsite(),
	}
}

func (me *Author) GetWebsite() global.Url {
	return me.Website
}
