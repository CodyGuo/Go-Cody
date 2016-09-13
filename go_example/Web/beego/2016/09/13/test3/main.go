package main

import (
	"github.com/astaxie/beego"
	"net/http"
)

func main() {
	beego.Handler("/", http.StripPrefix("/views/", http.FileServer(assetFS())))
	beego.Run()
}
