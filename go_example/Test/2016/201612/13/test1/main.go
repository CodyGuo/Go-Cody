package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	client, err := ssh.Dial("tcp", "10.10.2.252:22", &ssh.ClientConfig{
		User: "cisco",
		Auth: []ssh.AuthMethod{ssh.Password("cisco")},
		Config: ssh.Config{
			Ciphers: []string{"aes128-cbc"},
		},
	})
	ce(err, "dial")

	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", 25, 80, modes)
	ce(err, "request pty")

	err = session.Shell()
	ce(err, "start shell")

	err = session.Wait()
	ce(err, "return")
}
