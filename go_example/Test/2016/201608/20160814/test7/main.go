package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	j := make(chan time.Time)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j <- tick
	}

	for range j {
		<-j
	}
	fmt.Println("main down.")

}
