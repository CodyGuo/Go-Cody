package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(assetFS())))

	http.HandleFunc("/", web)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}
