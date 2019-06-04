package global

// Flags
var NoCache bool
var ProjectDir Dir

var AllBlueprintTypes = BlueprintTypes{
	AddOnBlueprint,
	LibraryBlueprint,
	ModuleBlueprint,
	PluginBlueprint,
	ThemeBlueprint,
	WebsiteBlueprint,
}

var AllDialects = Dialects{
	WordPressDialect,
	ClassicPressDialect,
	PantheonWPDialect,
}
