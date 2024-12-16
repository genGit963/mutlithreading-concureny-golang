package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	matches []string
)

func searchFile(root, filename string) {

	fmt.Println("Searching in: ", root)

	// get all files
	files, _ := os.ReadDir(root)

	// check all files
	for _, file := range files {

		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, file.Name()))
		}

		if file.IsDir() {
			searchFile(filepath.Join(root, file.Name()), filename)
		}

	}

}

func main() {
	searchFile("/Users/maheshbogati/Desktop/multi-threading-golang", "README.md")

	for _, file := range matches {
		fmt.Println("Machted: ", file)

	}

}
