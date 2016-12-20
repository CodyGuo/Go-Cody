package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	// "io"
	"log"
	"net/http"
	// "strconv"
	// "time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "hello cody.guo.")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methmod:", r.Method)
	if r.Method == http.MethodGet {
		// crutime := time.Now().Unix()
		h := md5.New()
		// io.WriteString(w, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("static/tpl/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token ok")
		} else {
			fmt.Println("token err")
		}

		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
