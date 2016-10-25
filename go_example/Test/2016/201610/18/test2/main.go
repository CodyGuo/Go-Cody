package main

import (
	"fmt"
)

func main() {
	n := make([]int, 0)
	t(n, 0)
	fmt.Println(n)
	var nn = []int{1, 2}
	tt(nn)
	fmt.Println(nn)
}

func t(n []int, p int) {
	if p > 10 {
		return
	}

	n = append(n, p)
	t(n, p+1)
}

func tt(n []int) {
	n[0] = 11
}
