package main

import (
	"github.com/codyguo/logs"
)

func main() {
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15}
	sum := 30
	for _, x := range numbers {
		for _, y := range numbers {
			for _, z := range numbers {
				if x+y+z == sum {
					logs.Noticef("x = %d, y = %d, z = %d", x, y, z)
				}
				logs.Errorf("x = %d, y = %d, z = %d, sum = %d", x, y, z, x+y+z)
			}
		}
	}
}
