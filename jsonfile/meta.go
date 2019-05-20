package jsonfile

import "blueprintz/global"

type Metaer interface {
	GetAboutUrl() global.Url
	GetSchemaVer() global.Version
	GetCreatedBy() string
}

type Meta struct {
	CreatedBy string         `json:"created_by"`
	AboutUrl  global.Url     `json:"about_url"`
	SchemaVer global.Version `json:"schema_ver"`
}

func NewMeta() *Meta {
	return &Meta{
		CreatedBy: global.JsonSchemaCreatedBy,
		AboutUrl:  global.AboutBlueprintzUrl,
		SchemaVer: global.JsonSchemaVersion,
	}
}

func NewMetaFromMetaer(metaer Metaer) *Meta {
	return &Meta{
		CreatedBy: metaer.GetCreatedBy(),
		AboutUrl:  metaer.GetAboutUrl(),
		SchemaVer: metaer.GetSchemaVer(),
	}
}
