package main

import (
	"fmt"
)

func even(a int) (array []int) {
	for i := 0; i < a; i++ {
		if i&1 == 0 {
			array = append(array, i)
		}
	}

	return
}

func main() {
	fmt.Printf("0~n之间的偶数: %v\n", even(100))
}
