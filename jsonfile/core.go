package jsonfile

import "blueprintz/global"

type Coreer interface {
	GetVersion() global.Version
}

type Core struct {
	Version global.Version `json:"version"`
}

func NewCore() *Core {
	return &Core{
		Version: global.UnknownVersion,
	}
}
func NewCoreFromCoreer(coreer Coreer) *Core {
	return &Core{
		Version: coreer.GetVersion(),
	}
}
