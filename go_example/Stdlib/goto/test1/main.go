package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(3 * time.Second)

	go func() {
		select {
		case <-t:
			println("已超时")
		}
	}()
	fmt.Println("等待超时中...")
	time.Sleep(2 * time.Second)
	fmt.Println("正常输出...")
	fmt.Println("超时输出...")
}
