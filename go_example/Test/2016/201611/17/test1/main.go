package main

import (
	"fmt"
	// "syscall"
	"unsafe"
)

type Point struct {
	x int
	y int
	z bool
	c byte
}

func main() {
	var p Point
	p.x = 10
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(uint32(unsafe.Sizeof(p)))
	fmt.Println(unsafe.Sizeof(&p))
}
