package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	for i := 0; i < 100; i++ {
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Write([]byte("\r\n"))
		time.Sleep(1 * time.Second)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
