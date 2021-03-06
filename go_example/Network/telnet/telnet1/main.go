package main

import (
    "github.com/ziutek/telnet"
    "log"
    "os"
    // "strings"
    "time"
)

const timeout = 1 * time.Second

func checkErr(info string, err error) {
    if err != nil {
        log.Println("[ERROR] ", info, "失败.", err)
    } else {
        log.Println("[INFO] ", info, "成功.")
    }
}

func expect(t *telnet.Conn, d ...string) (err error) {
    t.SetReadDeadline(time.Now().Add(timeout))

    err = t.SkipUntil(d...)
    if err != nil {
        return err
    }

    return nil
}

func sendln(t *telnet.Conn, s string) (err error) {
    err = t.SetWriteDeadline(time.Now().Add(timeout))
    if err != nil {
        return err
    }

    buf := make([]byte, len(s)+1)
    copy(buf, s)
    buf[len(s)] = '\n'

    _, err = t.Write(buf)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    t, err := telnet.Dial("tcp", "10.10.2.209:23")
    checkErr("连接服务器", err)
    t.SetUnixWriteMode(false)

    var data []byte
    typ := "cisco"
    switch typ {
    case "cisco":
        err := expect(t, "ssword: ")
        // checkErr("密码验证", err)

        sendln(t, "cisco")
        err = expect(t, "#")
        checkErr("进入特权模式", err)

        // terminal length 0
        sendln(t, "terminal length 0")

        err = expect(t, "#")
        sendln(t, "show arp")
        checkErr("show arp", err)

        data, err = t.ReadBytes('#')

    default:
        log.Fatalln("bad host type: " + typ)
    }

    os.Stdout.Write(data)
    os.Stdout.WriteString("\n")
}
