package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"log"
)

type Win32_NetworkAdapter struct {
	Name            string
	NetConnectionID string
}

func main() {
	var dst []Win32_NetworkAdapter
	q := wmi.CreateQuery(&dst, "WHERE NetConnectionID != null")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dst {
		fmt.Printf("[%d] : Name:%s NetConnectionID:%s\n", i, v.Name, v.NetConnectionID)
	}
}
