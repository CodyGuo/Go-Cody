package router

import (
    "github.com/lessgo/lessgo"

    "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/biz_handler/home"
    "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/middleware"
)

func init() {
    lessgo.Root(
        lessgo.Leaf("/websocket", home.WebSocket, middleware.ShowHeader),
        lessgo.Branch("/home", "前台",
            lessgo.Leaf("/index", home.Index, middleware.ShowHeader),
        ).Use(middleware.Print),
    )
}
