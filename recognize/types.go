package recognize

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-status"
)

type (
	Status = status.Status
)

type Componenter interface {
	GetType() global.ComponentType
	GetSlug() global.Slug
	GetVersion() global.Version
	GetWebsite() global.Url
}

type Map map[global.RecognizerName]Recognizer
type List []Recognizer

type Recognizer interface {
	ValidTypes() global.ComponentTypes
	Matches(Componenter) bool
	GetDownloadUrl(Componenter) string
}
