package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	tmp := <-ch
	fmt.Println(tmp)
}
