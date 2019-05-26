package jsonfile

import (
	"blueprintz/global"
	"blueprintz/util"
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"io/ioutil"
	"os"
)

type Blueprinter interface {
	GetJsonName() string
	GetJsonDesc() string
	GetJsonType() global.BlueprintType
	GetJsonLocal() global.Domain
	GetJsonCore() *Core
	GetJsonTheme() global.ComponentName
	GetJsonLegend() *Legend
	GetJsonLayout() *Layout
	GetJsonThemes() Themes
	GetJsonPlugins() Plugins
	GetJsonMuPlugins() Plugins
	GetJsonMeta() *Meta
}

type Blueprintz struct {
	Name      string               `json:"name"`
	Desc      string               `json:"desc"`
	Type      global.BlueprintType `json:"type"`
	Local     global.Domain        `json:"local"`
	Theme     global.ComponentName `json:"theme"`
	Core      *Core                `json:"core"`
	Meta      *Meta                `json:"meta"`
	Legend    *Legend              `json:"legend"`
	Layout    *Layout              `json:"layout"`
	MuPlugins Plugins              `json:"mu-plugins"`
	Plugins   Plugins              `json:"plugins"`
	Themes    Themes               `json:"themes"`
}

func LoadJsonFile() (jbp *Blueprintz, sts Status) {
	jbp = &Blueprintz{}
	sts = jbp.LoadFile()
	return jbp, sts
}

func NewBlueprintz() *Blueprintz {
	b := Blueprintz{}
	b.Renew()
	return &b
}

func (me *Blueprintz) Renew() {
	if me.Core == nil {
		me.Core = NewCore()
	}
	if me.Meta == nil {
		me.Meta = NewMeta()
	}
	if me.Legend == nil {
		me.Legend = NewLegend()
	}
	if me.Layout == nil {
		me.Layout = NewLayout()
	}
	if me.Themes == nil {
		me.Themes = NewThemes()
	}
	if me.Plugins == nil {
		me.Plugins = NewPlugins()
	}
	if me.MuPlugins == nil {
		me.MuPlugins = NewPlugins()
	}
}

func NewBlueprintzFromBlueprintz(blueprintz Blueprinter) *Blueprintz {
	return &Blueprintz{
		Name:      blueprintz.GetJsonName(),
		Desc:      blueprintz.GetJsonDesc(),
		Type:      blueprintz.GetJsonType(),
		Local:     blueprintz.GetJsonLocal(),
		Core:      blueprintz.GetJsonCore(),
		Theme:     blueprintz.GetJsonTheme(),
		Legend:    blueprintz.GetJsonLegend(),
		Layout:    blueprintz.GetJsonLayout(),
		Themes:    blueprintz.GetJsonThemes(),
		Plugins:   blueprintz.GetJsonPlugins(),
		MuPlugins: blueprintz.GetJsonMuPlugins(),
		Meta:      blueprintz.GetJsonMeta(),
	}
}

func (me *Blueprintz) LoadFile() (sts Status) {
	for range only.Once {
		fp := GetFilepath()
		if !util.FileExists(fp) {
			sts = status.YourBad("file '%s' does not exist", fp)
			break
		}
		b, err := ioutil.ReadFile(fp)
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot read '%s'", fp)
			break
		}
		err = json.Unmarshal(b, me)
		if err != nil {
			sts = status.Wrap(err).SetMessage("unable to unmarshal '%s'", fp)
			break
		}
		me.Plugins.Dedup()
		me.Themes.Dedup()

	}
	return sts
}

func GetFilepath() global.Dir {
	return fmt.Sprintf("%s%c%s",
		util.GetCurrentDir(),
		os.PathSeparator,
		global.BlueprintzFile,
	)
}
func GetBasefile() string {
	return global.BlueprintzFile
}

func (me *Blueprintz) WriteFile() (sts Status) {
	for range only.Once {
		b, err := json.MarshalIndent(me, "", "\t")
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot marshal Blueprintz")
			break
		}
		fp := GetFilepath()
		err = ioutil.WriteFile(fp, b, os.ModePerm)
		if err != nil {
			sts = status.Wrap(err).SetMessage("cannot write '%s'", fp)
			break
		}
	}
	return sts
}
