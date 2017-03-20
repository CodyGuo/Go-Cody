package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
)

func main() {
	var host string
	flag.StringVar(&host, "ip", "192.168.1.6", "host ip")
	var user string
	flag.StringVar(&user, "user", "user", "ssh user name")
	var passwd string
	flag.StringVar(&passwd, "passwd", "password", "ssh user's password")
	port := flag.Int("port", 22, "ssh server's port")
	flag.Parse()
	// fmt.Printf("host = %s, user = %s, password = %s, port = %d\n", host, user, passwd, *port)
	// fmt.Printf("numbers = %d\n", flag.NArg())
	var sshHost string
	sshHost = fmt.Sprintf("%s:%d", host, *port)
	// fmt.Printf("host = %s port=%d\n", sshHost, *port)
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGUSR2, syscall.SIGUSR1, syscall.SIGPIPE)
	// client, err := ssh.Dial("tcp", "10.10.2.252:22", &ssh.ClientConfig{
	client, err := ssh.Dial("tcp", sshHost, &ssh.ClientConfig{
		// User: "cisco",
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(passwd)},
		Config: ssh.Config{
			Ciphers: []string{"aes128-cbc"},
		},
	})
	defer client.Close()
	checkErr(err, "dial")

	session, _ := newSession(client, true)
	defer session.Close()
	go func() {
		switch <-sigs {
		case syscall.SIGUSR1:
			session, w := newSession(client, false)
			w.Write([]byte("exit\nexit\nexit\nquit\nquit\nquit\n"))
			session.Close()
			os.Exit(0)
		}
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
	} else {
		w, _ = session.StdinPipe()
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Shell()
	checkErr(err, "start shell")

	return session, w
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)
	}
}
