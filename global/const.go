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
	WordPressDialect    Dialect = "wordpress"
	ClassicPressDialect Dialect = "classicpress"
	PantheonWPDialect   Dialect = "pantheon-wordpress" //https://github.com/pantheon-systems/WordPress
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)

const (
	YesState   YesNo = "yes"
	NoState    YesNo = "no"
	UnsetState YesNo = "   "
)

type HelpIdType = string

const (
	CoreHelpId      HelpIdType = "core"
	LayoutHelpId    HelpIdType = "layout"
	MuPluginHelpId  HelpIdType = "mu-plugin"
	MuPluginsHelpId HelpIdType = "mu-plugins"
	PluginHelpId    HelpIdType = "plugin"
	PluginsHelpId   HelpIdType = "plugins"
	ProjectHelpId   HelpIdType = "project"
	ThemeHelpId     HelpIdType = "theme"
	ThemesHelpId    HelpIdType = "themes"
	ComponentHelpId HelpIdType = "component"
	UnknownHelpId   HelpIdType = "unknown"
)
