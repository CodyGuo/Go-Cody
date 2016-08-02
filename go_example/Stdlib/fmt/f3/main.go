package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 12; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%-2dx %-2d = %-3d  ", j, i, i*j)
			if i == j {
				fmt.Print("\n")
			}
		}
	}
}
