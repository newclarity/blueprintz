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
	CustomCode     SourceType = "custom"
	CommercialCode SourceType = "commercial"
	OpenSourceCode SourceType = "opensource"
)

const (
	WordPressCore    CoreType = "wordpress"
	ClassicPressCore CoreType = "classicpress"
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)
