package router

import (
	"github.com/CodyGuo/Go-Cody/go_example/Web/thinkgo/2017/02/09/test1/handler"
	"github.com/CodyGuo/Go-Cody/go_example/Web/thinkgo/2017/02/09/test1/middleware"
	"github.com/henrylee2cn/thinkgo"
)

// Route register router in a tree style.
func Route(frame *thinkgo.Framework) {
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/", handler.Index),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
	)
}
