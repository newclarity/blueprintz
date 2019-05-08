package jsonfile

import "blueprintz/global"

type PluginMap map[global.PartName]*Plugin
type Plugins []*Plugin

type Plugin struct {
	Name      global.PartName `json:"name"`
	Version   global.Version  `json:"version"`
	SourceUrl global.Url      `json:"source_url"`
	LocalSlug global.Slug     `json:"local_slug"`
}
