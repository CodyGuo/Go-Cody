/**
 * 替换文件内容
 */
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"
)

// 并发打开文件数限制
const numModFiles = 50

const (
	old = "github.com/lxn"
	new = "github.com/codyguo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%s path\n", os.Args)
		return
	}
	start := time.Now()
	err := modFiles(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n所有文件处理完成，用时 -> %s\n", time.Since(start))
}

func modFiles(root string) error {
	done := make(chan struct{})
	errc := make(chan error)
	defer close(done)

	files, errw := walkFiles(done, root)

	var wg sync.WaitGroup
	wg.Add(numModFiles)
	for i := 0; i < numModFiles; i++ {
		go func() {
			readWrteFiles(done, files, errc)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(errc)
	}()

	for err := range errc {
		if err != nil {
			log.Printf("[ERROR] modFiles -> errc -> %v\n", err)
		}
	}
	if err := <-errw; err != nil {
		return err
	}
	return nil
}

func readWrteFiles(done <-chan struct{}, files <-chan string, errc chan<- error) {
	for file := range files {
		err := rwFiles(file)
		select {
		case errc <- err:
		case <-done:
			return
		}
	}
}

func rwFiles(file string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_SYNC, os.ModeType)
	if err != nil {
		return err
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// fmt.Printf("src -> %s\n\n old -> %s\n", src, []byte(old))
	if bytes.Contains(src, []byte(old)) {
		fmt.Printf("正在处理文件 -> %s\n", file)
		out := bytes.Replace(src, []byte(old), []byte(new), -1)
		if _, err := f.Seek(0, 0); err != nil {
			return err
		}
		if _, err := f.Write(out); err != nil {
			return err
		}
	}
	return nil
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	files := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(files)
		errc <- filepath.Walk(root, func(file string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			if path.Ext(file) == ".go" {
				select {
				case files <- file:
				case <-done:
					return errors.New("walkFile canceld")
				}
			}

			return nil
		})
	}()
	return files, errc
}
