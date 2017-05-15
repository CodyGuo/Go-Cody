package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var StaticDir = "/static"

func dealStaticFiles(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, StaticDir) {
		file := stDir + r.URL.Path
		fmt.Println(file)
		f, err := os.Open(file)
		defer f.Close()

		if err != nil && os.IsNotExist(err) {
			fmt.Fprintln(w, "File not exist")
			return
		}
		http.ServeFile(w, r, file)
		return
	} else {
		fmt.Fprintln(w, "Hello world")
	}
}
func main() {
	http.HandleFunc("/", dealStaticFiles)    //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
