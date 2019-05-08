package jsonfile

import (
	"blueprintz/global"
	"blueprintz/only"
	"blueprintz/util"
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/go-status"
	"io/ioutil"
	"os"
)

type Blueprinter interface {
	GetName() string
	GetDesc() string
	GetType() global.BlueprintType
	GetLocal() global.Domain
	GetTheme() global.ComponentName
	GetLayout() *Layout
	GetThemes() Themes
	GetPlugins() Plugins
	GetMeta() *Meta
}

type Blueprintz struct {
	Name    string               `json:"name"`
	Desc    string               `json:"desc"`
	Type    global.BlueprintType `json:"type"`
	Local   global.Domain        `json:"local"`
	Theme   global.ComponentName `json:"theme"`
	Meta    *Meta                `json:"meta"`
	Layout  *Layout              `json:"layout"`
	Themes  Themes               `json:"themes"`
	Plugins Plugins              `json:"plugins"`
}

func NewBlueprintz(blueprintz Blueprinter) *Blueprintz {
	return &Blueprintz{
		Name:    blueprintz.GetName(),
		Desc:    blueprintz.GetDesc(),
		Type:    blueprintz.GetType(),
		Local:   blueprintz.GetLocal(),
		Theme:   blueprintz.GetTheme(),
		Layout:  blueprintz.GetLayout(),
		Themes:  blueprintz.GetThemes(),
		Plugins: blueprintz.GetPlugins(),
		Meta:    blueprintz.GetMeta(),
	}
}

func (me *Blueprintz) WriteFile() (sts Status) {
	for range only.Once {
		b, err := json.MarshalIndent(me, "", "\t")
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot marshal Blueprintz")
			break
		}
		fp := fmt.Sprintf("%s%c%s",
			util.GetCurrentDir(),
			os.PathSeparator,
			global.BlueprintzFile,
		)
		err = ioutil.WriteFile(fp, b, os.ModePerm)
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot write '%s'", fp)
			break
		}
	}
	return sts
}
