package main

import (
	"fmt"
)

func main() {
	a, b := order(2, 7)
	fmt.Println(a, b)
}

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}

	return a, b
}
