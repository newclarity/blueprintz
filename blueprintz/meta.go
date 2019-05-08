package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilMeta = (*Meta)(nil)
var _ jsonfile.Metaer = NilMeta

type Meta struct {
	AboutUrl  global.Url
	SchemaVer global.Version
}

func NewMeta() *Meta {
	return &Meta{
		AboutUrl:  "https://blueprintz.dev",
		SchemaVer: "0.1.0",
	}
}

func (me *Meta) GetAboutUrl() global.Url {
	return me.AboutUrl
}

func (me *Meta) GetSchemaVer() global.Version {
	return me.SchemaVer
}
