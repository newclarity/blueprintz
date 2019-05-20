package global

const BlueprintzFile = "blueprintz.json"

const (
	WebsiteBlueprint BlueprintType = "website"
	ThemeBlueprint   BlueprintType = "theme"
	PluginBlueprint  BlueprintType = "plugin"
	LibraryBlueprint BlueprintType = "library"
	ModuleBlueprint  BlueprintType = "module"
	AddOnBlueprint   BlueprintType = "add-on"
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)

const UnknownVersion Version = "?.?.?"

const JsonSchemaCreatedBy = "Blueprintz Composer for WordPress"
const AboutBlueprintzUrl = "https://blueprintz.dev"
const JsonSchemaVersion = "0.1.0"
