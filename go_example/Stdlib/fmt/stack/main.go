package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i > 9 {
		return
	}

	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"

	}

	return str
}

func main() {
	// s := new(stack)
	var s stack
	for i := 0; i < 10; i++ {
		s.push(i)
	}
	s.data[9] = 100

	fmt.Println(s.data)

	for i := 0; i < 5; i++ {
		fmt.Println(s.pop())
	}
	fmt.Println(s.data)

	fmt.Printf("stack %v\n", s)
}
