package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var f1 float32
	f1 = 1.2

	fmt.Println(f1)

	var uintPtr uintptr
	uintPtr = uintptr(unsafe.Pointer(&f1))
	fmt.Println(uintPtr)

	tmp := *(*[512]uint32)(unsafe.Pointer(uintPtr))
	fmt.Println(tmp)

	fmt.Println(math.Float32frombits(tmp[5]))

	fmt.Println(unsafe.Sizeof(uintPtr))

}
