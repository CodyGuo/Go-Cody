package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	client, err := ssh.Dial("tcp", "10.10.2.252:22", &ssh.ClientConfig{
		User: "cisco",
		Auth: []ssh.AuthMethod{ssh.Password("cisco")},
		Config: ssh.Config{
			Ciphers: []string{"aes128-cbc"},
		},
	})
	defer client.Close()
	checkErr(err, "dial")

	session, _ := newSession(client, true)
	defer session.Close()

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("\ntimeout, exit session.")
		session.Close()
		session, w := newSession(client, false)
		w.Write([]byte("exit\n"))

		session.Close()
	}()

	err = session.Wait()
	checkErr(err, "return")
}

func newSession(client *ssh.Client, stdin bool) (*ssh.Session, io.WriteCloser) {
	session, err := client.NewSession()
	checkErr(err, "new session")

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", 25, 80, modes)
	checkErr(err, "request pty")

	var w io.WriteCloser
	if stdin {
		session.Stdin = os.Stdin
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
	} else {
		w, _ = session.StdinPipe()
	}

	err = session.Shell()
	checkErr(err, "start shell")

	return session, w
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)
	}
}
