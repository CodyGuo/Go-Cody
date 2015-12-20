package main

/*
real    0m55.720s
user    0m1.095s
sys     0m1.449s
*/
import (
	"fmt"
)

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	var p bool
	for i := 1; i <= 1000000; i++ {
		p = false
		if i%FIZZ == 0 {
			fmt.Printf("[%d] Fizz", i)
			p = true
		}
		if i%BUZZ == 0 {
			if p {
				fmt.Printf("Buzz [%d]", i)
			} else {
				fmt.Printf("[%d] Buzz", i)
			}
			p = true
		}
		if !p {
			fmt.Println(i)
		} else {
			fmt.Print("\n")
		}
	}
}
