package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8888")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))

	buf := make([]byte, 1<<16)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	// result, _ := ioutil.ReadAll(conn)
	// fmt.Println(string(result))
}
