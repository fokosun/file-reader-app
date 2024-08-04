package main

import (
	"log"
	"os"
	"path/filepath"

	filehandler "github.com/fokosun/file-reader-app/pkg/handlers/file_handler"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the path relative to the current working directory
	rootPath := filepath.Join(cwd, "files")

	files, err := filehandler.ReadDirectory(rootPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	filehandler.DisplayFiles(files)
}
