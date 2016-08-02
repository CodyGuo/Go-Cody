package main

import (
	"fmt"
)

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {

	var str string
	str = "helloolleh"
	ok := IsPalindrome(str)
	fmt.Printf("%s %t\n", str, ok)
}
