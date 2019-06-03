package global

import "strings"

type BlueprintTypes []BlueprintType

func (me BlueprintTypes) Pad(n int) BlueprintTypes {
	for i, t := range me {
		me[i] = strings.Repeat(" ", n) + t + strings.Repeat(" ", n)
	}
	return me
}

type BlueprintType = string
type ComponentName = string
type ComponentTypeBoolMap = map[ComponentType]bool
type ComponentTypes = []ComponentType
type ComponentType = string
type Dialect = string

type UrlBoolMap = map[Url]bool
type Urls []Url
type Url = string
type Slug = string
type Version = string
type Path = string
type Domain = string

type Dirs = []Dir
type Dir = string
type Filepath = string
type Basefile = string
type Entry = string

type YesNo = string

type FileHeader = string

type RecognizerName = string

type Author = string

type LockerType = string

type StepType = string
type StepStatus = string
type StepStatusMap map[StepType]StepStatus

type Labels = []Label
type Label = string

const ProjectNode Label = "Project"
const CoreNode Label = "Core"
const LayoutNode Label = "Layout"
const ThemesNode Label = "Themes"
const PluginsNode Label = "Plugins"
const MuPluginsNode Label = "MU-Plugins"
const UnknownNode Label = "Unknown"
