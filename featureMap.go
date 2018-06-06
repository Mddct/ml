package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var filelist = []string{
	"./Cxx/C1.txt",
	"./Cxx/C14.txt",
	"./Cxx/C15.txt",
	"./Cxx/C16.txt",
	"./Cxx/C17.txt",
	"./Cxx/C18.txt",
	"./Cxx/C19.txt",
	"./Cxx/C20.txt",
	"./Cxx/C21.txt",
}

func main() {

	var wg sync.WaitGroup
	for i := range filelist {
		wg.Add(1)
		go func(path string) {
			m := make(map[int]int)
			defer wg.Done()
			file, _ := os.Open(path)
			buf := bufio.NewScanner(file)
			for buf.Scan() {
				i, err := strconv.Atoi(buf.Text())
				if err != nil {
					panic(err)
				}
				m[i]++
			}
			file.Close()
			// ommit buf.Error()
			filew, _ := os.Create(fmt.Sprintf("%sMap.txt", strings.Split(filepath.Base(path), ".")[0]))
			defer filew.Close()

			c1 := make([]int, 0, len(m))
			for k := range m {
				c1 = append(c1, k)
			}
			sort.Ints(c1)
			for i = range c1 {
				fmt.Fprintf(filew, "%v:%d\n", c1[i], m[c1[i]])
			}
			fmt.Println(path, "DONE")

		}(filelist[i])
	}
	wg.Wait()
	fmt.Println("ALL DONE")

}
