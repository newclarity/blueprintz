package blueprintz

import (
	"blueprintz/fileheaders"
	"github.com/gearboxworks/go-status"
)

type (
	Status = status.Status
)

type (
	Componenters   = fileheaders.Componenters
	ComponenterMap = fileheaders.ComponenterMap
)

type (
	HeaderFieldsKeyMap = map[HeaderField]bool
	HeaderFields       = []HeaderField
	HeaderField        = string
)
