package grepdir

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GrepDir(fileinfo os.DirEntry, path, stringMatch string) {
	fullpath := filepath.Join(path, fileinfo.Name())
	if !fileinfo.IsDir() {
		content, err := os.ReadFile(fullpath)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(content), stringMatch) {
			fmt.Println(fullpath, " contains string ", stringMatch)
		}
	}
}