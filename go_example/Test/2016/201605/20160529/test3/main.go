package main

import (
	"fmt"
)

func main() {
	str := []string{}

	for i := 1; i < 10; i++ {
		str = append(str, "h"+fmt.Sprint(i))
	}
	fmt.Println(str)
}
