package main

import (
	"fmt"
	// "time"
	// "unicode/utf8"
	"unsafe"
)

import (
	"github.com/CodyGuo/win"
)

const (
	layout = "2006-01-02 15:04:05"
)

func main() {
	var AdapterInfo *win.IP_ADAPTER_INFO
	var Adapter *win.IP_ADAPTER_INFO = nil
	var ulOutBufLen uint32 = uint32(unsafe.Sizeof(AdapterInfo))

	AdapterInfo = new(win.IP_ADAPTER_INFO)
	ret := win.GetAdaptersInfo(nil, &ulOutBufLen)
	if win.ERROR_BUFFER_OVERFLOW == ret {
		AdapterInfo = &win.IP_ADAPTER_INFO{}
		AdapterInfo = new(win.IP_ADAPTER_INFO)
	}

	win.GetAdaptersInfo(AdapterInfo, &ulOutBufLen)

	Adapter = AdapterInfo
	for Adapter != nil {
		fmt.Printf("Adapter Name: %s\n", win.TrimRight(Adapter.AdapterName))
		// des := win.UTF8Decode(Adapter.Description[0:])
		fmt.Printf("Desc: %s\n", win.TrimRight(Adapter.Description))
		fmt.Printf("MAC address: %02X-%02X-%02X-%02X-%02X-%02X\n", Adapter.Address[0], Adapter.Address[1], Adapter.Address[2],
			Adapter.Address[3], Adapter.Address[4], Adapter.Address[5])
		fmt.Printf("IP address: %s\n", Adapter.IpAddressList.IpAddress.String)

		fmt.Printf("IP Mask: %s\n", Adapter.IpAddressList.IpMask.String)
		fmt.Printf("Gateway: %s\n", Adapter.GatewayList.IpAddress.String)
		fmt.Printf("DHCP: %t\n", win.DHCP_ENABLE == Adapter.DhcpEnabled)
		if win.DHCP_ENABLE == Adapter.DhcpEnabled {
			fmt.Printf("DHCP Server: %s\n", Adapter.DhcpServer.IpAddress.String)
			fmt.Printf("Lease Obtained: %s\n", win.ItoaTime(Adapter.LeaseObtained, layout))
			fmt.Printf("Lease Expires: %s\n", win.ItoaTime(Adapter.LeaseExpires, layout))

		}
		fmt.Printf("%s\n", "#########################################")
		Adapter = Adapter.Next
	}
	win.Close()
}
