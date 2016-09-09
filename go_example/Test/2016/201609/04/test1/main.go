package main

import (
	"fmt"
)

func main() {
	s := "hello world."
	fmt.Println(string(s[1]))
	b := []byte(s)
	fmt.Println(string(b[2]))
}
