package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 2
	c <- 1
	fmt.Println(<-c)
	fmt.Println(<-c)
}
