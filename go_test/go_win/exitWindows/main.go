package main

import (
	"flag"
	"fmt"
)

import (
	. "github.com/CodyGuo/win"
)

var (
	arg string
)

func init() {
	flag.StringVar(&arg, "uFlags", "", "shutdown logoff reboot")
}

func main() {
	flag.Parse()

	switch arg {
	case "logoff":
		logoff()
	case "reboot":
		reboot()
	case "shutdown":
		shutdown()
	default:
		fmt.Println("您输入的参数有误.")
	}
}

func logoff() {
	ExitWindowsEx(EWX_LOGOFF, 0)
}

func reboot() {
	getPrivileges()
	ExitWindowsEx(EWX_REBOOT, 0)
}

func shutdown() {
	getPrivileges()
	ExitWindowsEx(EWX_SHUTDOWN, 0)
}

func getPrivileges() {
	var hToken HANDLE
	var tkp TOKEN_PRIVILEGES

	OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY, &hToken)
	LookupPrivilegeValueA(nil, StringToBytePtr(SE_SHUTDOWN_NAME), &tkp.Privileges[0].Luid)
	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED
	AdjustTokenPrivileges(hToken, false, &tkp, 0, nil, nil)
}
