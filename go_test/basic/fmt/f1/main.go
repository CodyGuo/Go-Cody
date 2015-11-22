package main

import (
    "fmt"
)

func main() {
    var a int
    var f float32
    var str string

    fmt.Println("Please input: ")
    fmt.Scanf("%d, %f, %s", &a, &f, &str)

    fmt.Println("Output: ", a, f, str)
}
