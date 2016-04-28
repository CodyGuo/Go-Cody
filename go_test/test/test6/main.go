package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "http://10.10.3.227/index.jsp?ip=20.20.20.33&mac=20-20-20-33-10-13&toDo=1&asmType=1&natType=0&toUrl=union.click.jd.com&eth0_mac=b0-51-8e-03-9c-93&netapp=1&userAgent=2&netApp_error=0"
	ipTemp := strings.Split(str, "=")[1]
	macTemp := strings.Split(str, "=")[2]
	ip := strings.Split(ipTemp, "&")[0]
	mac := strings.Split(macTemp, "&")[0]
	fmt.Println(ip, mac)
}
