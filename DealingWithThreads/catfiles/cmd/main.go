package main

import (
	"fmt"
	"github.com/devashishRaj/goConcurrency/DealingWithThreads/catfiles"
	"os"
	"time"
)

/*
Write a program that accepts a list of text file-names as arguments.
For each filename, the program should spawn a new goroutine that will output
the contents of that file to the console. You can use the time.Sleep() function
to wait for the child goroutines to complete (until you know how to do this better).
Call the program catfiles.go. Here’s how you can execute this Go program:

cd in catfiles folder then
go run cmd/main.go testdata/txtfile1 testdata/txtfile2 testdata/txtfile3
*/

func main() {
	// Check if a file name is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide the file name as a command-line argument.")
		return
	}
	// first os.arg[0] is name of file executed
	filenames := os.Args[1:]
	for _, filename := range filenames {
		go catfiles.Catfiles(filename)
	}
	// using time.sleep for now , later there be mechanisms introduced to provide
	// synchronization
	time.Sleep(5 * time.Second)
}