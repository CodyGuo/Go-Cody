package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    for {
        var login, passwd string
        fmt.Print("h3c>")
        fmt.Scanln(&login)
        if login == "exit" {
            os.Exit(0)
        }
        if login == "" {
            continue
        }
        if login == "admin" {
            fmt.Print("passwd:")
            fmt.Scanln(&passwd)
            if passwd == "admin2015" {
                fmt.Print("h3c#\n")
                fmt.Print("h3c#恭喜您登录成功个.\n")
                break
            } else {
                fmt.Print("h3c>密码错误.\n")
                continue
            }
        } else {
            fmt.Print("h3c>输入指令错误.\n")
            continue
        }

    }

    for {
        bash := exec.Command("/bin/sh", "-c", "/bin/sh")
        bash.Run()
    }

}
