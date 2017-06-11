package main

import (
	"syscall"
	"unsafe"
)

const (
	// 获取屏幕保护开关
	SPI_GETSCREENSAVEACTIVE = 0x0010
	// 设置屏保开关
	SPI_SETSCREENSAVEACTIVE = 0x0011
	// 设置屏保等待时间
	SPI_SETSCREENSAVETIMEOUT = 0x000F
	// 设备屏保在恢复时显示屏幕
	SPI_SETSCREENSAVESECURE = 0x0077

	// 设置桌面背景
	SPI_SETDESKWALLPAPER = 0x0014

	// Writes the new system-wide parameter setting to the user profile.
	SPIF_UPDATEINIFILE = 1

	// Broadcasts the WM_SETTINGCHANGE message after updating the user profile.
	SPIF_SENDWININICHANGE = 2

	FALSE = 0
	TRUE  = 1
)

var (
	// Library
	libuser32   uintptr
	libkernel32 uintptr

	// Functions
	systemParametersInfo uintptr
	getVersion           uintptr
)

func init() {
	// Library
	libuser32 = MustLoadLibrary("user32.dll")
	libkernel32 = MustLoadLibrary("kernel32.dll")
	// Functions
	systemParametersInfo = MustGetProcAddress(libuser32, "SystemParametersInfoW")
	getVersion = MustGetProcAddress(libkernel32, "GetVersion")
}

func MustLoadLibrary(name string) uintptr {
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}

	return uintptr(lib)
}

func MustGetProcAddress(lib uintptr, name string) uintptr {
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}

	return uintptr(addr)
}

/* 通过调用Win32 API函数SystemParametersInfo 设置桌面壁纸
之前我们已经设置了壁纸的类型，但是壁纸图片的实际文件路径还没设置。SystemParametersInfo 这个函数位于user32.dll中，它支持我们从系统中获得硬件和配置信息。它有四个参数，第一个指明调用这个函数所要执行的操作，接下来的两个参数指明将要设置的数据，依据第一个参数的设定。最后一个允许指定改变是否被保存。这里第一个参数我们应指定SPI_SETDESKWALLPAPER，指明我们是要设置壁纸。接下来是文件路径。在Vista之前必须是一个.bmp的文件。Vista和更高级的系统支持.jpg格式。
SPI_SETDESKWALLPAPER参数使得壁纸改变保存并持续可见。
*/
func SystemParametersInfo(uiAction, uiParam uint32, pvParam unsafe.Pointer, fWinIni uint32) bool {
	ret, _, _ := syscall.Syscall6(systemParametersInfo, 4,
		uintptr(uiAction),
		uintptr(uiParam),
		uintptr(pvParam),
		uintptr(fWinIni),
		0,
		0)

	return ret != 0
}

func GetVersion() int64 {
	ret, _, _ := syscall.Syscall(getVersion, 0,
		0,
		0,
		0)
	return int64(ret)
}

func stringToPointer(str string) unsafe.Pointer {
	return unsafe.Pointer(syscall.StringToUTF16Ptr(str))
}
