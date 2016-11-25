package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	TRUE  = 1
	FALSE = 0
	NULL  = 0
)

var (
	dhnetsdkDll *syscall.DLL
)

var (
	client_Init            *syscall.Proc
	client_SetNetworkParam *syscall.Proc
	client_LoginEx2        *syscall.Proc
	client_Logout          *syscall.Proc
	client_Cleanup         *syscall.Proc
	client_SetDevConfig    *syscall.Proc
	client_GetDevConfig    *syscall.Proc
	client_GetLastError    *syscall.Proc
	client_QuerySystemInfo *syscall.Proc
)

type DH_DEV_ENABLE_INFO struct {
	IsFucEnable [512]uint32
}

func init() {
	dhnetsdkDll = syscall.MustLoadDLL("dhnetsdk.dll")

	client_Init = dhnetsdkDll.MustFindProc("CLIENT_Init")
	client_SetNetworkParam = dhnetsdkDll.MustFindProc("CLIENT_SetNetworkParam")
	client_LoginEx2 = dhnetsdkDll.MustFindProc("CLIENT_LoginEx2")
	client_Logout = dhnetsdkDll.MustFindProc("CLIENT_Logout")
	client_Cleanup = dhnetsdkDll.MustFindProc("CLIENT_Cleanup")
	client_SetDevConfig = dhnetsdkDll.MustFindProc("CLIENT_SetDevConfig")
	client_GetDevConfig = dhnetsdkDll.MustFindProc("CLIENT_GetDevConfig")
	client_GetLastError = dhnetsdkDll.MustFindProc("CLIENT_GetLastError")
	client_QuerySystemInfo = dhnetsdkDll.MustFindProc("CLIENT_QuerySystemInfo")
}

func main() {
	ok := CLIENT_Init(0, 0)
	if !ok {
		fmt.Println("初始化dll失败.")
		return
	}

	pNetParam := new(NET_PARAM)
	pNetParam.nConnectTryNum = 2
	pNetParam.nGetDevInfoTime = 3000
	CLIENT_SetNetworkParam(pNetParam)
	ip := "10.10.2.15"
	pchDVRIP := StringToBytePtr(ip)
	pchUserName := StringToBytePtr("admin")
	pchPassword := StringToBytePtr("admin2016")
	var err int32
	var deviceInfo NET_DEVICEINFO_Ex
	lLoginID := CLIENT_LoginEx2(pchDVRIP, 37777, pchUserName, pchPassword, EM_LOGIN_SPEC_CAP_TCP, 0, &deviceInfo, &err)
	if lLoginID == 0 {
		fmt.Printf("登录摄像头[%s]失败.%v\n", ip, err)
		return
	}
	fmt.Printf("登录摄像头[%s]成功,登录ID: %d.\n", ip, lLoginID)

	var pSysInfoBuffer DH_DEV_ENABLE_INFO
	var nSysInfolen int32
	ok = CLIENT_QuerySystemInfo(lLoginID, ABILITY_DEVALL_INFO, (*byte)(unsafe.Pointer(&pSysInfoBuffer)), int32(unsafe.Sizeof(pSysInfoBuffer)), &nSysInfolen, 1000)
	if !ok {
		fmt.Printf("获取摄像头[%s]支持的功能列表失败.", ip)
		return
	}
	fmt.Printf("获取摄像头[%s]支持的功能列表成功.\n", ip)

	fmt.Printf("摄像头[%s]支持的功能如下:\n", ip)
	for index, name := range names {
		if pSysInfoBuffer.IsFucEnable[index] != FALSE {
			fmt.Printf("-- %s\n", name)
		}
	}
}

func StringToBytePtr(str string) *byte {
	return syscall.StringBytePtr(str)
}

func CLIENT_Init(cbDisConnect uintptr, dwUser uint32) bool {
	ret, _, _ := client_Init.Call(cbDisConnect, uintptr(dwUser))

	return ret == TRUE
}

func CLIENT_SetNetworkParam(pNetParam *NET_PARAM) {
	client_SetNetworkParam.Call(uintptr(unsafe.Pointer(pNetParam)))
}

