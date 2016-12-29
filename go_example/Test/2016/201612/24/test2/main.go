package main

import "fmt"

func main() {
	var p struct {
		Name string
		Age  int
	}
	p.Name = "codyguo"
	p.Age = 25

	var nums struct {
		first  int
		second int
	}
	nums.first = 1
	nums.second = 2
	fmt.Println(p, nums)

	var s []interface{}
	s = append(s, p, nums)
	fmt.Println(s)
}
