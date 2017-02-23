package handler

import (
	"github.com/henrylee2cn/thinkgo"
)

/*
Index
*/
var Index = thinkgo.HandlerFunc(func(ctx *thinkgo.Context) error {
	return ctx.Render(200, thinkgo.JoinStatic("index.html"), thinkgo.Map{
		"TITLE":   "thinkgo",
		"VERSION": thinkgo.VERSION,
		"CONTENT": "Welcome To Thinkgo",
		"AUTHOR":  "HenryLee",
	})
})
