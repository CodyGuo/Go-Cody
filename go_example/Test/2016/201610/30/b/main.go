package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("loop...")
		time.Sleep(2 * time.Second)
	}
}
