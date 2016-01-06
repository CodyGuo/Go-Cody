package main

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    // "golang.org/x/crypto/ssh/terminal"
    // "bufio"
    "os"
)

func main() {
    client("root", "123456", "10.10.3.100:22")
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

    session.Stdout = os.Stdout
    session.Stdin = os.Stdin

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
