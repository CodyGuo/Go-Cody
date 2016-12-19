package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()

	for {
		select {
		case c1 <- 2:
			fmt.Println("send 2")
		case v := <-c1:
			fmt.Println("recve c1", v)
		default:
			println("default")
		}
		time.Sleep(1 * time.Second)
	}
}
