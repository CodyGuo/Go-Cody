package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS, runtime.NumCPU())
	fmt.Println(runtime.Version())
}
