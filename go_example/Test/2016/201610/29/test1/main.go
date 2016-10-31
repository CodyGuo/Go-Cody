package main

import (
	"fmt"

	"github.com/gabriel-samfira/go-wmi/wmi"
)

func main() {
	// wm, err := wmi.NewConnection(".", `\root\CIMV2`)
	// if err != nil {
	// 	fmt.Println("NewConnection", err)
	// 	return
	// }
	// qParams := []wmi.WMIQuery{
	// 	&wmi.WMIAndQuery{wmi.QueryFields{Key: "MACAddress", Value: "78-0C-B8-EF-0F-CE", Type: wmi.Equals}},
	// }
	// sw, _ := wm.GetOne("Win32_NetworkAdapterConfiguration", []string{}, qParams)
	// t, _ := sw.Get("IPAddress")

	// test, _ := t.GetText(0)
	// fmt.Println(test)

	type s struct {
		Name string
	}

	var dst []s

	wm, _ := wmi.NewConnection(".")
	sw, _ := wm.Get("Win32_Process")
	ra := sw.Raw()
	result := ra.ToIDispatch()
	defer ra.Clear()
	enupty, _ := result.GetProperty("_NewEnum")
	enum, _ := enupty.ToIUnknown()
	defer enum.Release()
	n, _ := sw.Count()
	sw.Set("Win32_Process", 0, n)

	count
}
