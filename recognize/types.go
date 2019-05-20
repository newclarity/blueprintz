package recognize

import "blueprintz/global"

type Componenter interface {
	GetSlug() global.Slug
	GetVersion() global.Version
}
