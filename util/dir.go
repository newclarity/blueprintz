package util

import (
	"blueprintz/global"
	"os"
	"path/filepath"
)

func EntryExists(file global.Entry) bool {
	_, err := os.Stat(string(file))
	return !os.IsNotExist(err)
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

func FileExists(file global.Filepath) bool {
	return EntryExists(global.Entry(file))
}
