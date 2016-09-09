package main

import (
	"fmt"
)

func recPrint(n int) {
	if n <= 0 {
		return
	}
	fmt.Printf("%d\n", n)
	recPrint(n - 1)
}

func main() {
	recPrint(10)
}
