package main

import (
	"fmt"
	"unsafe"
)

import (
	"github.com/CodyGuo/win"
)

func main() {
	v := win.GetVersion()
	major := byte(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	fmt.Println(major, minor, build)

	var osvi *win.OSVERSIONINFOEX
	osvi = new(win.OSVERSIONINFOEX)
	osvi.DwOSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	fmt.Println(osvi.DwOSVersionInfoSize)
	ok := win.GetVersionExA(osvi)
	err := win.GetLastError()
	fmt.Println(osvi.DwMajorVersion, osvi.DwMinorVersion, ok, err)

}
