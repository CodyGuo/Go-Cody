package main

import (
	"fmt"
	"time"
)

func hello(in <-chan string) {
	for {
		select {
		case msg := <-in:
			fmt.Println(msg)
		default:
			fmt.Println("没有数据")
			time.Sleep(1 * time.Second)
		}
	}
}

func world() {
	in := make(chan string)
	go hello(in)
	for i := 0; i < 100; i++ {
		in <- fmt.Sprintf("hello %d", i)
	}
}

func main() {
	world()
}
