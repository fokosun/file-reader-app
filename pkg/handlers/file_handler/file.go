package filehandler

import (
	"fmt"
	"io/fs"
	"os"
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

func DisplayFiles(files []FileInfo) {
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
