package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usages: %s [string]\n", os.Args[0])
	} else if ok, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); ok {
		fmt.Printf("%s is number.\n", os.Args[1])
	} else {
		fmt.Printf("%s is not number.\n", os.Args[1])
	}
}
