package main

import (
    "fmt"
)

func test(ch chan bool) {
    fmt.Println("Go..")
    ch <- true
}

func main() {
    chs := make(chan bool)
    go test(chs)

    // go func() {
    //     for i := 0; i < 5; i++ {
    //         chs <- i
    //     }
    //     close(chs)
    // }()

    // for value := range chs {
    //     fmt.Println(value)
    // }
    <-chs
    go func() {
        chs <- false
        fmt.Println("wirte 1")
    }()

    value, ok := <-chs
    if ok {
        fmt.Println("Done..", value)
    }

    fmt.Println("me...")
}
