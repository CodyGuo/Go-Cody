/*
#include <utmp.h>

struct utmp *getutent(){
    return *utmp
}

*/
package main

import (
	"C"
)

import (
	"fmt"
)

const (
	EMPTY = iota
	RUN_LVL
	BOOT_TIME
	NEW_TIME
	OLD_TIME
	INIT_PROCESS
	LOGIN_PROCESS
	USER_PROCESS
	DEAD_PROCESS
	ACCOUNTING
)

const (
	UT_LINESIZE = 32
	UT_NAMESIZE = 32
	UT_HOSTSIZE = 256
)

type exit_status struct {
	termiation int32
	exit       int32
}
type timeval struct {
	tv_sec  int64
	tv_usec int64
}

type Utmp struct {
	ut_type int32
	ut_pid  int
	ut_id   [4]byte
	ut_user [UT_NAMESIZE]byte
	ut_host [UT_HOSTSIZE]byte
	ut_exit exit_status
	// ut_line [UT_LINESIZE]byte
	ut_session int
	ut_tv      timeval
	addr_v6    [4]int32
	_unused    [20]byte
}

func main() {
	var p_utent *C.utmp
	C.getutent()

	fmt.Println(p_utent.ut_user)
}
