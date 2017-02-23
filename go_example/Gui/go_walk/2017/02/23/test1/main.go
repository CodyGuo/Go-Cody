package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

const (
	TRUE  = 1
	FALSE = 0
	NULL  = 0
)

var next = 0

func callback(fn interface{}) uintptr {
	return syscall.NewCallback(fn)
}

func windowText(hWnd win.HWND) string {
	textLength := win.SendMessage(hWnd, win.WM_GETTEXTLENGTH, 0, 0)
	buf := make([]uint16, textLength+1)
	win.SendMessage(hWnd, win.WM_GETTEXT, uintptr(textLength+1), uintptr(unsafe.Pointer(&buf[0])))
	return syscall.UTF16ToString(buf)
}

func EnumMainTVWindow_cn(hWnd win.HWND, lParam uintptr) uintptr {
	ret := windowText(hWnd)
	if ret != "" {
		switch next {
		case 1:
			fmt.Printf("TeamViewer ID  : %s\n", ret)
			next += 1
		case 2:
			fmt.Printf("TeamViewer PASS: %s\n", ret)
			next += 1
		}
		if ret == "控制远程计算机" {
			next += 1
		}
		// fmt.Printf("EnumMainTVWindow_cn --> %v\n", ret)
	}

	return TRUE
}

func FindWindow(windowName string) win.HWND {
	lpWindowName := syscall.StringToUTF16Ptr(windowName)

	return win.FindWindow(nil, lpWindowName)
}

func EnumChildWindows(hWndParent win.HWND, lpEnumFunc, lParam uintptr) bool {
	ret := win.EnumChildWindows(hWndParent, lpEnumFunc, lParam)
	return ret
}

func main() {
	process := os.Args[1]
	hwndTeamViewer_cn := FindWindow(process)
	if hwndTeamViewer_cn != NULL {
		ok := EnumChildWindows(hwndTeamViewer_cn, callback(EnumMainTVWindow_cn), 0)

		fmt.Printf("enumchildwindows --> %v\n", ok)
	} else {
		fmt.Printf("No FindWindow --> %v,ErrCode --> %v\n", process, hwndTeamViewer_cn)
	}
}
