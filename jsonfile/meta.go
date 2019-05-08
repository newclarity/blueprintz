package jsonfile

import "blueprintz/global"

type Metaer interface {
	GetAboutUrl() global.Url
	GetSchemaVer() global.Version
}

type Meta struct {
	AboutUrl  global.Url     `json:"about_url"`
	SchemaVer global.Version `json:"schema_ver"`
}

func NewMeta(metaer Metaer) *Meta {
	return &Meta{
		AboutUrl:  metaer.GetAboutUrl(),
		SchemaVer: metaer.GetSchemaVer(),
	}
}
