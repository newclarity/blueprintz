package blueprintz

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/only"
	"fmt"
	"github.com/gearboxworks/go-status"
	"io/ioutil"
)

type PluginMap map[global.ComponentName]*Plugin
type Plugins []*Plugin

type Plugin struct {
	Component
}

func (me Plugins) Scandir(layouter jsonfile.Layouter) (sts Status) {
	for range only.Once {
		pid := layouter.GetPluginsDir()
		files, err := ioutil.ReadDir(pid)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to read directory '%s'", pid)
		}
		for _, f := range files {
			fmt.Println(f)
		}
	}
	return sts
}
