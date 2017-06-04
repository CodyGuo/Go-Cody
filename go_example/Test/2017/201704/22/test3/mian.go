package main

import "fmt"

type Name struct {
	Name string
	Age  int
}

func main() {
	var list []int
	var name Name
	fmt.Printf("list --> %#v, name --> %#v\n", list, name)
}
