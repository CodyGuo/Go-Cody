package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world.")
	f, _ := os.Create("test.log")
	defer f.Close()
}
