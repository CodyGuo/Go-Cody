package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := os.Args[1]
	base := filepath.Base(path) // base
	fmt.Println(base)
	p, f := filepath.Split(path)
	fmt.Printf("%s -> %s\n", p, f)

	es, _ := filepath.EvalSymlinks(path)
	fmt.Println(es)

	fmt.Println(filepath.Ext(path))

	ext := filepath.Ext(path)
	fmt.Println(strings.Split(base, ext)[0])
}
