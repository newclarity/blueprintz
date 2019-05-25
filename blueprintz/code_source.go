package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilSource = (*Source)(nil)
var _ jsonfile.Sourcer = NilSource

type Source struct {
	Custom     global.Urls
	Commercial global.Urls
	OpenSource global.Urls
}

func (me *Source) GetCustom() global.Urls {
	return me.Custom
}

func (me *Source) GetCommercial() global.Urls {
	return me.Commercial
}

func (me *Source) GetOpenSource() global.Urls {
	return me.OpenSource
}

func ConvertJsonfileSource(jfcs *jsonfile.Source) *Source {
	return &Source{
		Custom:     jfcs.Custom,
		Commercial: jfcs.Commercial,
		OpenSource: jfcs.OpenSource,
	}
}
