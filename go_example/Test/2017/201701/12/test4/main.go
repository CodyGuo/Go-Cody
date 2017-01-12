package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/codyguo/walk"
	"github.com/codyguo/win"
)

const mainWindowWindowClass = `\o/ USB_Class \o/`

func init() {
	walk.MustRegisterWindowClass(mainWindowWindowClass)
}

type USB struct {
	walk.MainWindow
}

func NewUSB() (*USB, error) {
	usb := new(USB)
	if err := walk.InitWindow(
		usb,
		nil,
		mainWindowWindowClass,
		win.WS_OVERLAPPEDWINDOW,
		win.WS_EX_CONTROLPARENT); err != nil {
		return nil, err
	}

	return usb, nil
}

func (usb *USB) RegisterDeviceNotification() bool {
	var notificationFilter = new(win.DEV_BROADCAST_DEVICEINTERFACE)
	notificationFilter.Dbcc_size = uint32(unsafe.Sizeof(*notificationFilter))
	notificationFilter.Dbcc_devicetype = win.DBT_DEVTYP_DEVICEINTERFACE
	notificationFilter.Dbcc_classguid = win.GUID_DEVINTERFACE_USB_DEVICE

	handle := win.HANDLE(usb.Handle())
	ret := win.RegisterDeviceNotificationW(handle, (uintptr)(unsafe.Pointer(notificationFilter)), win.DEVICE_NOTIFY_WINDOW_HANDLE)

	return ret != 0
}

func (usb *USB) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_DEVICECHANGE:
		switch wParam {
		case win.DBT_DEVICEARRIVAL:
			fmt.Printf("USB插入 -> %v\n", msg)
		case win.DBT_DEVICEREMOVECOMPLETE:
			fmt.Printf("USB拔出 -> %v\n", msg)
		}
	}
	return usb.FormBase.WndProc(hwnd, msg, wParam, lParam)
}

func main() {
	usb, err := NewUSB()
	if err != nil {
		log.Fatal(err)
	}

	usb.RegisterDeviceNotification()
	usb.Run()
}
