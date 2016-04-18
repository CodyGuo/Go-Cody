package main

import "fmt"

func main() {
	fmt.Printf("%d \t %[1]b \t %#X \t %q\n", 42, 42, "测试")
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
