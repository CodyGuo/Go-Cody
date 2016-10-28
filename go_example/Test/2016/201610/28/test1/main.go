package main

import (
	"fmt"
	// "github.com/codyguo/win"
	"net"
	"syscall"
)

import (
	"github.com/StackExchange/wmi"
)

func main() {
	// syscall.GetHostByName("codyguo")
	const maxSiz = 50
	var n uint32 = maxSiz
	buf := make([]uint16, maxSiz)
	syscall.GetComputerName(&buf[0], &n)
	name := syscall.UTF16ToString(buf)
	fmt.Println(name)
	tmp, _ := net.InterfaceByName("以太网")
	tmp2, _ := net.InterfaceByIndex(3)

	inters, _ := net.Interfaces()

	fmt.Println(inters)
	fmt.Println(inters[0])
	fmt.Println(tmp2)
	fmt.Println(tmp)
	type s struct {
		Description string
		SettingID   string
		MACAddress  string
		// ArpAlwaysSourceRoute string
		// ArpUseEtherSNAP  int
		// DefaultIPGateway []string
		// DefaultTOS uint8
		// DHCPEnabled      int
	}
	var dst []s
	err := wmi.Query("SELECT * FROM Win32_NetworkAdapterConfiguration", &dst)
	if err != nil {
		fmt.Println("Expected err field mismatch", err)
		return
	}
	for _, s := range dst {
		fmt.Println(s)
	}
}
