package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {
	usr, _ := user.Current()
	codeFolder := usr.HomeDir+"/Code/"
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
