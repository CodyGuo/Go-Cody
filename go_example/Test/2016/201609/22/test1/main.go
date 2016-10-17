package main

import (
	"fmt"
	// "reflect"
)

// func Test(v interface{}) {
// 	vv := reflect.ValueOf(v).Elem()
// 	vv.SetInt(20)
// }

func main() {
	num := 10
	fmt.Println(num)
	num = 20
	// Test(&num)
	fmt.Println(num)
}
