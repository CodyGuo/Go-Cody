package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()
		go func() {
			for {
				daytime := time.Now().String()
				conn.Write([]byte(daytime))
				conn.Write([]byte("\r\n"))
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
