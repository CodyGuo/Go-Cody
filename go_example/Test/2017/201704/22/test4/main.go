package main

import "fmt"

func main() {
	var v interface{}
	v = "hello world."
	switchTest(v)
	v = 1
	switchTest(v)
	v = []int{1, 2, 3}
	switchTest(v)
}

func switchTest(v interface{}) {
	switch i := v.(type) {
	case string:
		fmt.Printf("The string is %s.\n", i)
	case int:
		fmt.Printf("The intger is %d.\n", i)
	default:
		fmt.Printf("Unsuported value. (type=%T)\n", i)
	}
}
