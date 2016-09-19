package main

import (
	"fmt"
	"os"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	os.Stdin.Read(make([]byte, 1))
	close(done)
	fmt.Println("done")
}
