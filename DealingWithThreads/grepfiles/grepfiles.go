package grepfiles

import (
	"fmt"
	"os"
	"strings"
)

func Grepfiles(matchString string, filename string) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	present := strings.Contains(fileString, matchString)
	if present {
		fmt.Printf("%s contains %s\n", filename, matchString)
	}
}