package main

import (
	"fmt"
)

func main() {
	var a = map[int]string{1: "a", 2: "b"}
	var b = map[int]string{1: "c", 2: "d", 3: "e"}
	var c = make(map[int][]string)
	for k, v := range a {
		c[k] = append(c[k], v)
	}
	for k, v := range b {
		c[k] = append(c[k], v)
	}

	fmt.Printf("a --> %v\n", a)
	fmt.Printf("b --> %v\n", b)
	fmt.Printf("c --> %v\n", c)

}
