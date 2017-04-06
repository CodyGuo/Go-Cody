package main

import (
	"fmt"
)

func main() {
	var (
		a   int
		f   float32
		str string
	)

	fmt.Println("准备录入数据: ")
	fmt.Scanf("%d,%f,%s", &a, &f, &str)
	fmt.Println("输出录入结果: ")
	fmt.Println(a, f, str)

}
