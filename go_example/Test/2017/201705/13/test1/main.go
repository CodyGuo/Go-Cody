package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	addr := "localhost:8888"
	ipAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", ipAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		checkErr(err)
		fmt.Println("A client connected : " + conn.RemoteAddr().String())
		go tcpPipe(conn)
	}
}

func tcpPipe(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("diconnected: " + ipStr)
		fmt.Println("-----------------------------------------\n\n")
	}()

	reader := bufio.NewReader(conn)
	for {
		p := make([]byte, 16<<1)
		n, err := reader.Read(p)
		if err != nil {
			return
		}

		message := fmt.Sprintf("%s", p[:n])
		message = strings.TrimSpace(message)

		fmt.Printf("recive from client: %s, message: %s\n", ipStr, message)
		switch message {
		case "hello":
			conn.Write([]byte("<nac>world你好世界.</nac>"))
		default:
			msg := time.Now().Format("2006-01-02 15:04:05")
			conn.Write([]byte("<nac>" + msg + "</nac>"))
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
