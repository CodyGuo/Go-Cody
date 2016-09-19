package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"text/template"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.HandlerFunc("GetVip")
}

func GetVip(w http.ResponseWriter, r *http.Request) {
	index, _ := viewsIndexTplBytes()
	t, err := template.New("index").Parse(string(index)) //解析模板文件

	t.Execute(w, nil)
}
