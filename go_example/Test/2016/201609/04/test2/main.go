package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   1    2       3         4    one two  测试  张三      李四     "

	fmt.Println(filterSpace(str))
}

func filterSpace(str string) []string {
	const space = ' '
	var ok bool
	mapping := func(r rune) rune {
		switch r {
		case space:
			if ok {
				return -1
			}
			ok = true
		default:
			ok = false
		}
		return r
	}

	mapS := strings.Map(mapping, str)
	return strings.Split(strings.TrimSpace(mapS), " ")
}
