package main

import (
	"fmt"
)

func main() {
	var u uint = 4294967295
	fmt.Printf("uint: %d, int: %d, uint_b: %b, int_b: %b\n", u, int(u), u, int(u))
}
