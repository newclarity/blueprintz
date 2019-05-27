package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
)

var NilMeta = (*Meta)(nil)
var _ jsonfile.Metaer = NilMeta

type Meta struct {
	CreatedBy     string
	AboutUrl      global.Url
	SchemaVer     global.Version
	StepStatusMap global.StepStatusMap
}

func NewMeta() *Meta {
	return &Meta{
		CreatedBy: global.JsonSchemaCreatedBy,
		AboutUrl:  global.AboutBlueprintzUrl,
		SchemaVer: global.JsonSchemaVersion,
	}
}
func ConvertJsonfileMeta() *Meta {
	return NewMeta()
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
