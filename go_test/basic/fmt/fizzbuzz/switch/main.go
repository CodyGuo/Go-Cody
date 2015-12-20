package main

/*
real    0m56.954s
user    0m1.394s
sys     0m1.614s
*/
import (
	"fmt"
)

func main() {
	const (
		FIZZ = 3
		BUZZ = 5
	)

	for i := 1; i <= 1000000; i++ {
		switch {
		case i%FIZZ == 0, i%BUZZ == 0:
			fmt.Printf("[%d] FizzBuzz", i)
		case i%FIZZ == 0:
			fmt.Printf("[%d] Fizz", i)
		case i%BUZZ == 0:
			fmt.Printf("[%d] Buzz", i)
		default:
			fmt.Printf("%v", i)
		}
		fmt.Print("\n")
	}
}
