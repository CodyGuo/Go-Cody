package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")
	one := re.Find([]byte(a))
	fmt.Println("find", string(one))

	all := re.FindAll([]byte(a), -1)
	fmt.Println("findall", all)
	for _, a := range all {
		fmt.Println("all", string(a))
	}
}
