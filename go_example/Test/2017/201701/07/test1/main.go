package main

import (
	"fmt"
)

func main() {
	fmt.Println(^5)
	fmt.Println(0 ^ 5)
	fmt.Println(^4, ^3, ^2, ^1, -1^1, -1^2)
	fmt.Printf("%b %#b\n", -2, -1)
	var fs []float64 = []float64{1.1234456, 1.1234567, 1.1234678, 1.1}
	for _, f := range fs {
		s := fmt.Sprintf("%.5f", f)
		fmt.Println(f, "->", s)
	}
}
