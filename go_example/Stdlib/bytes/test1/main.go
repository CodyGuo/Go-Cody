package main

import (
	"bytes"
	"fmt"
)

func main() {
	var hello string = "hello world."

	fmt.Println(bytes.Contains([]byte(hello), []byte("hll")))
}
