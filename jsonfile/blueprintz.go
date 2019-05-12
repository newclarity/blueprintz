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
	GetJsonName() string
	GetJsonDesc() string
	GetJsonType() global.BlueprintType
	GetJsonLocal() global.Domain
	GetJsonTheme() global.ComponentName
	GetJsonLayout() *Layout
	GetJsonThemes() Themes
	GetJsonPlugins() Plugins
	GetJsonMeta() *Meta
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
		Name:    blueprintz.GetJsonName(),
		Desc:    blueprintz.GetJsonDesc(),
		Type:    blueprintz.GetJsonType(),
		Local:   blueprintz.GetJsonLocal(),
		Theme:   blueprintz.GetJsonTheme(),
		Layout:  blueprintz.GetJsonLayout(),
		Themes:  blueprintz.GetJsonThemes(),
		Plugins: blueprintz.GetJsonPlugins(),
		Meta:    blueprintz.GetJsonMeta(),
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
