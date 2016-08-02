package main

import (
    "fmt"
)

func MyPrint(any ...interface{}) {
    fmt.Println(any...)
}

func MycallBack(info interface{}, status bool, mfunc func(any ...interface{})) {
    if status {
        mfunc(info)
        fmt.Println("I'm callback.")
    } else {
        fmt.Println("I'm not callback.")
    }
}

func main() {
    MycallBack(false, true, MyPrint)
    fmt.Println("================分隔符====================")
    MycallBack("hello everybody.", false, MyPrint)

}
