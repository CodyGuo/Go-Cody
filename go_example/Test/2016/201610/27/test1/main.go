package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
		regexp.Match(pattern, b)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println(os.Args[1], "是数字.")
	} else {
		fmt.Println(os.Args[1], "不是数字.")
	}
}
