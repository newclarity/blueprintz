package run

import (
	"blueprintz/jsonfile"
	"blueprintz/util"
	"github.com/gearboxworks/go-status"
)

func EnsureBlueprintJsonExists() (sts status.Status) {
	if !util.FileExists(jsonfile.GetFilepath()) {
		sts = status.YourBad("The file '%s' does not exist; aborting.",
			jsonfile.GetBasefile(),
		)
	}
	return sts
}
func EnsureBlueprintJsonDoesNotExist() (sts status.Status) {
	if util.FileExists(jsonfile.GetFilepath()) {
		sts = status.YourBad("The file '%s' already exists; aborting.",
			jsonfile.GetBasefile(),
		)
	}
	return sts
}
