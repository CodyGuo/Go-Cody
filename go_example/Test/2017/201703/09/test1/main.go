package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "127.0.0.1")

	cmd.Run()
	for {
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
