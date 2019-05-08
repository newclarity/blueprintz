package jsonfile

import "blueprintz/global"

type Meta struct {
	AboutUrl  global.Url     `json:"about_url"`
	SchemaVer global.Version `json:"schema_ver"`
}

func NewMeta() *Meta {
	return &Meta{
		AboutUrl:  "https://blueprintz.dev",
		SchemaVer: "0.1.0",
	}
}