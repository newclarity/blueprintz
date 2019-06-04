package global

const JsonSchemaCreatedBy = "Blueprintz Composer for WordPress"
const AboutBlueprintzUrl = "https://blueprintz.dev"
const JsonSchemaVersion = "0.1.0"

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
