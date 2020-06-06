package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func walk(dir string, size chan<- int64) {
	dirFileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir err:", err)
	}
	for _, v := range dirFileInfo {
		if v.IsDir() {
			walk(filepath.Join(dir, v.Name()), size)
		} else {
			size <- v.Size()
		}
	}
}
func main() {
	dir := flag.String("dir", ".", "dirname")
	flag.Parse()
	fileSize := make(chan int64)
	go func() {
		walk(*dir, fileSize)
		close(fileSize)
	}()
	var size int64
	for val := range fileSize {
		size += val
	}
	fmt.Println("dir size:", size/1000000)
}
