package recognize

import "blueprintz/global"

type Componenter interface {
	GetType() global.ComponentType
	GetSlug() global.Slug
	GetVersion() global.Version
}
