package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fokosun/file-reader-app/pkg/handlers"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the path relative to the current working directory
	rootPath := filepath.Join(cwd, "files")

	files, err := handlers.ReadDirectory(rootPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	fmt.Printf("%v\n", files)
	handlers.DisplayFiles(files)
}
