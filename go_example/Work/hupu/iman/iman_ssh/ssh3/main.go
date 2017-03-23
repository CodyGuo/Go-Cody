package main

import (
    // "bufio"
    // "bytes"
    "fmt"
    "github.com/dynport/gossh"
    "golang.org/x/crypto/ssh"
    "log"
    "os"
    // "strings"
)

// returns a function of type gossh.Writer func(...interface{})
// MakeLogger just adds a prefix (DEBUG, INFO, ERROR)
func MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}

func main() {
    client := gossh.New("10.10.3.227", "root")
    // my default agent authentication is used. use
    client.SetPassword("123456")
    // for password authentication
    client.DebugWriter = MakeLogger("DEBUG")
    client.InfoWriter = MakeLogger("INFO ")
    client.ErrorWriter = MakeLogger("ERROR")
    defer client.Close()

    conn, _ := client.Connection()
    defer conn.Close()
    session, err := conn.NewSession()
    if err != nil {
        panic("Failed to create session: " + err.Error())
    }
    defer session.Close()
    session.Stdout = os.Stdout
    session.Stderr = os.Stderr
    session.Stdin = os.Stdin

    // Set up terminal modes
    modes := ssh.TerminalModes{
        ssh.ECHO:          1,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    // Request pseudo terminal
    if err := session.RequestPty("xterm", 25, 100, modes); err != nil {
        log.Fatalf("request for pseudo terminal failed: %s", err)
    }
    // Start remote shell
    if err := session.Shell(); err != nil {
        log.Fatalf("failed to start shell: %s", err)
    }
    fmt.Sprint("admin\n")
    fmt.Sprint("admin2015\n")
    err = session.Wait()
    if err != nil {
        log.Fatal(err)
    }

}
