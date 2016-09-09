package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, "hello world. version 1."); err != nil {
		log.Fatal(err)
	}
}
