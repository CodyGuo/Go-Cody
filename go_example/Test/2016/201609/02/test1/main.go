package main

import (
	"fmt"
)

func main() {
	s := []byte("a")
	fmt.Printf("%d\n", s[0]%10)
}
