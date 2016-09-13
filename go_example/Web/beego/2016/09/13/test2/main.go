package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe("127.0.0.1", nil); err != nil {
		log.Fatal(err)
	}
}
