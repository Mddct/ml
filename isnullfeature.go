package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	paths := make([]string, 0, 100)
	filepath.Walk("dataProcess", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	var wg sync.WaitGroup
	// fmt.Println(paths[1:]) // ommit the .DS_store file
	for i := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			// fmt.Println("%s is processing", path)
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			csvF := csv.NewReader(file)
			line := 1
			for {
				records, err := csvF.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				for i := range records {
					if records[i] == "" {
						log.Fatalf("%s (%dcol , %d)", path, line, i)
					}
				}
				line++
			}

		}(paths[i])
	}
	wg.Wait()
}
