package main

import (
	// "bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

import (
	"github.com/axgle/mahonia"
)

var system [2]string = [2]string{
	"本地连接",
	"以太网",
}

var sw bool

func init() {
	flag.BoolVar(&sw, "on", true, "enable")
}

func main() {
	flag.Parse()
	// err := disable(sw)
	// fmt.Println(err)
	name := checkInterface()
	fmt.Printf("网卡 [%s] 状态正常.\n", name)

	inter, _ := net.InterfaceByName(name)

	ip, _ := inter.Addrs()

	fmt.Println(ip[len(ip)-1])
	// fmt.Println(ip)
	for {
	}
}

func getAdapterList() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}

func localAddresses() error {
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	aList, err := getAdapterList()
	if err != nil {
		return err
	}

	for _, ifi := range ifaces {
		for ai := aList; ai != nil; ai = ai.Next {
			index := ai.Index

			if ifi.Index == int(index) {
				ipl := &ai.IpAddressList
				gwl := &ai.GatewayList
				for ; ipl != nil; ipl = ipl.Next {
					fmt.Printf("%s: %s (%s) %s\n", ifi.Name, ipl.IpAddress, ipl.IpMask, gwl.IpAddress)
				}
			}
		}
	}

	return err
}

func checkInterface() string {
	for win := range system {
		_, err := net.InterfaceByName(system[win])
		if err == nil {
			return system[win]
		}
	}
	name, err := enableInterface()
	if err != nil {
		fmt.Printf("网卡 [%s] 启用失败. %v\n", name, err)
		return ""
	}
	fmt.Printf("网卡 [%s] 启用成功.", name)

	return name
}

func enableInterface() (name string, err error) {
	for win := range system {
		cmd := exec.Command("cmd.exe", "/c", `netsh interface show interface name=`+system[win])
		out, _ := cmd.StdoutPipe()

		err = cmd.Start()
		if err != nil {
			fmt.Println("enableInteface", err)
			return "", err
		}
		ma := mahonia.NewDecoder("gbk")
		buf := ma.NewReader(out)
		data := make([]byte, 1<<16)
		n, _ := buf.Read(data)
		if strings.Contains(string(data[:n]), "管理状态: 已禁用") {
			name = system[win]
			fmt.Printf("网卡 [%s] 禁用状态,正在启用网卡,请稍等.\n", name)
			err = enabled(true, name)
		}

		cmd.Wait()
	}

	return name, err
}

func enabled(ok bool, name string) error {
	var on string = " disable"
	if ok {
		on = " enable"
	}
	cmd := exec.Command("cmd.exe", "/c", `netsh interface set interface `+name+on)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}
