package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

import (
	"github.com/ziutek/telnet"
)

type systemFlag uint

const (
	windows systemFlag = iota
	linux
	macOS
	switchOS
)

type Telneter interface {
	Connect() error
	ReadUnitl(d ...string) error
	WriteCmd(s string) error
}

type TelnetConfig struct {
	*telnet.Conn

	IP     string
	User   string
	Passwd string

	Debug   bool
	System  systemFlag
	Timeout time.Duration
}

type SwitchConfig struct {
	TelnetConfig
	enablePasswd string
}

func (t *TelnetConfig) Connect() (err error) {
	t.Conn, err = telnet.DialTimeout("tcp", t.IP+":23", t.Timeout)
	return err
}

func (t *TelnetConfig) ReadUnitl(d ...string) (err error) {
	var data []byte
	if data, err = t.ReadUntil(d...); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if t.Debug {
		os.Stdout.WriteString("\n")
		os.Stdout.WriteString("[DEBUG] ")
		os.Stdout.Write(data)
		os.Stdout.WriteString("\n")
	}
	return nil
}

func (t *TelnetConfig) WriteCmd(s string) (err error) {
	buf := make([]byte, len(s))
	copy(buf, s)
	if _, err := t.Write(buf); err != nil {
		return err
	}

	switch t.System {
	case windows:
		t.Write([]byte{telnet.CR, telnet.LF})
	case linux, switchOS:
		t.Write([]byte{telnet.CR})
	case macOS:
		t.Write([]byte{telnet.LF, telnet.CR})
	}
	return nil
}

func main() {
	ok1 := make(chan bool, 1)
	ok2 := make(chan bool, 1)
	go window2003(ok1)
	go switchCisco(ok2)

	<-ok1
	<-ok2
}

func window2003(ok chan bool) {
	var tWin Telneter
	twinConf := &TelnetConfig{
		IP:      "10.10.2.116",
		User:    "administrator",
		Passwd:  "hpiMan!#%",
		Debug:   true,
		System:  windows,
		Timeout: 3 * time.Second,
	}
	tWin = twinConf

	err := tWin.Connect()
	checkErr("连接服务器"+twinConf.IP, err)

	err = tWin.ReadUnitl("login:")
	checkErr("用户名"+twinConf.User, err)
	tWin.WriteCmd(twinConf.User)

	err = tWin.ReadUnitl("password:")
	checkErr("密码验证", err)
	tWin.WriteCmd(twinConf.Passwd)

	err = tWin.ReadUnitl("Administrator>")
	checkErr("登录服务器"+twinConf.IP, err)

	tWin.WriteCmd("type cody.log")
	err = tWin.ReadUnitl("Administrator>")
	checkErr("查看日志内容", err)

	tWin.WriteCmd("mkdir cody")
	err = tWin.ReadUnitl("Administrator>")
	checkErr("创建文件夹", err)

	tWin.WriteCmd("exit")

	ok <- true
	close(ok)
}

func switchCisco(ok chan bool) {
	var sw Telneter
	swConf := &SwitchConfig{
		TelnetConfig: TelnetConfig{IP: "10.10.3.252",
			Passwd:  "cisco",
			Debug:   true,
			System:  switchOS,
			Timeout: 3 * time.Second},
		enablePasswd: "cisco",
	}
	sw = swConf

	sw.Connect()
	sw.ReadUnitl("assword:")
	sw.WriteCmd(swConf.Passwd)
	sw.ReadUnitl(">")
	sw.WriteCmd("enable")
	sw.ReadUnitl("assword:")
	sw.WriteCmd(swConf.enablePasswd)
	sw.ReadUnitl("#")
	sw.WriteCmd("show arp")
	sw.ReadUnitl("#")
	ok <- true
	close(ok)
}

func checkErr(info string, err error) {
	if err != nil {
		log.Println("[ERROR] ", info, "失败.", err)
	} else {
		log.Println("[INFO] ", info, "成功.")
	}
}
