package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/olekukonko/tablewriter"
)

type FileInfo struct {
	Name    string
	Size    int64
	Mode    fs.FileMode
	ModTime time.Time
	IsDir   bool
}

func readDirectory(path string) ([]FileInfo, error) {
	var files []FileInfo

	items, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		fileInfo, err := item.Info()

		if err != nil {
			return nil, err
		}

		files = append(files, FileInfo{
			Name:    item.Name(),
			Size:    fileInfo.Size(),
			Mode:    fileInfo.Mode(),
			ModTime: fileInfo.ModTime(),
			IsDir:   item.IsDir(),
		})
	}
	return files, nil
}

func displayFiles(files []FileInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Size (Bytes)", "Modified Time", "Permissions", "Type"})

	for _, file := range files {
		table.Append([]string{
			file.Name,
			fmt.Sprintf("%d", file.Size),
			file.ModTime.Format("2006-01-02 15:04:05"),
			file.Mode.String(),
			map[bool]string{true: "Directory", false: "File"}[file.IsDir],
		})
	}
	table.Render()
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the path relative to the current working directory
	rootPath := filepath.Join(cwd, "files")

	// path := "../files/"

	files, err := readDirectory(rootPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	displayFiles(files)
}
