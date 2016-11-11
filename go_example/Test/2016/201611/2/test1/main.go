package main

import (
	"fmt"
	"os"

	"github.com/gabriel-samfira/go-wmi/wmi"
)

func main() {
	w, err := wmi.NewConnection(".", `\Root\CIMV2`, nil, nil, nil, nil)
	if err != nil {
		fmt.Println("NewConnect", err)
		os.Exit(1)
	}

	// AddressFamily:
	//  2 - IPv4
	//  23 - IPv6
	// qParams := []wmi.WMIQuery{
	// 	&wmi.WMIAndQuery{wmi.QueryFields{Key: "TotalVirtualMemory", Value: 2, Type: wmi.Equals}},
	// }
	// See documentation on MSFT_NetIPAddress class at: https://msdn.microsoft.com/en-us/library/hh872425(v=vs.85).aspx
	netip, err := w.Gwmi("Win32_USBController", []string{}, nil)
	if err != nil {
		fmt.Println("Gwmi", err)
		os.Exit(1)
	}

	// TotalVirtualMemory, _ := netip.GetProperty("Name")
	// fmt.Println(TotalVirtualMemory.Value())
	elements, err := netip.Elements()
	if err != nil {
		fmt.Println("element", err)
		os.Exit(1)
	}

	if len(elements) > 0 {
		for i := 0; i < len(elements); i++ {
			address, err := elements[i].GetProperty("Description")
			if err != nil {
				fmt.Println("Description", err)
				os.Exit(1)
			}
			iface, err := elements[i].GetProperty("DeviceID")
			if err != nil {
				fmt.Println("DeviceID", err)
				os.Exit(1)
			}
			fmt.Printf(" %v -->  %v\n\n", address.Value(), iface.Value())
		}
	}

	// elements, err := netip.Elements()
	// if err != nil {
	// 	fmt.Println("element", err)
	// 	os.Exit(1)
	// }
	// if len(elements) > 0 {
	// 	for i := 0; i < len(elements); i++ {
	// 		address, err := elements[i].GetProperty("TotalVirtualMemory")
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 		}
	// 		iface, err := elements[i].GetProperty("TotalPhysicalMemory")
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			os.Exit(1)
	// 		}
	// 		fmt.Printf("Found IP %v on interface %v\n", address.Value(), iface.Value())
	// 	}
	// }

	// tmp, _ := elements[0].GetProperty("TotalVirtualMemory")

	return
}
