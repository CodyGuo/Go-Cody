package main

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    mycall = syscall.NewCallback(CallBack)
)

func MyPrint() uintptr {
    fmt.Println("myprint")
    pfunc()
}

func CallBack(cfun func()) uintptr {
    MyPrint(CallBack)
    return 0
}

func MyName() {
    fmt.Println("My name is cody.guo.")
}

func main() {
    CallBack(MyName)
}
