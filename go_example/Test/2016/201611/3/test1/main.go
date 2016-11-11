package main

import (
	"fmt"
)

var _VERSION_ = "unknown"

func switchTest(n int) int {
	switch n {
	case 0:
		return 0
	default:
		return n + 1
	}
}

func ifTest(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + 1
	}
}

func main() {
	fmt.Println("version:", _VERSION_)
}
