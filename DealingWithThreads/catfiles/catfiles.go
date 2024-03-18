package catfiles

import (
	"fmt"
	"os"
)

func Catfiles(filename string) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fmt.Println(fileString)
}