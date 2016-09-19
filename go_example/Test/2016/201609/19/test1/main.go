package main

import (
	"errors"
	"log"
)

import (
	"github.com/CodyGuo/win"
)

var (
	winExecError = map[uint32]string{
		0:  "The system is out of memory or resources.",
		2:  "The .exe file is invalid.",
		3:  "The specified file was not found.",
		11: "The specified path was not found.",
	}
)

func main() {
	err := execRun("cmd /c start http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
}

func execRun(cmd string) error {
	lpCmdLine := win.StringToBytePtr(cmd)
	// http://baike.baidu.com/link?url=51sQomXsIt6OlYEAV74YZ0JkHDd2GbmzXcKj_4H1R4ILXvQNf3MXIscKnAkSR93e7Fyns4iTmSatDycEbHrXzq
	ret := win.WinExec(lpCmdLine, win.SW_HIDE)
	if ret <= 31 {
		return errors.New(winExecError[ret])
	}

	return nil
}
