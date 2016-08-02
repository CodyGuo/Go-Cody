package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	panic("error")
	defer fmt.Println("three")

}
