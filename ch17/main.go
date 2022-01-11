package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	waitGroup = sync.WaitGroup{}
	lock = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename ) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGroup.Add(1)
			fileSearch(filepath.Join(root, file.Name()), filename)
		}

	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)

	go fileSearch("/Users/sanghoonkim/Documents/golang/src/github.com/concurrency","main.go")
	
	waitGroup.Wait()

	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}