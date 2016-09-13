package router

import (
    "github.com/lessgo/lessgo"

    "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/sys_handler/admin"
    "github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test5/sys_handler/admin/login"
)

func init() {
    lessgo.Root(
        lessgo.Branch("/admin", "后台管理",
            lessgo.Leaf("/index", admin.Index),
            lessgo.Branch("/login", "后台登陆",
                lessgo.Leaf("/", login.Index),
            ),
        ),
    )
}
