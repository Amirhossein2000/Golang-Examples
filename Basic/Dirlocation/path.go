package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("PWD ==>> ", root)
	err := filepath.Walk(root, processpath)
	if err != nil {
		panic(err)
	}
}
func processpath(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("directory ==>", path)
		} else {
			fmt.Println("file ==>", path)
		}
	}
	return nil
}
