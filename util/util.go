package util

import (
	"blueprintz/global"
	"encoding/json"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"syscall"
)

func GetProjectDir() global.Dir {
	dir := global.ProjectDir
	if len(dir) == 0 {
		dir = GetCurrentDir()
	}
	if dir[0] == '~' {
		dir, _ = homedir.Expand(dir)
	}
	return dir
}

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func ToAbsoluteDir(path global.Path) global.Dir {
	return fmt.Sprintf("%s%c%s",
		GetProjectDir(),
		os.PathSeparator,
		path,
	)
}

func GetExecutableFilepath() global.Filepath {
	fp, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	return global.Filepath(fp)
}

var domainExtractRegexp = regexp.MustCompile("^https?://([^/]+)")
var wwwRemoveRegexp = regexp.MustCompile("^www\\.(.+)$")

func ExtractDomain(url global.Url) (d global.Domain) {
	for range only.Once {
		match := domainExtractRegexp.FindStringSubmatch(url)
		if len(match) < 2 {
			break
		}
		d = match[1]
		match = wwwRemoveRegexp.FindStringSubmatch(d)
		if len(match) < 2 {
			break
		}
		d = match[1]
	}
	return d
}

func ErrorIsFileDoesNotExist(err error) bool {
	pe, ok := err.(*os.PathError)
	return ok && pe.Op == "open" && pe.Err == syscall.ENOENT
}

func ReadBytes(filepath global.Filepath) (b []byte, sts status.Status) {
	for range only.Once {
		var err error
		b, err = ioutil.ReadFile(string(filepath))
		if err != nil && ErrorIsFileDoesNotExist(err) {
			sts = status.Success("read %d bytes from '%s'",
				len(b),
				filepath,
			)
		}
		if err != nil {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("cannot read from '%s' file", filepath),
				Help:    fmt.Sprintf("confirm file '%s' is readable", filepath),
			})
			break
		}
	}
	return b, sts
}

func UnmarshalJson(j []byte, obj interface{}) (sts status.Status) {
	for range only.Once {
		err := json.Unmarshal(j, obj)
		if err != nil {
			sts = status.Wrap(err, &status.Args{
				Message: fmt.Sprintf("failed to unmarshal JSON for '%T'", obj),
				//Help: fmt.Sprintf("ensure '%s' is in correct format per %s",
				//	obj.GetFilepath(),
				//	obj.GetHelpUrl(),
				//),
			})
			break
		}
	}
	return sts
}
