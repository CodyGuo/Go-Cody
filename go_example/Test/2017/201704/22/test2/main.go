package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	if v != nil {
		var v = 123
		fmt.Printf("if --> %v\n", v)
	}
	fmt.Printf("main --> %v\n", v)
}
