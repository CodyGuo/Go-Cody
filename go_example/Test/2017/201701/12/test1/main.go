package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	root := os.Args[1]
	files := walkFiles(root)
	for file, suf := range files {
		fmt.Printf("%s -> %s\n", file, suf)
	}
}

func walkFiles(root string) map[string][]string {
	var files = make(map[string][]string)
	filepath.Walk(root, func(curPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		suf := path.Ext(curPath)
		if suf != "" {
			files[suf] = append(files[suf], curPath)
		}
		return nil
	})
	return files
}
