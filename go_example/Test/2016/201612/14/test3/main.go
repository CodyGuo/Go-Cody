package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("name")
	t, _ = t.Parse("hello {{.UserName}}")
	p := Person{UserName: "codyguo"}
	t.Execute(os.Stdout, p)
}
