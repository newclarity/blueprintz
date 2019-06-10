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

var AllDialects = DialectNames{
	WordPressDialect,
	ClassicPressDialect,
	PantheonWPDialect,
}

var YesNoOptions = YesNos{
	UnsetState,
	YesState,
	NoState,
}
