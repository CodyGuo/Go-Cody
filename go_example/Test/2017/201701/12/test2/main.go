package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyWindow struct {
	walk.WidgetBase
}

func NewMyWindow(parent walk.Container) (*MyWindow, error) {
	m := &MyWindow{}
	if err := walk.InitWidget(
		m,
		parent,
		"MY",
		win.WS_OVERLAPPEDWINDOW,
		win.WS_EX_CONTROLPARENT); err != nil {
		fmt.Println("init", err)
		return nil, err
	}
	fmt.Println("测试新事件", m)
	return m, nil
}

func (mw *MyWindow) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	fmt.Printf("msg: %v\n", msg)
	switch msg {
	case win.WM_COMMAND:
		fmt.Printf("WM_COMMAND msg=%v\n", msg)
		if lParam != 0 { //Reflect message to control
			h := win.HWND(lParam)
			fmt.Printf("WM_COMMAND h=%v\n", h)
		}
		fmt.Println("WM_COMMAND DONE")
		return 0
	case win.WM_CLOSE:
		win.DestroyWindow(hwnd)
		return 0
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
		return 1
	}

	return mw.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

func main() {
	mw, _ := walk.NewMainWindow()
	fmt.Printf("mw %v create.\n", mw)
	m, _ := NewMyWindow(mw)
	fmt.Printf("newmywind %v\n", m)
	win.IDC_CROSS
	mw.Run()
}

func register() {
	var wc win.WNDCLASSEX
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	wc.LpfnWndProc = wndProcPtr
	wc.HInstance = hInst
	wc.HIcon = hIcon
	wc.HCursor = hCursor
	wc.HbrBackground = win.COLOR_BTNFACE + 1
	wc.LpszClassName = syscall.StringToUTF16Ptr(className)

	if atom := win.RegisterClassEx(&wc); atom == 0 {
		panic("RegisterClassEx")
	}
}
