package main

import "fmt"

func deferfmt(i int) {
	defer func() {
		fmt.Println("i为: ", i)
	}()
	if i > 10 {
		i = 1
	} else {
		i = 0
	}
}

func main() {
	deferfmt(9)
}
