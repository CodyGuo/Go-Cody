package main

import (
    // "fmt"
    "golang.org/x/crypto/ssh"
    // "io"
    "log"
    // "os"
    "bytes"
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
    if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
        log.Fatal(err)
    }
    w, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }

    var b bytes.Buffer
    session.Stdout = &b
    err = session.Shell()
    if err != nil {
        log.Println("show u", err)
    }
    w.Write([]byte("admin\n"))
    w.Write([]byte("admin2014\n"))
    w.Write([]byte("echo 2233 >>1.txt\n"))

    log.Println(b.String())

    session.Wait()
}
