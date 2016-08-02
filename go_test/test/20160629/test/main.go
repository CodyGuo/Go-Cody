package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

import (
	"github.com/CodyGuo/win"
)

const (
	address = "baidu.com:80"
)

var (
	shutTime int
)

func main() {
	shutTime = 3
	for {
		time.Sleep(1 * time.Second)
		if err := checkNetwork(); err != nil {
			fmt.Println("您已合规.")
		} else {
			fmt.Printf("您违规了,将在%2d秒后关机,请注意.\n", shutTime)
			shutTime--
			if 0 == shutTime {
				shutdown(shutTime)
				os.Exit(0)
			}
		}
	}
}

func checkNetwork() error {
	_, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return err
	}
	return nil

}

func shutdown(shutTime int) {
	var hToken win.HANDLE
	var tkp win.TOKEN_PRIVILEGES

	win.OpenProcessToken(win.GetCurrentProcess(), win.TOKEN_ADJUST_PRIVILEGES|win.TOKEN_QUERY, &hToken)
	win.LookupPrivilegeValueA(nil, win.StringToBytePtr(win.SE_SHUTDOWN_NAME), &tkp.Privileges[0].Luid)
	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = win.SE_PRIVILEGE_ENABLED

	win.AdjustTokenPrivileges(hToken, false, &tkp, 0, nil, nil)
	win.ExitWindowsEx(win.EWX_SHUTDOWN|win.EWX_FORCE, 0)
}
