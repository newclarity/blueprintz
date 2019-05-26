package recognize

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-status"
)

type (
	Status = status.Status
)

type Coreer interface {
	GetVersion() global.Version
}

type Componenter interface {
	GetType() global.ComponentType
	GetSlug() global.Slug
	GetVersion() global.Version
	GetWebsite() global.Url
}

type ComponentRecognizerMap map[global.RecognizerName]ComponentRecognizer
type ComponentRecognizerList []ComponentRecognizer

type ComponentRecognizer interface {
	ValidComponentTypes() global.ComponentTypes
	MatchesComponent(Componenter) bool
	GetComponentDownloadUrl(Componenter) string
	GetCoreDownloadUrl(Coreer) string
}
