package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	codeFolder := "/Users/antoinegagnon/Code/"
	files, err := os.ReadDir(codeFolder)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 1 {
		log.Fatal("Please provide at least one argument.")
	}
	for _, f := range files {
		if strings.Contains(strings.ToLower(f.Name()), strings.ToLower(os.Args[1])) {
			fmt.Print(codeFolder + f.Name())
			return
		}
	}
}
