package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	checkErr(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\r\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
