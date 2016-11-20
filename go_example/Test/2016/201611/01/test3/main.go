package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "hello world"
	f := bytes.Fields([]byte(str))
	fmt.Println(f)

	for i := range f {
		fmt.Println(string(f[i]))

	}
}
