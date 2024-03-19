package main

import (
	"github.com/devashishRaj/goConcurrency/DealingWithThreads/grepDirFiles"
	"os"
	"time"
)

/*
Adapt the grepDir program to continue searching recursively in
any subdirectories. If you give your search goroutine a file, it should search
for a string match in that file, just like in the previous exercises.Otherwise,
if you give it a directory, it should recursively spawn a new goroutine for
each file or directory found inside :
cd into grepDirFiles
go run cmd/main.go raj testdata
*/

func main() {
	if len(os.Args) < 3 {
		panic("less args than expected")
	}
	searchStr := os.Args[1]
	path := os.Args[2]
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fileinfo := range files {
		go grepDirFiles.GrepDirFile(fileinfo, path, searchStr)
	}
	time.Sleep(2 * time.Second)
}