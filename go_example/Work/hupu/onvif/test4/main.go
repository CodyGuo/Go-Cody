package main

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

type Login struct {
	Username string
	Password string
	Nonce64  string
	Created  string
}

func main() {
	t := template.New("tmp")
	t.Parse(loginXML)

	login := Login{
		Username: "admin",
		Password: "admin2016",
		Nonce64:  "noce",
		Created:  "2016-11-25T03:28:16Z",
	}

	// buf := make([]byte, 1<<16)
	var buf bytes.Buffer
	t.Execute(&buf, login)
	result := buf.String()

	result = strings.Replace(result, "&lt;", "<", -1)
	fmt.Println(result + deviceInfoXML)
}
