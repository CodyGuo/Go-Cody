/*
   翻转字符串
*/
package main

import (
	"fmt"
)

func main() {
	str := "hello world."

	fmt.Print("方式1: ")
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[len(str)-i-1]))
	}
	fmt.Print("\n")

	fmt.Print("方式2: ")
	for i := len(str) - 1; i > -1; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Print("\n")

	defer fmt.Print("\n")
	for i := 0; i < len(str); i++ {
		defer fmt.Print(string(str[i]))
	}
	defer fmt.Print("方式3: ")
}
