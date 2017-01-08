package main

import "fmt"

func f(left, right chan int) {
	r := <-right
	left <- 1 + r
	fmt.Printf("(left: %v -> %d) <- (right: %v -> 1+  %d)\n", left, 1+r, right, r)
}

func main() {
	const n = 10
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	fmt.Printf("left: %v, right: %v, leftmost: %v\n", left, right, leftmost)
	for i := 0; i < n; i++ {
		right = make(chan int)
		fmt.Printf("left = right, %v = %v\n", left, right)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		fmt.Printf("go right -> %v\n", c)
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}
