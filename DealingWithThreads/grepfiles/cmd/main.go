package main

/*
Expand the program you wrote in the first exercise so that instead of printing
the contents of the text files, it searches for a string match.
The string to search for is the first argument on the command line.
When you spawn a new gorou- tine, instead of printing the file’s contents,
it should read the file and search for a match. If the goroutine finds a match,
it should output a message saying that the filename contains a match.
Call the program grepfiles.go. Here’s how you can execute this Go program
(“raj” is the search string in this example):
cd to grepfiles folder then
 go run cmd/main.go raj testdata/txtf1 testdata/txtf2 testdata/txtf3

*/
import (
	"github.com/devashishRaj/goConcurrency/DealingWithThreads/grepfiles"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		panic("not given enough requirements")
	}
	matchString := os.Args[1]
	filenames := os.Args[2:]
	for _, filename := range filenames {
		go grepfiles.Grepfiles(matchString, filename)
	}
	time.Sleep(3 * time.Second)
}