package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilMeta = (*Meta)(nil)
var _ jsonfile.Metaer = NilMeta

type Meta struct {
	CreatedBy string
	AboutUrl  global.Url
	SchemaVer global.Version
}

func NewMeta() *Meta {
	return &Meta{
		CreatedBy: "Blueprintz Composer for WordPress",
		AboutUrl:  "https://blueprintz.dev",
		SchemaVer: "0.1.0",
	}
}

func (me *Meta) GetCreatedBy() global.Url {
	return me.CreatedBy
}

func (me *Meta) GetAboutUrl() global.Url {
	return me.AboutUrl
}

func (me *Meta) GetSchemaVer() global.Version {
	return me.SchemaVer
}
