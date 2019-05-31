package util

import (
	"blueprintz/global"
	"fmt"
	"github.com/gearboxworks/go-status/only"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

func DirExists(dir global.Dir) bool {
	return EntryExists(global.Entry(dir))
}
func MaybeMakeDir(dir global.Dir, perms os.FileMode) (err error) {
	if !DirExists(dir) {
		err = os.MkdirAll(string(dir), perms)
	}
	return err
}
func FileDir(file global.Filepath) global.Dir {
	return global.Dir(filepath.Dir(string(file)))
}
func ParentDir(file global.Dir) global.Dir {
	return global.Dir(filepath.Dir(string(file)))
}

func EntryExists(file global.Entry) bool {
	_, err := os.Stat(string(file))
	return !os.IsNotExist(err)
}
func FileExists(file global.Filepath) bool {
	return EntryExists(global.Entry(file))
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
