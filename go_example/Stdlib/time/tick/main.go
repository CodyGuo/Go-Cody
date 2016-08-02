package main

import "time"
import "fmt"

func main() {
	t := time.Tick(1 * time.Second)
	for now := range t {
		fmt.Printf("%v %s\n", now.Format("2006-01-02 15:04:05"), "hi")
	}
}
