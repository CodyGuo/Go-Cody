/*
ssh 连接双层验证的linux shell

Last login: Sat Sep 26 20:11:58 2015 from 192.168.119.1

                Hello, this is test cli (version 2.0).
                Copyright 2002-2013 test Tech. Co., Ltd. All rights reserved.

Login> admin
Password:
[root@local ~]#

*/
package main

import (
	"flag"
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

var (
	sshConf SSHConfig
)

func init() {
	sshConf.setFlag()

}

func main() {
	flag.Parse()
	sshConf.Login()
	sshConf.Do()
	sshConf.Exit()
	sshConf.LogWrite()
	sshConf.session.Wait()
}

type SSHConfig struct {
	host       string
	port       string
	user       string
	passwd     string
	iManUser   string
	iManPasswd string
	shellCmd   string
	cmd        []string
	debug      bool

	session *ssh.Session

	in  chan<- string
	out <-chan string
}

func (s *SSHConfig) setFlag() {
	flag.StringVar(&s.host, "h", "", "The remote host ip.")
	flag.StringVar(&s.port, "port", "22", "The remote host port.")
	flag.StringVar(&s.user, "u", "", "The user name login host.")
	flag.StringVar(&s.passwd, "p", "", "Log on to the host password.")
	flag.StringVar(&s.iManUser, "iManu", "", "The user name login iMan.")
	flag.StringVar(&s.iManPasswd, "iManp", "", "Login iMan's password.")
	flag.StringVar(&s.shellCmd, "c", "", "To perform a shell command.")
	flag.BoolVar(&s.debug, "d", false, "The debug switch.")
}

func (s *SSHConfig) Login() {
	if os.Args == nil || s.host == "" || s.user == "" ||
		s.passwd == "" || len(s.shellCmd) == 0 {
		s.host = "10.10.3.100"
		s.port = "22"
		s.user = "root"
		s.passwd = "123456"
		s.cmd = []string{"show p", "show u"}
		s.debug = true
	} else {
		s.port = "22"
		s.cmd = strings.Split(s.shellCmd, ";")
	}

	// ssh登录配置,用户名、密码
	config := &ssh.ClientConfig{
		User: s.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.passwd),
		},
	}

	fmt.Println(s.host, s.port, s.user, s.passwd, s.iManUser, s.iManPasswd)
	// 使用tcp的默认22号端口连接ssh
	client, err := ssh.Dial("tcp", s.host+":"+s.port, config)
	CheckError(err)
	defer client.Close()

	// Create a session
	s.session, err = client.NewSession()
	CheckError(err)
	defer s.session.Close()

	// shell终端的模式,vt100
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	err = s.session.RequestPty("vt100", 80, 40, modes)
	CheckError(err)

	// session的标准输入管道
	w, err := s.session.StdinPipe()
	CheckError(err)
	// session的标准输出管道
	r, err := s.session.StdoutPipe()
	CheckError(err)
	// session的错误输出管道
	e, err := s.session.StderrPipe()
	CheckError(err)

	// channel 读写shell
	s.in, s.out = MuxShell(w, r, e)
	err = s.session.Shell()
	CheckError(err)

	//等待shell返回
	<-s.out

	// 登录iMan系统
	if iManLogin() {
		s.in <- s.iManUser
		s.in <- s.iManPasswd
	}
}

func (s *SSHConfig) Exit() {
	// 退出系统 shell
	s.in <- "exit"

	// 退出iMan shell
	if iManLogin() {
		// fmt.Println("iMan 退出系统中...")
		s.in <- "exit"
	}
}

func (s *SSHConfig) Do() {
	for _, cmd := range s.cmd {
		fmt.Println("正在发送命令: ", cmd)
		s.in <- cmd
	}
}

func (s *SSHConfig) LogWrite() {
	if s.debug {
		fmt.Printf("----------正在远程连接: %s:%s,请稍等.------------\n", s.host, s.port)
		if iManLogin() {
			fmt.Println("\n############### iMan登录 ####################")
			fmt.Printf("命令执行返回: %s\n", <-s.out)
			fmt.Printf("\t\t%s\n", <-s.out)
		}

		fmt.Println("\n############### 执行命令结果 ####################")
		for range s.cmd {
			fmt.Printf("命令执行返回: \n%s\n\n", <-s.out)
		}

		fmt.Println("############### 退出系统      ####################")
		fmt.Printf("命令执行返回: %s\n\n", <-s.out)
		if iManLogin() {
			fmt.Println("退出系统shell.")
			fmt.Printf("\t\t%s\n", <-s.out)
		}

		fmt.Println("############### 本次SSH结束   ####################")
	} else {
		if iManLogin() {
			_, _, _ = <-s.out, <-s.out, <-s.out
		} else {
			_ = <-s.out
		}
	}
}

func Usage() {
	fmt.Printf(`Usage of cssh:
  -c string
        To perform a shell command.
  -d bool The debug statu.
  -h string
        The remote host ip.
  -iManp string
        The host login iman user password.
  -iManu string
        The host login iman user name.
  -p string
        The host login user password.
  -port string
        The remote host port. (default "22")
  -u string
        The host login user name.
  [cssh -h 192.168.1.1 -port 22 -u root -p root -iManu admin -iManp admin -c "uptime;whoami" -d true]`)
	os.Exit(1)
}

func iManLogin() bool {
	if sshConf.iManUser == "" && sshConf.iManPasswd == "" {
		return false
	}
	return true
}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 5)
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
			buf [1 << 15]byte
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

			tmpLogIn := string(buf[:t])
			//assuming the $PS1 == 'Login>'; $PS1 == 'Password:'; $PS1 == '[test@local ~]$ "; $PS1 == '[root@local ~]# '
			if strings.Contains(tmpLogIn, "Login>") || strings.Contains(tmpLogIn, "Password:") ||
				strings.Contains(tmpLogIn, "$") || strings.Contains(tmpLogIn, "~#") ||
				strings.Contains(tmpLogIn, "NAC>") {
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "iMan[ERROR]: %s\n", err.Error())
		os.Exit(1)
	}
}
