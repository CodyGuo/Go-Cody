package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 3}
	fmt.Println(r1, area(r1))
}
