package middleware

import (
    "github.com/lessgo/lessgo"
)

var Print = lessgo.ApiMiddleware{
    Name:   "打印测试",
    Desc:   "打印测试",
    Config: nil,
    Middleware: func(confObject interface{}) lessgo.MiddlewareFunc {
        return lessgo.WrapMiddleware(func(c *lessgo.Context) error {
            c.Log().Info("测试中间件-打印一些东西：1234567890")
            c.Log().Info("param:%v(len=%v),%v(len=%v)", c.PathParamKeys(), len(c.PathParamKeys()), c.PathParamValues(), len(c.PathParamValues()))
            return nil
        })
    },
}.Reg()

var ShowHeader = lessgo.ApiMiddleware{
    Name:   "显示Header",
    Desc:   "显示Header测试",
    Config: nil,
    Middleware: func(c *lessgo.Context) error {
        c.Log().Info("测试中间件-显示Header：%v", c.Request().Header)
        return nil
    },
}.Reg()
