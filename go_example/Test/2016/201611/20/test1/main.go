package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(file, "hello %s\n", "world.")
	file.Close()
}
