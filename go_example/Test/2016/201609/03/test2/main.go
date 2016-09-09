package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	d, _ := time.ParseDuration("-24h")

	fmt.Println(now.Add(d))
}
