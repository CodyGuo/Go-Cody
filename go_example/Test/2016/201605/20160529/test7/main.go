package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		fmt.Println(info.Name())

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	getFileList("./")
}
