package main

import (
	"fmt"
	"time"
)

func main() {
	natuals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			natuals <- x
		}
		close(natuals)
	}()

	go func() {
		for {
			x, ok := <-natuals
			if !ok {
				break
			}
			squares <- x * x

		}
		close(squares)
	}()

	// for x := range squares {
	// 	fmt.Println(x)
	// }
	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