func CLIENT_LoginEx2(pchDVRIP *byte, wDVRPort uint16, pchUserName, pchPassword *byte,
	emSpecCap EM_LOGIN_SPAC_CAP_TYPE, pCapParam uintptr, lpDeviceInfo *NET_DEVICEINFO_Ex, err *int32) int64 {
	ret, _, _ := client_LoginEx2.Call(uintptr(unsafe.Pointer(pchDVRIP)),
		uintptr(wDVRPort),
		uintptr(unsafe.Pointer(pchUserName)),
		uintptr(unsafe.Pointer(pchPassword)),
		uintptr(emSpecCap),
		pCapParam,
		uintptr(unsafe.Pointer(lpDeviceInfo)),
		uintptr(unsafe.Pointer(err)))

	return int64(ret)
}

func CLIENT_SetDevConfig(lLoginID int64, dwCommand uint32, lChannel int32, lpInBuffer uintptr, dwInBufferSize uint32, waittime int32) bool {
	ret, _, _ := client_SetDevConfig.Call(uintptr(lLoginID),
		uintptr(dwCommand),
		uintptr(lChannel),
		lpInBuffer,
		uintptr(dwInBufferSize),
		uintptr(waittime))

	return ret == TRUE
}

func CLIENT_GetDevConfig(lLoginID int64, dwCommand uint32, lChannel int32, lpOutBuffer uintptr, dwOutBufferSize uint32, lpBytesReturned *uint32, waittime int32) bool {
	ret, _, _ := client_SetDevConfig.Call(uintptr(lLoginID),
		uintptr(dwCommand),
		uintptr(lChannel),
		lpOutBuffer,
		uintptr(dwOutBufferSize),
		uintptr(unsafe.Pointer(lpBytesReturned)),
		uintptr(waittime))

	return ret == TRUE
}

func CLIENT_QuerySystemInfo(lLoginID int64, nSystemType int32, pSysInfoBuffer *byte, maxlen int32, nSysInfolen *int32, waittime int) bool {
	ret, _, _ := client_QuerySystemInfo.Call(uintptr(lLoginID),
		uintptr(nSystemType),
		uintptr(unsafe.Pointer(pSysInfoBuffer)),
		uintptr(maxlen),
		uintptr(unsafe.Pointer(nSysInfolen)),
		uintptr(waittime))

	return ret == TRUE
}

func CLIENT_GetLastError() uint32 {
	ret, _, _ := client_GetLastError.Call()

	return uint32(ret)
}

// var netSnmpCFG DHDEV_NET_SNMP_CFG
// netSnmpCFG.bEnable = '1'
// netSnmpCFG.bSNMPV1 = '1'
// netSnmpCFG.bSNMPV2 = '1'
// netSnmpCFG.iSNMPPort = 161
// netSnmpCFG.iTrapPort = 162
// var public [DH_MAX_SNMP_COMMON_LEN]byte
// copy(public[:], "public")

// var privite [DH_MAX_SNMP_COMMON_LEN]byte
// copy(public[:], "privite")
// netSnmpCFG.szReadCommon = public
// netSnmpCFG.szWriteCommon = privite

// ok = CLIENT_SetDevConfig(lLoginID, DH_DEV_SNMP_CFG, -1, uintptr(unsafe.Pointer(&netSnmpCFG)), uint32(unsafe.Sizeof(netSnmpCFG)), 3000)

// if ok {
//  fmt.Printf("设置 %s SNMP成功: %v\n", ip, ok)
// } else {
//  fmt.Printf("设置 %s SNMP失败: %v,ID: %v\n", ip, ok, lLoginID)
// }

// var curDateTime NET_TIME = NET_TIME{2016, 12, 1, 10, 00, 00}
// ok = CLIENT_SetDevConfig(lLoginID, DH_DEV_TIMECFG, -1, uintptr(unsafe.Pointer(&curDateTime)), uint32(unsafe.Sizeof(curDateTime)), 1000)
// if ok {
//  fmt.Println("设置成功")
// } else {
//  fmt.Println("设置失败.")
// }

// var dwRet uint32 = 0

// ok = CLIENT_GetDevConfig(lLoginID, DH_DEV_TIMECFG, -1, uintptr(unsafe.Pointer(&curDateTime)), uint32(unsafe.Sizeof(curDateTime)), &dwRet, 3000)
// if ok {
//  fmt.Println("get 成功", ok)
// }
// fmt.Println(curDateTime.dwYear, dwRet)
