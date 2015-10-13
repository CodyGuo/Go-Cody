package main

import (
    "fmt"
)

func Recv(ch <-chan int, lock chan<- bool) {
    for value := range ch {
        fmt.Println(value)
    }

    // lock channel
    lock <- true
}

func Send(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }

    close(ch)
}

func main() {
    ch := make(chan int)
    lock := make(chan bool)

    go Send(ch)
    go Recv(ch, lock)

    // unlock channel
    <-lock
}
