package main

import (
    "fmt"
    "os"
)

func main() {
    // 获取系统名字
    fmt.Println(os.Hostname())
    // 获取系统内存
    fmt.Println(os.Getpagesize())
    // 获取系统环境变量
    for index, env := range os.Environ() {
        fmt.Println(index, " : ", env)
    }
    // 获取指定key的环境变量,环境变量不区分大小写
    fmt.Println("当前系统目录为:", os.Getenv("windir"))
    // 设置环境变量
    fmt.Println("cody的环境变量为:", os.Getenv("cody"))
    os.Setenv("Cody", "guo")
    fmt.Println("cody的环境变量为:", os.Getenv("cody"))
    // 删除所有环境变量
    os.Clearenv()
    fmt.Println(os.Environ())

    // 如果存在os.Exit()就不会执行defer
    // defer fmt.Println("我在退出吗?")
    // os.Exit(0)
    fmt.Println("程序已退出,不打印了...")

    fmt.Println(os.Getuid(), os.Getgid())
    fmt.Println(os.Getgroups())
    fmt.Println(os.Getpid(), os.Getppid())

    fmt.Println(os.TempDir())

}
