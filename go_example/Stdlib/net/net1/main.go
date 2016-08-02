package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "10.10.3.227:6002")
	if err != nil {
		log.Fatalln(err)
	}

	rAddr := conn.RemoteAddr()
	fmt.Println("rAddr:", rAddr)

	var tmp string
	fmt.Scan(&tmp)
}
