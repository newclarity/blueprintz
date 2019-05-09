package util

import (
	"blueprintz/global"
	"blueprintz/only"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type DirWalker func(fp, bf string, f os.FileInfo, level int) error

func WalkDirFilesFirst(dir string, fw, dw DirWalker, level ...int) (err error) {
	var _lvl int
	if len(level) == 0 {
		_lvl = 0
	} else {
		_lvl = level[0]
	}
	for range only.Once {
		var files []os.FileInfo
		files, err = ioutil.ReadDir(dir)
		if err != nil {
			break
		}
		dirs := make([]os.FileInfo, 0, len(files))
		for _, f := range files {
			if f.IsDir() {
				dirs = append(dirs, f)
				continue
			}
			fp := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, f.Name())
			bf := filepath.Base(fp)
			err = fw(fp, bf, f, _lvl)
			if err != nil {
				break
			}
		}
		for _, d := range dirs {
			dp := fmt.Sprintf("%s%c%s", dir, os.PathSeparator, d.Name())
			bf := filepath.Base(dp)
			err = dw(dp, bf, d, _lvl)
			if err == filepath.SkipDir {
				continue
			}
			err = WalkDirFilesFirst(dp, fw, dw, _lvl+1)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
func ToAbsoluteDir(reldir global.Path) global.Dir {
	return fmt.Sprintf("%s%c%s",
		GetCurrentDir(),
		os.PathSeparator,
		reldir,
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
