package main

import (
    "github.com/lessgo/lessgo"
    "github.com/lessgo/lessgoext/swagger"

    _ "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/middleware"
    _ "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/router"
)

func main() {
    // 开启自动api文档
    // 参数为true表示自定义允许访问的ip前缀
    // 参数为false表示只允许局域网访问
    swagger.Reg(false)

    // 指定根目录URL
    lessgo.SetHome("/home")
    
    // 开启网络服务
    lessgo.Run()
}
