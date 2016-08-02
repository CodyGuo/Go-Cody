package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	stack := list.New()

	var input string
	var sum int
	var stnum, conum float64 = 0, 2

	fmt.Println("请输入一段二进制数字:")
	fmt.Scanf("%s", &input)
	for _, c := range input {
		stack.PushBack(c)
	}

	length := stack.Len()
	fmt.Println("栈的当前容量是 ", length)

	// 出栈
	for e := stack.Back(); e != nil; e = e.Prev() {
		v := e.Value.(int32)
		sum += int(v-48) * int(math.Pow(conum, stnum))
		stnum++
	}

	fmt.Println("二进制转换为十进制结果是：", sum)

}
