package main

import (
	"fmt"
)

func main() {
	var x int = -1
	var u uint = 2147483648

	fmt.Printf("x = %d = %d\n", x, uint(x))
	fmt.Printf("u = %d = %d\n", u, int(u))
}
