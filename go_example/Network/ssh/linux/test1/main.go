package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		Timeout: time.Second,
	}
	start := time.Now()
	time.AfterFunc(1*time.Second, func() {
		fmt.Printf("超时 %d\n", 1)
		goto LABEL
	})
	conn, err := ssh.Dial("tcp", "10.10.2.226:22", config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	seesion, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer seesion.Close()

	seesion.Stdout = os.Stdout

	seesion.Run("who")
	fmt.Printf("ssh连接执行时间: %v\n", time.Since(start))
LABEL:
	log.Println("连接ssh超时")
}
