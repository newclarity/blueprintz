package global

import "github.com/rivo/tview"

type BlueprintType = string
type ComponentName = string
type ComponentTypeBoolMap = map[ComponentType]bool
type ComponentTypes = []ComponentType
type ComponentType = string
type CoreType = string

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

type TreeNodeMap = map[NodeLabel]*tview.TreeNode
type NodeLabels = []NodeLabel
type NodeLabel = string

const ProjectNode NodeLabel = "Project"
const CoreNode NodeLabel = "Core"
const LayoutNode NodeLabel = "Layout"
const ThemesNode NodeLabel = "Themes"
const PluginsNode NodeLabel = "Plugins"
const MuPluginsNode NodeLabel = "MU-Plugins"
const UnknownNode NodeLabel = "Unknown"
