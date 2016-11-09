package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai","serverIP":"10.10.1.1"},{"serverName":"Bejing","serverIP":"10.10.2.1"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
