package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var files []string
	//src := "C:/Users/lenovo/Downloads" //copy source
	src := "D:/Go stuff"
	dst := "" // destination
	abs := "C:/Users/lenovo/Desktop"
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, file := range files {
		fmt.Println(file)
		dst = file
		wg.Add(1)
		go copy(&wg, abs, src, dst)
		if err != nil {
			panic(err)
		}
	}
	wg.Wait()

}

func copy(wg *sync.WaitGroup, abs, src, dst string) {
	_, err := os.Stat(src)
	if err != nil {
		panic(err)
	}

	// if !filestat.Mode().IsRegular() {
	// 	return fmt.Errorf("%s is not a regular file", src)
	// }

	// open source file
	source, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer source.Close()

	// create file
	destination, err := os.Create(strings.Join([]string{abs, dst}, "/"))
	if err != nil {
		panic(err)
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		panic(err)
	}
	wg.Done()
}
