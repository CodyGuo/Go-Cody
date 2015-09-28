package main

import (
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    fmt.Println(os.Hostname())
    // name, _ := os.Hostname()
    // fmt.Println(net.LookupHost(name))

    conn, err := net.Dial("udp", "10.10.3.100:80")
    if err != nil {
        fmt.Println(err)
    }

    defer conn.Close()
    fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])

    var tmp string
    fmt.Scan(&tmp)
}
