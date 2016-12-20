package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 23}
	clean(&arr)
	fmt.Println(arr)
}

func clean(arr *[]int) {
	// arr = arr[:0:0]
	*arr = []int{}
}
