package config

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

type Config struct {
	About         string    `json:"about"`
	LearnMore     string    `json:"learn_more"`
	OsBridge      OsBridger `json:"-"`
	SchemaVersion string    `json:"schema_version"`
}

func UnmarshalConfig(b []byte) *Config {
	c := Config{}
	_ = json.Unmarshal(b, &c)
	return &c
}

func NewConfig(OsBridge OsBridger) *Config {
	c := &Config{
		About:     "This is a Blueprintz user configuration file.",
		LearnMore: "To learn about Blueprintz visit https://blueprintz.dev",
		OsBridge:  OsBridge,
	}
	return c
}

func (me *Config) Initialize() (sts Status) {
	sts = me.Load()
	if status.IsError(sts) {
		sts = me.WriteFile()
	}
	return sts
}
func (me *Config) Bytes() []byte {
	b, _ := json.Marshal(me)
	return b
}

func (me *Config) GetDir() global.Dir {
	return me.OsBridge.GetUserConfigDir()
}

func (me *Config) GetFilepath() global.Filepath {
	return fmt.Sprintf("%s%cconfig.json",
		me.OsBridge.GetUserConfigDir(),
		os.PathSeparator,
	)
}

func (me *Config) WriteFile() (sts Status) {
	for range only.Once {
		j, err := json.MarshalIndent(me, "", "    ")
		if err != nil {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("unable to marshal config"),
				Help:    util.ContactSupportHelp(),
			})
			break
		}
		sts = me.MaybeMakeDir(me.GetDir(), os.ModePerm)
		if status.IsError(sts) {
			break
		}
		err = ioutil.WriteFile(string(me.GetFilepath()), j, os.ModePerm)
		if err != nil {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("unable to write to config file '%s'", me.GetFilepath()),
				Help:    fmt.Sprintf("check '%s' for write permissions", util.FileDir(me.GetFilepath())),
			})
			break
		}
		sts = status.Success("project config file written")
	}
	return sts
}

func (me *Config) MaybeMakeDir(dir global.Dir, mode os.FileMode) (sts Status) {
	for range only.Once {
		err := util.MaybeMakeDir(dir, mode)
		if err == nil {
			sts = status.Success("directory '%s' created", dir)
			break
		}
		sts = status.Wrap(err, &status.Args{
			Message: fmt.Sprintf("failed to create directory '%s'", dir),
			Help:    fmt.Sprintf("confirm directory '%s' is readable", util.ParentDir(dir)),
		})

	}
	return sts
}

func (me *Config) ReadBytes() (b []byte, sts Status) {
	for range only.Once {
		fp := me.GetFilepath()
		b, sts = util.ReadBytes(fp)
		if status.IsError(sts) {
			break
		}
		sts = status.Success("read %d bytes from file '%s'.", len(b), fp)
	}
	return b, sts
}

func (me *Config) Unmarshal(j []byte) (sts Status) {
	for range only.Once {
		sts := util.UnmarshalJson(j, me)
		if status.IsError(sts) {
			break
		}
		sts = status.Success("bytes unmarshalled")
	}
	return sts
}

func (me *Config) Load() (sts Status) {
	for range only.Once {
		var j []byte
		j, sts = me.ReadBytes()
		if status.IsError(sts) {
			break
		}
		if len(j) > 0 {
			sts = me.Unmarshal(j)
		}
		if status.IsError(sts) {
			break
		}
	}
	return sts
}
