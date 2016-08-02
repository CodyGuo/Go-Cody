package main

import (
	"fmt"
)

func main() {
	var n int
Loop:
	if n < 10 {
		n++
		goto Loop
	}

	fmt.Println("n:", n)
}
