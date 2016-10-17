package main

import (
	"fmt"
)

func main() {
	var str string = "abc你好"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
	for _, s := range str {
		fmt.Printf("%c ", s)
	}
	fmt.Println()
}
