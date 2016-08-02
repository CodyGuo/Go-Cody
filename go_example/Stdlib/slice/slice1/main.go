package main

import (
    "fmt"
)

func main() {
    var s1 []int
    s1 = []int{1, 2, 3, 4, 5}
    fmt.Println(s1, []int{1, 2, 3, 4, 5}[2])
}
