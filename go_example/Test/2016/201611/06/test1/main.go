package main

import (
	"fmt"
	"strings"
)

func main() {
	printSplit("你好啊")
	printSplit("编译中")
	printSplit("打印")
	printSplit("title")
}

func printSplit(title string) {
	t := 0
	for range title {
		t++
	}
	titleLen := t
	fmt.Printf("%s len %d\n", title, titleLen)
	sumLen := 50
	splitLen := sumLen - titleLen

	split2 := splitLen / 2

	splitStr := strings.Repeat("#", split2)
	splitSub := strings.Repeat("#", split2-1)
	if splitLen%2 == 0 {
		fmt.Printf("%s %s %[1]s\n", splitStr, title)
	} else {
		fmt.Printf("%s  %s %s\n", splitStr, title, splitSub)
	}
}
