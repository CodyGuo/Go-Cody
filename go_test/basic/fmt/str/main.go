package main

import (
	"fmt"
)

func main() {
	str := "hello world. 世界，你好！"

	fmt.Println("字节数组方式遍历: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %v\n", i, str[i])
	}

	fmt.Println("Unicode字符方式遍历: ")
	for i, ch := range str {
		fmt.Printf("str[%d] = %v\n", i, ch)
	}

}
