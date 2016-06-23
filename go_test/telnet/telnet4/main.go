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

const (
	INFO  = "[INFO ] "
	DEBUG = "[DEBUG] "
	ERROR = "[ERROR] "
)

type TelnetConfig struct {
	*telnet.Conn

	IP     string
	User   string
	Passwd string

	Debug   bool
	System  systemFlag
	Timeout time.Duration
}

func NewTelnetConfig() *TelnetConfig {
	conf := new(TelnetConfig)

	return conf
}

func (t *TelnetConfig) Connect() (err error) {
	t.Conn, err = telnet.DialTimeout("tcp", t.IP+":23", t.Timeout)
	return err
}

func (t *TelnetConfig) readAll() string {
	var line []byte
	for {
		b, err := t.ReadByte()
		if err != nil {
			return ""
		}
		line = append(line, b)
	}

	return string(line)
}

func (t *TelnetConfig) ReadUnitl(d ...string) (err error) {
	go func() {
		time.Sleep(3 * time.Second)
		// os.Exit(1)
		// err = fmt.Errorf("%s读取返回超时.\n", ERROR)
		return
	}()

	var data []byte
	if data, err = t.ReadUntil(d...); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if t.Debug {
		os.Stdout.WriteString("\n")
		os.Stdout.WriteString(DEBUG)
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
	conf := NewTelnetConfig()
	twinConf := &TelnetConfig{
		IP:      "10.10.2.76",
		User:    "administrator",
		Passwd:  "hupuruanjian",
		Debug:   true,
		System:  windows,
		Timeout: 3 * time.Second,
	}

	conf = twinConf

	err := conf.Connect()
	checkErr("连接服务器"+twinConf.IP, err)

	err = conf.ReadUnitl("login:")
	checkErr("用户名"+twinConf.User, err)
	conf.WriteCmd(twinConf.User)

	err = conf.ReadUnitl("password:")
	checkErr("密码验证", err)
	conf.WriteCmd(twinConf.Passwd)

	err = conf.ReadUnitl("Administrator>")
	checkErr("登录服务器"+twinConf.IP, err)

	err = conf.ReadUnitl("Helper to upload success", "Administrator>")
	checkErr("执行打包", err)

	conf.WriteCmd("exit")

}
func checkErr(info string, err error) {
	if err != nil {
		log.Println("[ERROR] ", info, "失败.", err)
	} else {
		log.Println("[INFO] ", info, "成功.")
	}
}
