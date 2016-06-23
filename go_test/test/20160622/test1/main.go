package main

import (
	"fmt"
)

func easySum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}

	return sum
}

func gaosiSum1(num int) int {
	sum := (1 + num) * num / 2
	// return (1 + num) * num / 2
	return sum
}

func gaosiSum2(num int) int {
	return (1 + num) * num / 2
}

func main() {
	fmt.Println(easySum(100))
	fmt.Println(gaosiSum1(100))
	fmt.Println(gaosiSum2(100))
}
