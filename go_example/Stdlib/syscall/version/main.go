package main

import (
	"os"
	"syscall"
)

func CheckError(err error) {
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
}

func printVersion(v uint32) {
	major := byte(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	print("windows version ", major, ".", minor, ".", build)
}

func main() {
	k32dll := syscall.MustLoadDLL("kernel32.dll")
	defer syscall.FreeLibrary(k32dll.Handle)

	getVersion, err := k32dll.FindProc("GetVersion")
	CheckError(err)

	ret, _, _ := getVersion.Call()
	printVersion(uint32(ret))

}
