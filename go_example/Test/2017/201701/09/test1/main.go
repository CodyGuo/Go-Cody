package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("aptech"),
		},
	}

	client, err := ssh.Dial("tcp", "192.168.1.188:22", &config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("pwd;ls"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.String())
}
