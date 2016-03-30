package main

import (
	"fmt"
)

func main() {
	test := []string{"hello"}
	fmt.Println(test)
	test = append(test, "hi")
	fmt.Println(test)
}
