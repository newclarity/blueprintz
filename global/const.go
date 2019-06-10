package global

const AppName Name = "Blueprintz Composer™ for WordPress"
const JsonSchemaCreatedBy = AppName
const AboutBlueprintzUrl Url = "https://blueprintz.dev"
const JsonSchemaVersion Version = "0.1.0"
const UserDataPath Path = ".blueprintz"

const NavMenuSpacer = "  "
const BrowseUiNavMenu = "[lightgrey]Select: [yellow]<Enter>[lightgrey]" +
	NavMenuSpacer + "[lightgrey]Move: [yellow]<Tab>[lightgrey]" +
	NavMenuSpacer + "[lightgrey]Return: [yellow]<Escape>[lightgrey]" +
	NavMenuSpacer + "[lightgrey]Exit: [yellow]<Ctrl-C>[lightgrey]"

const BlueprintzFile = "blueprintz.json"

const UnknownVersion Version = "?.?.?"

const (
	ThemeComponent    ComponentType = "theme"
	PluginComponent   ComponentType = "plugin"
	MuPluginComponent ComponentType = "mu-plugin"
	UnknownComponent  ComponentType = "unknown"
)

const (
	WebsiteBlueprint BlueprintType = "website"
	ThemeBlueprint   BlueprintType = "theme"
	PluginBlueprint  BlueprintType = "plugin"
	LibraryBlueprint BlueprintType = "library"
	ModuleBlueprint  BlueprintType = "module"
	AddOnBlueprint   BlueprintType = "add-on"
)
const (
	PendingStatus     StepStatus = "pending"
	NeedsReviewStatus StepStatus = "needsreview"
	NeedsInputStatus  StepStatus = "needsinput"
	CompleteStatus    StepStatus = "complete"
)

const (
	InitializeStep         StepType = "01. Initialize blueprintz.json"
	ScanSourceStep         StepType = "02. Scan core, plugins & themes"
	DeriveCompAttrsStep    StepType = "03. Derive component attributes"
	DesignateAuthAttrsStep StepType = "04. Designate author attributes"
)

const (
	WordPressDialect    DialectName = "wordpress"
	ClassicPressDialect DialectName = "classicpress"
	PantheonWPDialect   DialectName = "pantheon-wordpress" //https://github.com/pantheon-systems/WordPress
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)

const (
	YesState   YesNo = "yes"
	NoState    YesNo = "no"
	UnsetState YesNo = "   "
)

type HelpId = string

const (
	CoreHelpId       HelpId = "core"
	LayoutHelpId     HelpId = "layout"
	MuPluginHelpId   HelpId = "mu-plugin"
	MuPluginsHelpId  HelpId = "mu-plugins"
	PluginHelpId     HelpId = "plugin"
	PluginsHelpId    HelpId = "plugins"
	ProjectHelpId    HelpId = "project"
	ThemeHelpId      HelpId = "theme"
	ThemesHelpId     HelpId = "themes"
	ComponentHelpId  HelpId = "component"
	ComponentsHelpId HelpId = "components"
	UnknownHelpId    HelpId = "unknown"
)

const (
	HelpLabel Label = "Help"
)
