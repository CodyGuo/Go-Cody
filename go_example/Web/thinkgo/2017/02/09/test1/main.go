package main

import (
	"github.com/CodyGuo/Go-Cody/go_example/Web/thinkgo/2017/02/09/test1/router"
	"github.com/henrylee2cn/thinkgo"
)

func main() {
	router.Route(thinkgo.New("test1"))
	thinkgo.Run()
}
