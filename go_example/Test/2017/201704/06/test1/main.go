package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var a = make(chan int, 2)

func A() {
	defer wg.Done()
	a <- 1
	a <- 2
	time.Sleep(time.Second * 3)
	println("A done.")
	close(a)
}
func B() {
	defer wg.Done()
	for i := range a {
		fmt.Println("hahh")
		fmt.Println(i)
	}
	println("B done.")
}
func main() {
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
