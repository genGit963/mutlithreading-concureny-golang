package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitGroup sync.WaitGroup
	lock      sync.Mutex
)

func searchFile(root, filename string) {
	defer waitGroup.Done() // Ensure Done is called at the end of the function

	fmt.Println("Searching in:", root)

	// Get all files and handle errors
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", root, err)
		return
	}

	// Check all files
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			// Safely update matches using a mutex
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}

		// If the file is a directory, spawn a new goroutine
		if file.IsDir() {
			waitGroup.Add(1) // Increment before launching a new goroutine
			go searchFile(filepath.Join(root, file.Name()), filename)
		}
	}
}

func main() {
	root := "/Users/maheshbogati/Desktop"
	filename := "README.md"

	// Start the search
	waitGroup.Add(1)
	go searchFile(root, filename)

	// Wait for all goroutines to complete
	waitGroup.Wait()

	// Print matched files
	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
}
