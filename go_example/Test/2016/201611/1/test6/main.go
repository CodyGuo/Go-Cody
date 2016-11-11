package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string

	fmt.Scan(&input)

	fmt.Println(input)
	n, _ := fmt.Scanln()
	fmt.Println(n)
	bufio.NewWriter(os.Stdin).Reset(nil)
	fmt.Scan(&input)
	fmt.Println(input)

}
