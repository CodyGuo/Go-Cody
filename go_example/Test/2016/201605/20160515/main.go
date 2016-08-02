package main

import (
	"fmt"
)

type Str string

func (s Str) String() string {
	return fmt.Sprintf("Str: %s", string(s))
}

func main() {
	var s Str = "hi"
	fmt.Println(s)
}
