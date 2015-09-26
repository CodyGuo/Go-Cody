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
    debugOn      = flag.Bool("d", false, "The debug statu.")
    hostIp       = flag.String("h", "", "The remote host ip.")
    port         = flag.String("port", "22", "The remote host port.")
    userNmae     = flag.String("u", "", "The host login user name.")
    passWord     = flag.String("p", "", "The host login user password.")
    iManUser     = flag.String("iManu", "", "The host login iman user name.")
    iManPassword = flag.String("iManp", "", "The host login iman user password.")
    shellCmd     = flag.String("c", "", "To perform a shell command.")
)

func main() {
    flag.Parse()

    if os.Args == nil || *hostIp == "" || *userNmae == "" || *passWord == "" ||
        *shellCmd == "" {
        Usage()
    }

    config := &ssh.ClientConfig{
        User: *userNmae,
        Auth: []ssh.AuthMethod{
            ssh.Password(*passWord),
        },
    }
    client, err := ssh.Dial("tcp", *hostIp+":"+*port, config)
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
    in <- *iManUser
    in <- *iManPassword
    in <- *shellCmd

    // 退出系统 shell
    in <- "exit"

    // 退出iMan shell
    in <- "exit"

    if *debugOn {
        fmt.Printf("----------正在远程连接: %s,请稍等.------------\n", *hostIp)
        fmt.Println("############### 登录 ####################")
        fmt.Printf("命令执行返回: %s\n %s\n", <-out, <-out)
        fmt.Print("\n")

        fmt.Println("############### 执行命令结果 ####################")
        fmt.Printf("命令执行返回: %s\n", <-out)
        fmt.Print("\n")

        fmt.Println("############### 退出系统 ####################")
        fmt.Printf("命令执行返回: \n%s\n %s\n", <-out, <-out)
        fmt.Println("############### 本次SSH结束 ####################")
    } else {
        _, _, _, _, _ = <-out, <-out, <-out, <-out, <-out
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

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
    in := make(chan string, 3)
    out := make(chan string, 5)
    var wg sync.WaitGroup
    wg.Add(1) //for the shell itself
    go func() {
        for cmd := range in {
            wg.Add(1)
            w.Write([]byte(cmd + "\n"))
            // if debugOn {
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

            if strings.Contains(tmpLogIn, "NAC>") || strings.Contains(tmpLogIn, "Password:") ||
                strings.Contains(tmpLogIn, "$") || strings.Contains(tmpLogIn, "]#") { //assuming the $PS1 == 'sh-4.3$ '
                out <- string(buf[:t])
                t = 0
                wg.Done()
            }
        }
    }()
    return in, out
}
