package main

import (
	"time"
)

var (
	bakPath string
)

func init() {
	bakPath = time.Now().Format(LAYOUT) + "研发部备份"
	svn = new(Svn)
	svn.loginUser = "cody.guo"
	svn.loginPasswd = "iman2015"
}
