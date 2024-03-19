package grepDirFiles

import (
	"os"
	"path/filepath"
	"strings"
)

func GrepDirFile(fileinfo os.DirEntry, path, searchStr string) {
	fullpath := filepath.Join(path, fileinfo.Name())
	if fileinfo.IsDir() {
		files, err := os.ReadDir(fullpath)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			go GrepDirFile(file, fullpath, searchStr)
		}
	} else {
		content, err := os.ReadFile(fullpath)
		if err != nil {
			panic(err)
		}
		present := strings.Contains(string(content), searchStr)
		if present {
			println(fullpath, " contains the string :", searchStr)
		}
	}
}