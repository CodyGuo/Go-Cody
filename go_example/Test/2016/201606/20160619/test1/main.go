package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	pause()
	fmt.Println("world.")
}

func pause() {
	fmt.Print("请输入回车继续...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
	}
}
