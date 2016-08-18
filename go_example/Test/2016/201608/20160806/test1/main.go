package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Errno(1)
	fmt.Println(err)
}
