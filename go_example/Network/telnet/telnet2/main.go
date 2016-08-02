package main

import (
    "log"
    "os"
    "time"
)

import (
    "github.com/ziutek/telnet"
)

var (
    tnConfig *TelnetConfig
)

type TelnetConfig struct {
    Ip           string
    User         string
    Passwd       string
    EnablePasswd string

    Debug   bool
    Timeout time.Duration
}

func NewTelnet() (tn *telnet.Conn, err error) {
    return telnet.DialTimeout("tcp", tnConfig.Ip+":23", tnConfig.Timeout)
}

func (t *TelnetConfig) readUnitl(tn *telnet.Conn, d ...string) (err error) {
    var data []byte
    data, err = tn.ReadUntil(d...)
    if err != nil {
        return err
    }

    if t.Debug {
        os.Stdout.WriteString("\n")
        os.Stdout.Write(data)
        os.Stdout.WriteString("\n")
    }

    return nil
}

func (t *TelnetConfig) writeCmd(tn *telnet.Conn, s string) (err error) {
    buf := make([]byte, len(s)+1)
    copy(buf, s)
    buf[len(s)] = '\n'

    _, err = tn.Write(buf)
    if err != nil {
        return err
    }

    return nil
}

func init() {
    tnConfig = new(TelnetConfig)
    tnConfig.Ip = "10.10.2.200"
    tnConfig.User = "h3c"
    tnConfig.Passwd = "h3c"
    tnConfig.EnablePasswd = "cisco"
    tnConfig.Timeout = 3 * time.Second
    tnConfig.Debug = true
}

func main() {
    tn, err := NewTelnet()
    defer tn.Close()
    checkErr("连接服务器", err)
    tn.SetUnixWriteMode(false)

    err = tnConfig.readUnitl(tn, "Username:")
    checkErr("用户名验证", err)
    tnConfig.writeCmd(tn, tnConfig.User)

    err = tnConfig.readUnitl(tn, "Password:")
    checkErr("密码验证", err)
    tnConfig.writeCmd(tn, tnConfig.Passwd)

    err = tnConfig.readUnitl(tn, "<H3C>")
    checkErr("super模式验证", err)
    tnConfig.writeCmd(tn, "super")

    err = tnConfig.readUnitl(tn, "Password:")
    checkErr("super密码验证", err)
    tnConfig.writeCmd(tn, tnConfig.EnablePasswd)

    err = tnConfig.readUnitl(tn, ">")
    checkErr("进入super模式", err)
    tnConfig.writeCmd(tn, "system")

    err = tnConfig.readUnitl(tn, "]")
    checkErr("进入全局配置模式", err)

    tnConfig.writeCmd(tn, "dis arp")
    err = tnConfig.readUnitl(tn, "]")
    checkErr("dis arp", err)

}

func checkErr(info string, err error) {
    if err != nil {
        log.Println("[ERROR] ", info, "失败.", err)
    } else {
        log.Println("[INFO] ", info, "成功.")
    }
}
