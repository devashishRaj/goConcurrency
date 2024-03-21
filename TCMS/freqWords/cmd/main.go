package main

import (
	"fmt"
	freqWords "github.com/devashishRaj/goConcurrency/TCMS/freqwords"
	"time"
)

/*
Write a program to produce a list of word frequencies .
You can use  URLs form the RFC web pages (https://rfc-editor.org/rfc/... txt pages)
Once it’s finished, the program should output a list of words with the frequency
with which each word appears in the web page. Here’s some sample output:
cd to freqWorda

$ go run cmd/main.go or go run -race cmd/main.go
the -> 5
a -> 8
car -> 1
program -> 3

*/

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i < 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		// try using goroutines to detech race condition
		// go freqWords.CountWords(url, frequency)
		freqWords.CountWords(url, frequency)
	}
	time.Sleep(10 * time.Second)
	for k, v := range frequency {
		fmt.Println(k, " -> ", v)
	}
}