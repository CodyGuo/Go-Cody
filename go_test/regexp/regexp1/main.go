package main

import (
    "fmt"
    "regexp"
)

var digitsRegexp = regexp.MustCompile(`(\d+)\D+(\d+)`)

func main() {
    someString := "1000abcd123"
    fmt.Println(digitsRegexp.FindStringSubmatch(someString))
}
