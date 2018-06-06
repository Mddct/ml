package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	paths := make([]string, 0, 100)
	filepath.Walk("dataProcess", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	for i, path := range paths {
		if filepath.Base(path)[0] != '.' {
			newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf("split%02d.csv", i))
			os.Rename(path, newPath)
		}
	}
}
