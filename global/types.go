package global

type BlueprintTypes []BlueprintType

func (me BlueprintTypes) Index(t BlueprintType) int {
	return StringSliceIndex(me, t)
}

func (me BlueprintTypes) Pad(n int) BlueprintTypes {
	return StringSlicePad(me, n)
}

type BlueprintType = string
type ComponentName = string
type ComponentTypeBoolMap = map[ComponentType]bool
type ComponentTypes = []ComponentType
type ComponentType = string

type DialectNames []DialectName
type DialectName = string

func (me DialectNames) Index(d DialectName) int {
	return StringSliceIndex(me, d)
}
func (me DialectNames) Pad(n int) DialectNames {
	return StringSlicePad(me, n)
}

type UrlBoolMap = map[Url]bool
type Urls []Url
type Url = string
type Slug = string

type Versions []Version
type Version = string

func (me Versions) Index(v Version) int {
	return StringSliceIndex(me, v)
}
func (me Versions) Pad(n int) Versions {
	return StringSlicePad(me, n)
}

type Path = string
type Domain = string

type Dirs = []Dir
type Dir = string
type Filepath = string
type Basefile = string
type Entry = string

type YesNos []YesNo
type YesNo = string

func (me YesNos) Index(yn YesNo) int {
	return StringSliceIndex(me, yn)
}

type FileHeader = string

type RecognizerName = string

type Author = string

type LockerType = string

type StepType = string
type StepStatus = string
type StepStatusMap map[StepType]StepStatus

type Names = []Name
type Name = string

type Labels = []Label
type Label = string

const ProjectLabel Label = "Project"
const CoreLabel Label = "Core"
const LayoutLabel Label = "Layout"
const ThemesLabel Label = "Themes"
const PluginsLabel Label = "Plugins"
const MuPluginsLabel Label = "Mu-Plugins"
const ThemeLabel Label = "Theme"
const PluginLabel Label = "Plugin"
const MuPluginLabel Label = "Mu-Plugin"
const UnknownNode Label = "Unknown"
