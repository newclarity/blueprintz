package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilAuthor = (*Author)(nil)
var _ jsonfile.Authorer = NilAuthor

type Authors []*Author
type Author struct {
	Website global.Url
}

func NewAuthor(ws global.Url) *Author {
	return &Author{
		Website: ws,
	}
}

func (me *Author) GetWebsite() global.Url {
	return me.Website
}

func ConvertJsonfileAuthors(jfss jsonfile.Authors) Authors {
	bpzss := make(Authors, len(jfss))
	for i, s := range jfss {
		bpzss[i] = ConvertJsonfileAuthor(s)
	}
	return bpzss
}

func ConvertJsonfileAuthor(jfs *jsonfile.Author) *Author {
	return &Author{
		Website: jfs.GetWebsite(),
	}
}
