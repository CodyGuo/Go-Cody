package main

import (
    // "fmt"
    "github.com/codyguo/gosftp"
    "golang.org/x/crypto/ssh"
    "io"
    "log"
    // "os"
    "time"
    // "bytes"
    "strings"
)

var (
    timeout time.Duration
)

func main() {
    config := &ssh.ClientConfig{
        User: "root",
        Auth: []ssh.AuthMethod{
            ssh.Password("123456"),
        },
    }
    client, err := ssh.Dial("tcp", "10.10.3.100:22", config)
    if err != nil {
        panic(err)
    }
    // Each ClientConn can support multiple interactive sessions,
    // represented by a Session.
    defer client.Close()
    // Create a session
    session, err := client.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }
    defer session.Close()

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
        log.Fatal(err)
    }
    w, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }

    r, err := session.StdoutPipe()
    if err != nil {
        panic(err)
    }

    err = session.Shell()
    if err != nil {
        log.Println("创建seesion失败: ", err)
    }

    timeout = 10

    sendIn(w, "admin")
    result := except("NAC>", r)
    log.Println(result)

    sendIn(w, "admin2014")
    result = except("Password:", r)
    log.Println(result)

    sendIn(w, "echo cody.guo")
    result = except("#", r)

    log.Println(result)
    fp, err := gosftp.NewClient(session)
    if err != nil {
        log.Println("创建sftp失败.")
    }
    fp.Mkdir("/nac/codyguo", 0777)
    defer fp.Close()
    sendIn(w, "exit")

    sendIn(w, "exit")

    session.Wait()

    log.Fatalf("本次SSH结束 [%s:%d].", "10.10.3.100", 22)
}

func except(sep string, r io.Reader, timeoutSec ...time.Duration) (result string) {
    var (
        buf [65 * 1024]byte
        t   int
    )

    if len(timeoutSec) == 1 {
        if timeoutSec[0] != time.Nanosecond {
            timeout = timeoutSec[0]
        }
    }

    timer := time.NewTimer(time.Second * timeout)

    for {
        n, err := r.Read(buf[t:])
        // log.Println("正在读取", n)
        if err != nil {
            // log.Println("退出读取", n)
            log.Fatal(err)
        }

        t += n
        tmpLogIn := string(buf[:t])
        if strings.Contains(tmpLogIn, sep) {
            result = string(buf[:t])
            timer.Stop()
            break
        }

        // 超时器
        go func() {
            <-timer.C
            log.Fatalf("SSH超时: %d 秒,返回值中不存在: %s.\n", timeout, sep)
        }()
    }

    return
}

func sendIn(w io.Writer, cmd string) {
    w.Write([]byte(cmd + "\n"))
}
