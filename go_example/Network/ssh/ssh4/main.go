/* switch ssh
username cisco password cisco
ip domain name hupu
crypto key generate rsa
ip ssh time-out 30
line vty 0 4
transport input all
login local
*/
package main

import (
	"fmt"
	"io"
	// "log"
	"os"
	"strings"
	"sync"
)

import (
	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "cisco",
		Auth: []ssh.AuthMethod{
			ssh.Password("cisco"),
		},
		Config: ssh.Config{
			Ciphers: []string{"aes128-cbc"},
		},
	}
	// config.Config.Ciphers = append(config.Config.Ciphers, "aes128-cbc")
	clinet, err := ssh.Dial("tcp", "192.168.1.253:22", config)
	checkError(err, "连接交换机")

	session, err := clinet.NewSession()
	defer session.Close()
	checkError(err, "创建shell")

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
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
	in <- "show arp"
	in <- "show int status"

	in <- "exit"
	in <- "exit"

	fmt.Printf("%s\n%s\n", <-out, <-out)

	_, _ = <-out, <-out
	// tmp, _ := session.Output("show arp")
	// fmt.Println(string(tmp))
	session.Wait()
}

func checkError(err error, info string) {
	if err != nil {
		fmt.Printf("%s. error: %s\n", info, err)
		os.Exit(1)
	}
}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 3)
	out := make(chan string, 5)
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
			result := string(buf[:t])
			if strings.Contains(result, "Username:") ||
				strings.Contains(result, "Password:") ||
				strings.Contains(result, "#") {
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}
