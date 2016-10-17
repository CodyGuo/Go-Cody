package main

import (
	"fmt"
)

var a string

func f(wait chan bool) {
	fmt.Println(a)
	wait <- true
}

func main() {
	wait := make(chan bool)

	a = "hello world"
	go f(wait)

	<-wait
}
