package main

import (
	"fmt"
)

func main() {
	var p *int
	fmt.Printf("%v\n", p)

	var i int
	p = &i
	fmt.Printf("%v\n", p)
	*p = 8
	*p++
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v\n", i)

	var ps *[]int = new([]int)
	var v []int = make([]int, 100)

	// *ps = make([]int, 100, 100)
	fmt.Println(*ps)
	fmt.Println("---------------------")
	fmt.Println(v)
}
