package freqWords

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var wordRegex = regexp.MustCompile(`[a-zA-Z]+`)

func CountWords(url string, frequency map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("server's error " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	fmt.Println("Completed:", url)
}