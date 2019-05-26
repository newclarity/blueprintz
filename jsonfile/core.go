package jsonfile

import "blueprintz/global"

type Coreer interface {
	GetVersion() global.Version
	GetType() global.CoreType
	GetDownloadUrl() global.Url
}

type Core struct {
	Version     global.Version  `json:"version"`
	CoreType    global.CoreType `json:"type,omitempty"`
	DownloadUrl global.Url      `json:"url,omitempty"`
}

func NewCore() *Core {
	return &Core{
		Version: global.UnknownVersion,
	}
}

func NewCoreFromCoreer(coreer Coreer) *Core {
	return &Core{
		Version:     coreer.GetVersion(),
		CoreType:    coreer.GetType(),
		DownloadUrl: coreer.GetDownloadUrl(),
	}
}
