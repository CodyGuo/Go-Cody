package main

import (
	"fmt"
)

func main() {
	callback(printit)
}

func callback(f func(...interface{}), y ...interface{}) {
	f(y...)
}

func printit(args ...interface{}) {
	if len(args) < 1 {
		fmt.Println("空")
		return
	}

	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Print("\n")
}
