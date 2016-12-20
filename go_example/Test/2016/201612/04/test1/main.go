package main

import (
	"fmt"
)

func hello(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "hello"
	}
	close(ch)
}

func world(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "world"
	}
	close(ch)
}

func main() {
	h := make(chan string)
	w := make(chan string)
	go hello(h)
	go world(w)

	for i, v := range <-h {
		fmt.Println(i, string(v))
	}
	// for value := range h {
	// 	fmt.Println(value, <-w)
	// }

}
