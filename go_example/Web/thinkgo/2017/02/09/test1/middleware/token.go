package middleware

import (
	"github.com/henrylee2cn/thinkgo"
)

/*
Token
*/
var Token = thinkgo.HandlerFunc(func(ctx *thinkgo.Context) error {
	ctx.Log().Debugf("[ware] token:%q", ctx.QueryParam("token"))
	return nil
})
