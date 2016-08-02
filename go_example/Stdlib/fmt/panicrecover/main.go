package main

import (
	"fmt"
)

func main() {
	ok := throwsPanic(myPanic)
	if ok {
		fmt.Println("有没有搞错，复活了。", ok)
	} else {
		fmt.Println("没有崩溃啊。")
	}
}

func myPanic() {
	panic("哎呀，我操，又崩溃了。")
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("x:", x)
			b = true
		}
	}()

	f()

	return
}
