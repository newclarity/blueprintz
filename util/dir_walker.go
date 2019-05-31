package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DirWalker func(fp string) error

type ss []string
type mbs map[bool]ss
type smbs []mbs

func WalkDirFilesFirst(dir string, fw, dw DirWalker) (err error) {

	fs := make(smbs, 99)
	err = filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		//
		// Later @todo - Build a tree instead of a list of nodes indexed by depth
		// That will avoid having to implement inefficient skipDirs()
		//
		depth := strings.Count(path, string(os.PathSeparator))
		if depth > 99 {
			log.Fatal("Blueprintz cannot handle paths with greater than 99 segments")
		}
		if fs[depth] == nil {
			fs[depth] = make(mbs, 0)
		}
		isdir := file.IsDir()
		if _, ok := fs[depth][isdir]; !ok {
			fs[depth][isdir] = make(ss, 0)
		}
		fs[depth][isdir] = append(fs[depth][isdir], path)
		return nil
	})

	fs = stripSmbsNils(fs)

	for _, boolmap := range fs {
		if boolmap == nil {
			continue
		}
		for _, fp := range boolmap[false] {
			err = fw(fp)
			if err != nil {
				break
			}
		}
		for _, dp := range boolmap[true] {
			err = dw(dp)
			if err == filepath.SkipDir {
				fs = skipDirs(fs, dp)
			}
		}

	}
	return err
}

func skipDirs(s smbs, dp string) smbs {
	sdp := dp + "/"
	for i, mbs := range s {
		for j, ss := range mbs {
			for k, p := range ss {
				if p == "" {
					continue
				}
				if p == dp {
					s[i][j][k] = ""
					continue
				}
				if strings.HasPrefix(p, sdp) {
					s[i][j][k] = ""
					continue
				}
			}
		}
	}
	return s
}

func stripSmbsNils(s smbs) smbs {
	var first, last int
	for i, v := range s {
		if v == nil {
			continue
		}
		first = i
		break
	}
	for i := first + 1; i < 99; i++ {
		if s[i] != nil {
			continue
		}
		last = i
		break
	}
	return s[first:last]
}
