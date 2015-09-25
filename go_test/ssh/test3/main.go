package main

import (
    "fmt"
    "golang.org/x/crypto/ssh"
    "io"
    "log"
    // "strings"
    "sync"
)

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
    in := make(chan string, 1)
    out := make(chan string, 4)
    var wg sync.WaitGroup
    wg.Add(1) //for the shell itself
    go func() {
        for cmd := range in {
            wg.Add(1)
            w.Write([]byte(cmd + "\n"))
            wg.Wait()
        }
    }()

    go func() {
        var (
            buf [65 * 1024]byte
            t   int
        )
        for {
            n, err := r.Read(buf[t:])
            if err != nil {
                fmt.Println(err.Error())
                close(in)
                close(out)
                return
            }

            t += n
            tmpStr := string(buf[t-2])
            if tmpStr == ">" {
                w.Write([]byte("admin\n"))
            }

            fmt.Println("tmpstr[len -1]:", tmpStr)
            fmt.Println("-------------------------------------------------------")

            if tmpStr == "#" { //assuming the $PS1 == 'sh-4.3$ '
                out <- string(buf[:t])
                t = 0
                wg.Done()
            }
        }
    }()
    return in, out
}
func main() {
    config := &ssh.ClientConfig{
        User: "root",
        // User: "nacupdate",
        Auth: []ssh.AuthMethod{
            ssh.Password("123456"),
            // ssh.Password("imanupdate"),
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
    e, err := session.StderrPipe()
    if err != nil {
        panic(err)
    }
    in, out := MuxShell(w, r, e)
    if err := session.Shell(); err != nil {
        log.Fatal(err)
    }
    <-out //ignore the shell output
    // w.Write([]byte("admin\n"))
    // w.Write([]byte("admin2014\n"))
    // in <- "admin"
    // in <- "admin2014"
    // in <- "admin2014"
    in <- "whoami"
    fmt.Printf("whoami: %s\n", <-out)
    // fmt.Printf("whoami: %s\n", <-out)
    // fmt.Printf("whoami: %s\n", <-out)
    in <- "exit"
    session.Wait()
}
