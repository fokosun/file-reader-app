package filehandler

import "os"

func ReadDirectory(path string) ([]FileInfo, error) {
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
