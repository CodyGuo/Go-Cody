package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// START1 OMIT
	quit := make(chan bool) // HL
	c := boring("Joe", quit)
	for i := rand.Intn(100); i >= 0; i-- {
		fmt.Println(<-c)
	}
	fmt.Println("quit<-true")
	quit <- true // HL
	// STOP1 OMIT
}

func boring(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() { // HL
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Printf("boring-> quit: %v\n", <-quit)
				return
			}

		}
	}()
	return c
}
