package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path := "/Users/joseph"
	readDir(path)
	walk(path)
}

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf(err.Error())
	}
	for _, file := range files {
		log.Printf("%s", file.Name())
	}
}

func walk(path string) {
	var files []string
	err := filepath.Walk(path, walkAllFilesHelper(&files))
	if err != nil {
		log.Printf(err.Error())
	}
	for _, file := range files {
		log.Printf("%s", file)
	}
}

func walkAllFilesHelper(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.Mode().IsDir() {
			log.Printf("Walker %s", path)
			*files = append(*files, path)
		}
		return nil
	}
}
