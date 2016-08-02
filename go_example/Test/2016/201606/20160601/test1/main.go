package main

import (
	"fmt"
	"log"
)

func add(x, y []int) (result []int) {
	if len(x) != len(y) {
		log.Fatal("The length is not equal.")
	}

	decimal := 0
	for i, j := range x {
		sum := j + y[i] + decimal
		if sum >= 10 {
			sum = sum % 10
			decimal = 1
		} else {
			decimal = 0
		}
		result = append(result, sum)
	}

	if decimal == 1 {
		result = append(result, 1)
	}
	return
}

func main() {
	result := add([]int{1, 1, 2, 1}, []int{5, 6, 7, 2})
	fmt.Println(result)
}
