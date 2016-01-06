package main

import (
    "bufio"
    "flag"
    "fmt"
    "golang.org/x/crypto/ssh"
    "golang.org/x/crypto/ssh/terminal"
    "os"
    "strings"
)

func main() {
    var path, ip string
    flag.StringVar(&path, "p", "cfg", "指定passwd文件 -p=cfg")
    flag.StringVar(&ip, "i", "127.0.0.1", "指定要登录的IP -i=127.0.0.1")
    flag.Parse()
    user := parseconfig(path, ip)
    if user == nil || len(user) != 4 {
        fmt.Println("匹配IP出错：", ip)
        return
    }
    client(user[0], user[1], fmt.Sprintf("%s:%s", user[2], user[3]))
}

func parseconfig(path, ip string) []string {
    File, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    defer File.Close()
    str := "密码文件格式：root 123456 127.0.0.1 22"
    IO := bufio.NewReader(File)
    line, _, err := IO.ReadLine()
    if err != nil {
        if err.Error() == "EOF" {
            fmt.Println(str)
            return nil
        }
        fmt.Println("解析配置文件: ", err)
        fmt.Println(str)
        return nil
    }
    if strings.Contains(string(line), ip) {
        return split(string(line))
    }
    fmt.Println(str)
    return nil
}
func split(str string) []string {
    var l []string
    list := strings.Split(str, " ")
    for _, v := range list {
        if len(v) == 0 {
            continue
        }
        if strings.Contains(v, "    ") {

            list := strings.Split(v, "  ")
            for _, v := range list {
                if len(v) == 0 {
                    continue
                }
                l = append(l, v)
            }
            continue
        }
        l = append(l, v)
    }
    return l
}
func client(user, passwd, ip string) {
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(passwd),
        },
    }
    client, err := ssh.Dial("tcp", ip, config)
    if err != nil {
        fmt.Println("建立连接: ", err)
        return
    }
    defer client.Close()
    session, err := client.NewSession()
    if err != nil {
        fmt.Println("创建Session出错: ", err)
        return
    }
    defer session.Close()

    fd := int(os.Stdin.Fd())
    oldState, err := terminal.MakeRaw(fd)
    if err != nil {
        fmt.Println("创建文件描述符: ", err)
        return
    }

    session.Stdout = os.Stdout
    session.Stderr = os.Stderr
    session.Stdin = os.Stdin

    termWidth, termHeight, err := terminal.GetSize(fd)
    if err != nil {
        fmt.Println("获取窗口宽高: ", err)
        return
    }
    defer terminal.Restore(fd, oldState)

    modes := ssh.TerminalModes{
        ssh.ECHO:          1,
        ssh.TTY_OP_ISPEED: 14400,
        ssh.TTY_OP_OSPEED: 14400,
    }

    if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
        fmt.Println("创建终端出错: ", err)
        return
    }
    err = session.Shell()
    if err != nil {
        fmt.Println("执行Shell出错: ", err)
        return
    }
    err = session.Wait()
    if err != nil {
        fmt.Println("执行Wait出错: ", err)
        return
    }
}
