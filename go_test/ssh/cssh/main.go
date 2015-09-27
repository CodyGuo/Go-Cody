package main

import (
    "io"
    "os"

    "fmt"
    "log"

    "flag"
    "sync"

    "strings"

    "golang.org/x/crypto/ssh"
)

var (
    debugSwitch            bool
    iManUser, iManPassword string

    hostIp, port, userNmae, passWord, shellCmd string
)

func init() {
    flag.StringVar(&hostIp, "h", "", "The remote host ip.")
    flag.StringVar(&port, "port", "22", "The remote host port.")
    flag.StringVar(&userNmae, "u", "", "The host login user name.")
    flag.StringVar(&passWord, "p", "", "The host login user password.")
    flag.StringVar(&iManUser, "iManu", "", "The host login iman user name.")
    flag.StringVar(&iManPassword, "iManp", "", "The host login iman user password.")
    flag.StringVar(&shellCmd, "c", "", "To perform a shell command.")
    flag.BoolVar(&debugSwitch, "d", false, "The debug statu.")
}
func main() {
    flag.Parse()

    if os.Args == nil || hostIp == "" || userNmae == "" ||
        passWord == "" || shellCmd == "" {
        Usage()
    }

    // ssh登录配置,用户名、密码
    config := &ssh.ClientConfig{
        User: userNmae,
        Auth: []ssh.AuthMethod{
            ssh.Password(passWord),
        },
    }

    // 使用tcp的默认22号端口连接ssh
    client, err := ssh.Dial("tcp", hostIp+":"+port, config)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    // Create a session
    session, err := client.NewSession()
    if err != nil {
        log.Fatalf("unable to create session: %s", err)
    }
    defer session.Close()

    // shell终端的模式,vt100
    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
        log.Fatal(err)
    }

    // session的标准输入管道
    w, err := session.StdinPipe()
    if err != nil {
        panic(err)
    }
    // session的标准输出管道
    r, err := session.StdoutPipe()
    if err != nil {
        panic(err)
    }
    // session的错误输出管道
    e, err := session.StderrPipe()
    if err != nil {
        panic(err)
    }

    // channel 读写shell
    in, out := MuxShell(w, r, e)
    if err := session.Shell(); err != nil {
        log.Fatal(err)
    }

    <-out //ignore the shell output

    if iManLogin() {
        in <- iManUser
        in <- iManPassword
    }

    in <- shellCmd

    // 退出系统 shell
    in <- "exit"

    // 退出iMan shell
    if iManLogin() {
        // fmt.Println("iMan 退出系统中...")
        in <- "exit"
    }

    if debugSwitch {
        fmt.Printf("----------正在远程连接: %s:%s,请稍等.------------\n", hostIp, port)
        if iManLogin() {
            fmt.Println("\n############### iMan登录 ####################")
            fmt.Printf("命令执行返回: %s\n", <-out)
            fmt.Printf("\t\t%s\n", <-out)
        }

        fmt.Println("\n############### 执行命令结果 ####################")
        fmt.Printf("命令执行返回: \n%s\n\n", <-out)

        fmt.Println("############### 退出系统      ####################")
        fmt.Printf("命令执行返回: %s\n\n", <-out)
        if iManLogin() {
            fmt.Println("退出系统shell.")
            fmt.Printf("\t\t%s\n", <-out)
        }

        fmt.Println("############### 本次SSH结束   ####################")
    } else {

        if iManLogin() {
            // _, _, _, _, _ = <-out, <-out, <-out, <-out, <-out
            _, _, _ = <-out, <-out, <-out
        } else {
            _ = <-out
        }
    }
    session.Wait()
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
  [cssh -h 192.168.1.1 -port 22 -u root -p root -iManU admin -iManP admin -c "uptime;whoami" -d true]`)
    os.Exit(1)
}

func iManLogin() bool {
    if iManUser == "" && iManPassword == "" {
        return false
    }
    return true
}

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
    in := make(chan string, 3)
    out := make(chan string, 1)
    var wg sync.WaitGroup
    wg.Add(1) //for the shell itself
    go func() {
        for cmd := range in {
            wg.Add(1)
            w.Write([]byte(cmd + "\n"))
            // if debugSwitch {
            //     fmt.Println("执行命令:", cmd)
            // }
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
            tmpLogIn := string(buf[:t])

            // if len(tmpLogIn) > 1 {
            //     fmt.Println("tmpLogIn:", tmpLogIn[:len(tmpLogIn)-1])
            //     fmt.Println("-------------------------------------------------------")
            // }

            //assuming the $PS1 == 'NAC>'; $PS1 == 'Password:'; $PS1 == '[test@hupu ~]$ "; $PS1 == '[root@hupu ~]# '
            if strings.Contains(tmpLogIn, "NAC>") || strings.Contains(tmpLogIn, "Password:") ||
                strings.Contains(tmpLogIn, "$") || strings.Contains(tmpLogIn, "]#") {
                out <- string(buf[:t])
                t = 0
                wg.Done()
            }
        }
    }()
    return in, out
}
