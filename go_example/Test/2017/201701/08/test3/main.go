package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("job", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Byte!"
	fmt.Printf("Job says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(100 * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
			case q := <-quit:
				quit <- q + " see you!"
				return
			}
		}
	}()

	return c
}
