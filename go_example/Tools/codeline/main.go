package main

import (
	"os"
	"path/filepath"
	"strings"

	"errors"

	"github.com/codyguo/logs"
)

func main() {
	dir, err := os.Getwd()
	checkErr(err)
	if len(os.Args) == 2 {
		dir = os.Args[1]
	}
	WalkFiles(dir)
}

func WalkFiles(dir string) {
	fullPath, err := filepath.Abs(dir)
	checkErr(err)
	filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		logs.Noticef("walkFiles --> %s", path)
		if info.IsDir() {
			return nil
		}
		nowPath := filepath.Dir(path)
		if fullPath != nowPath {
			logs.Info(fullPath, nowPath)
			return filepath.SkipDir
		}
		// }
		return nil
	})
}

func checkLines(fullPath, path string) error {
	path, err := filepath.Abs(path)
	checkErr(err)
	pLists := strings.Split(path, fullPath)
	var lines []string
	if len(pLists) > 1 {
		lines = strings.Split(pLists[1], string(os.PathSeparator))
	}
	logs.Info(lines)
	if len(lines) > 2 {
		return errors.New("line to 3.")
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		logs.Fatal(err)
	}
}
