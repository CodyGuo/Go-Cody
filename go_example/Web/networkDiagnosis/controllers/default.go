package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	cf := &session.ManagerConfig{
		CookieName:      "networktools",
		CookieLifeTime:  3600,
		Gclifetime:      3600,
		EnableSetCookie: true,
		Secure:          true,
	}
	globalSessions, _ = session.NewManager("memory", cf)
	go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	w, r := c.Ctx.ResponseWriter, c.Ctx.Request
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	c.Data["Host"] = r.Host
	c.TplName = "index.tpl"
	log.Println("sessionID -->", sess.SessionID())
}
