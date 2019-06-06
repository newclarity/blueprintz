package jsonfile

import "blueprintz/global"

type Coreer interface {
	GetVersion() global.Version
	GetDialect() global.DialectName
	GetDownloadUrl() global.Url
}

type Core struct {
	Version     global.Version     `json:"version"`
	Dialect     global.DialectName `json:"dialect,omitempty"`
	DownloadUrl global.Url         `json:"url,omitempty"`
}

func NewCore() *Core {
	return &Core{
		Version: global.UnknownVersion,
	}
}

func NewCoreFromCoreer(coreer Coreer) *Core {
	return &Core{
		Version:     coreer.GetVersion(),
		Dialect:     coreer.GetDialect(),
		DownloadUrl: coreer.GetDownloadUrl(),
	}
}
