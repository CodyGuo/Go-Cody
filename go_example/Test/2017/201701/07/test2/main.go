package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("gen接收到数据 %d\n", n)
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("sq接收到数据 %d\n", n*n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	c := gen(2, 3)
	out := sq(c)
	fmt.Println("out:", <-out)
	fmt.Println("out:", <-out)

	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}

	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)
	for n := range merge(c1, c2) {
		fmt.Println("merge ->", n)
	}
}
