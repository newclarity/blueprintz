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
	SiteBuilderAuthor AuthorType = "sitebuilder"
	OpenSourceAuthor  AuthorType = "opensource"
	CommercialAuthor  AuthorType = "commercial"
	ContractorAuthor  AuthorType = "contractor"
	OtherAuthorType   AuthorType = "other"
)

const (
	ProjectOwner      Maintainer = "owner"
	OwnerEmployees    Maintainer = "employees"
	SoftwareVendor    Maintainer = "vendor"
	OpenSourceProject Maintainer = "opensource"
	OtherMaintainer   Maintainer = "other"
)

const (
	WordPressOrgRecognizer RecognizerName = "wordpress.org"
)
