package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:/n%s 110.10.25.0/24/n", os.Args[0])
		return
	}
	_, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		fmt.Printf("ParseIP Error:/n%s/n", err)
		return
	}
	intfs, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Get Addr error:/n%s/n", err)
		return
	}
	_, LanIP10, _ := net.ParseCIDR("10.0.0.0/8")
	_, LanIP172, _ := net.ParseCIDR("172.18.0.0/16")
	_, LanIP192, _ := net.ParseCIDR("192.168.0.0/16")
	for _, intf := range intfs {
		list, _ := intf.Addrs()
		for _, addr := range list {
			if !strings.Contains(addr.String(), ":") {
				list := strings.Split(addr.String(), "/")
				if len(list) != 2 {
					continue
				}
				i := net.ParseIP(list[0])
				if i == nil {
					continue
				}
				if ipnet.Contains(i) {
					fmt.Printf("WanIP: %s,oldname:%s/n", i, intf.Name)
					fmt.Println("start rename to WAN")
					rename(intf.Name, "WAN")
					continue
				}
				if LanIP10.Contains(i) || LanIP172.Contains(i) || LanIP192.Contains(i) {
					fmt.Printf("LanIP: %s,oldname:%s/n", i, intf.Name)
					fmt.Println("start rename to LAN")
					rename(intf.Name, "LAN")
				}
			}
		}
	}
}
func rename(oldname, newname string) {
	cmd := exec.Command("netsh", "interface", "set", "interface", "name", "=", oldname, "newname", "=", newname)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
