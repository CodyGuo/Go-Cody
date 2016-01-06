package main

import (
    "bufio"
    "fmt"
    "golang.org/x/crypto/ssh"
    "log"
    "os"
)

func main() {
    check := func(err error, msg string) {
        if err != nil {
            log.Fatalf("%s error: %v", msg, err)
        }
    }

    client, err := ssh.Dial("tcp", "10.10.3.100:22", &ssh.ClientConfig{
        User: "root",
        Auth: []ssh.AuthMethod{ssh.Password("123456")},
    })
    check(err, "dial")

    session, err := client.NewSession()
    check(err, "new session")
    defer session.Close()

    session.Stdout = os.Stdout
    session.Stderr = os.Stderr
    session.Stdin = os.Stdin

    modes := ssh.TerminalModes{
        ssh.ECHO:          1,
        ssh.TTY_OP_ISPEED: 14400,
        ssh.TTY_OP_OSPEED: 14400,
    }
    err = session.RequestPty("xterm", 25, 100, modes)
    check(err, "request pty")

    err = session.Shell()
    check(err, "start shell")

    // // session.Stdout
    stdout, _ := session.StdoutPipe()
    r := bufio.NewReader(stdout)
    line, _, err := r.ReadLine()
    fmt.Println(line)

    err = session.Wait()
    check(err, "return")
}
