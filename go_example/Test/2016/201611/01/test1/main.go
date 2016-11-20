package main

import (
	"fmt"
	"net"
	"os"
	"reflect"
	"time"
)

func timeSum(sum func()) {
	now := time.Now()
	sum()
	fmt.Println(time.Now().Sub(now).String())
}

func getIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}

	fmt.Println(addrs[1])

}

func main() {
	timeSum(getIP)
}
