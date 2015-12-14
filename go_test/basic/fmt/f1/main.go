package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var a int
	var f float32
	var str string

	fmt.Print("Please input : ")
	n, _ := fmt.Scanf("%d %f %s", &a, &f, &str)
	fmt.Println("Output: ", a, f, str, n)

	fmt.Println("----------------------------")

	// 初始化stdin
	os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")

	var sIP, sMac string
	fmt.Print("Please input: ")
	fmt.Scanf("%s %s", &sIP, &sMac)

	fmt.Println("Output: ", sIP, sMac)
}
