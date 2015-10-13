package main

import (
    "fmt"
    "net"
    "strings"
)

func main() {
    conn, err := net.Dial("udp", "10.10.3.100:21")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}
