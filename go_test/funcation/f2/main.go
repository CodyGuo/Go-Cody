package main

import (
    "fmt"
)

type CallFunc func()

func CallBack(pFunc CallFunc) {
    pFunc()
}

func main() {

    // var f = func() {
    //     fmt.Println("我是f")
    // }

    CallBack(func() {
        fmt.Println("我是匿名.")
    })
}
