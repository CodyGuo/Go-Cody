package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func walkDir2(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir2(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

var verbose = flag.Bool("v", false, "debug true.")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir2(root, &n, fileSize)
	}
	go func() {
		n.Wait()
		close(fileSize)
	}()
	// go func() {
	// 	for _, root := range roots {
	// 		walkDir(root, fileSize)
	// 	}
	// 	close(fileSize)
	// }()

	// var nfiles, nbytes int64
	// for size := range fileSize {
	// 	nfiles++
	// 	nbytes += size
	// }

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(nbytes)/1e9)
}
