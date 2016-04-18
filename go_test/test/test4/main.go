package main

import (
	"fmt"
	"sync"
)

type A struct {
	b string
}

func (a *A) update() {
	t := []A{
		A{b: "1"},
		A{b: "2"},
	}
	*a = t[0]
	fmt.Println(a)
}

func main() {
	// a := new(A)
	// a.update()
	// fmt.Println(a)

	go func() {
		i := 1
		fmt.Println(i)

	}()

	for {
		i := 0
		fmt.Println(i)
	}
}
