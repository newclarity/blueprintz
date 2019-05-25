package global

const JsonSchemaCreatedBy = "Blueprintz Composer for WordPress"
const AboutBlueprintzUrl = "https://blueprintz.dev"
const JsonSchemaVersion = "0.1.0"

const BlueprintzFile = "blueprintz.json"

const UnknownVersion Version = "?.?.?"

const (
	PluginComponent ComponentType = "plugin"
	ThemeComponent  ComponentType = "theme"
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
	CustomCode     Source = "custom"
	CommercialCode Source = "commercial"
	OpenSourceCode Source = "opensource"
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)
