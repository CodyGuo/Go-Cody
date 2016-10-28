package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Repeat("#", 50))

	now := time.Now()
	d, _ := time.ParseDuration("-24h")
	// 10天前
	fmt.Println(now.Add(d * 10))
}
