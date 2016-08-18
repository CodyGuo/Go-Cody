package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("%c\r", r)
			time.Sleep(1 * time.Second)
		}
	}
}
