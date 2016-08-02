package main

import (
    "fmt"
)

func main() {
    func() {
        fmt.Println("我是一个闭包函数.")
    }()

    f := func(value string) int {
        fmt.Println(value)
        return len(value)
    }
    result := f("中文长度3")
    fmt.Println("输入的字符长度:", result)

    Call(f("爱"))
}

func Call(pFunc interface{}) {
    pFunc()
}
