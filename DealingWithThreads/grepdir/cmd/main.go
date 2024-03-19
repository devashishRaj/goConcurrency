package main

import (
	"github.com/devashishRaj/goConcurrency/DealingWithThreads/grepdir"
	"os"
	"time"
)

/*
Change the program you wrote in grepfiles so that instead of passing a
list of text filenames, you pass a directory path. The program will look inside
this directory and list the files. For each file, you can spawn a goroutine
that will search for a string match (the same as before). Call the program
grepdir.go. Here’s how you can execute this Go program:
cd to grepdir
go run cmd/main.go raj testdata
*/

func main() {
	if len(os.Args) < 3 {
		panic("less than expected arguments")
	}
	stringMatch := os.Args[1]
	filesPath := os.Args[2]
	files, err := os.ReadDir(filesPath)
	if err != nil {
		panic(err)
	}
	for _, fileinfo := range files {
		go grepdir.GrepDir(fileinfo, filesPath, stringMatch)
	}
	time.Sleep(3 * time.Second)
}