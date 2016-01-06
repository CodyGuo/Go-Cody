package main

import (
	"fmt"
	// "time"
)

var c chan int

func ready(w string) {
	// time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	c = make(chan int)
	go ready("Tea")
	go ready("Coffee")

	fmt.Println("I'm waiting.")
	<-c
	<-c
}
