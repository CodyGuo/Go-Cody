package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var (
	libuser32 *syscall.DLL
	isIconic  *syscall.Proc
)

func init() {
	// Library
	libuser32 = syscall.MustLoadDLL("user32.dll")
	// functions
	isIconic = libuser32.MustFindProc("IsIconic")
}

/* 确定指定的窗口是否被最小化(图标)
https://msdn.microsoft.com/en-us/library/windows/desktop/ms633527(v=vs.85).aspx
*/
func IsIconic(hWnd win.HWND) bool {
	ret, _, _ := isIconic.Call(uintptr(hWnd))
	return ret == win.TRUE
}

type MyWindow struct {
	*walk.MainWindow
}

func main() {
	mw := new(MyWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "最小化测试",
		MinSize:  Size{280, 200},
		Layout:   HBox{},
		Children: []Widget{
			PushButton{
				Text: "测试",
				OnClicked: func() {
					walk.MsgBox(mw, "测试", "最小化事件捕捉!", walk.MsgBoxIconInformation)
				},
			},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}
	mw.SizeChanged().Attach(func() {
		if mw.X() == -32000 && mw.Y() == -32000 {
			fmt.Printf("X, Y == -32000 --> 窗口最小化, X = %d, Y = %d\n", mw.X(), mw.Y())
		}
		if IsIconic(mw.Handle()) {
			message := fmt.Sprintf("IsIconic --> 窗口最小化, X = %d, Y = %d", mw.X(), mw.Y())
			walk.MsgBox(mw, "事件", message, walk.MsgBoxIconWarning)
		}
	})
	mw.Run()
}
