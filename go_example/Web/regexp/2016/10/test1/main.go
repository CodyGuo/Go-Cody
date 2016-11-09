package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIP(ip string) (b bool) {
	if m, err := regexp.MatchString("^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-5]{2}).){3}.([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-5][0-4])$", ip); !m {
		fmt.Println(err)
		return false
	}
	return true
}

func main() {

	ip := os.Args[1]
	fmt.Println(IsIP(ip))
}
