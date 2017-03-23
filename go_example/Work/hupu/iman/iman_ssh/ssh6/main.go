package main

import (
    "bufio"
    "fmt"
    "golang.org/x/crypto/ssh"
    "golang.org/x/crypto/ssh/agent"
    "io"
    "log"
    "net"
    "os"
    "path/filepath"
    "regexp"
    "sftp"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("参数不正确：")
        fmt.Printf("上传文件：%s %s %s\n", os.Args[0], "/data/Test.txt", "127.0.0.1:/mnt/")
        fmt.Printf("下载文件：%s %s %s\n", os.Args[0], "127.0.0.1:/mnt/Test.txt", "/data/")
        os.Exit(1)
    }
    _, err := os.Stat(os.Args[1])
    if err != nil {
        get()
        return
    }
    push()
}

func get() {
    localpath := strings.Replace(os.Args[2], `\`, `/`, 20)
    re, _ := regexp.Compile(`((([1-9]?|1\d)\d|2([0-4]\d|5[0-5]))\.){3}(([1-9]?|1\d)\d|2([0-4]\d|5[0-5]))`)
    if re.MatchString(localpath) {
        fmt.Println("检查文件源地址.")
        return
    }
    list := strings.Split(os.Args[1], ":")
    if len(list) != 2 {
        fmt.Println("检查文件源地址.")
        return
    }
    fmt.Println("文件下载.")
    Sftp(list[0], list[1], localpath, "get")
}

func push() {
    list := strings.Split(os.Args[2], ":")
    if len(list) != 2 {
        fmt.Println("检查文件源地址.")
        return
    }
    fmt.Println("文件推送.")
    localpath := strings.Replace(list[1], `\`, `/`, 20)
    Sftp(list[0], os.Args[1], localpath, "push")
}

func Sftp(ip, path, localpath, action string) {
    var auths []ssh.AuthMethod
    if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
        auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))

    }
    auth_list := parse("passlist", ip)
    if len(auth_list) == 0 {
        log.Println("获取用户和密码失败.")
        os.Exit(-1)
    }
    auths = append(auths, ssh.Password(auth_list[2]))

    config := ssh.ClientConfig{
        User: auth_list[1],
        Auth: auths,
    }
    addr := fmt.Sprintf("%s:%s", ip, auth_list[3])
    conn, err := ssh.Dial("tcp", addr, &config)
    if err != nil {
        log.Fatalf("unable to connect to [%s]: %v", addr, err)
    }
    defer conn.Close()
    c, err := sftp.NewClient(conn, sftp.MaxPacket(6e9))
    if err != nil {
        log.Fatalf("unable to start sftp subsytem: %v", err)
    }
    defer c.Close()
    switch action {
    case "get":
        fs, e := c.Open(path)
        if e != nil {
            log.Println(e)
            os.Exit(-1)
        }
        filename := filepath.Base(path)
        info, _ := fs.Stat()
        File, err := os.Create(fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
        if err != nil {
            log.Println(err)
            os.Exit(-1)
        }
        log.Println("保存路径:", fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
        defer File.Close()
        io.Copy(File, io.LimitReader(fs, info.Size()))
    case "push":
        filename := filepath.Base(path)
        fs, err := c.Create(fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
        if err != nil {
            log.Println(err)
            os.Exit(-1)
        }
        defer fs.Close()
        log.Println("保存路径:", fmt.Sprintf(`%s/%s`, strings.TrimRight(localpath, `/`), filename))
        File, err := os.Open(path)
        if err != nil {
            log.Println(err)
            os.Exit(-1)
        }
        defer File.Close()
        io.Copy(io.MultiWriter(fs), File)
    }
}

func parse(cfg, ip string) []string {
    File, err := os.Open(cfg)
    if err != nil {
        log.Println("打开配置文件失败.")
        os.Exit(-1)
    }
    defer File.Close()
    buf := bufio.NewReader(File)
    for {
        line, _, err := buf.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            break
        }
        if strings.Contains(string(line), ip) {
            return split(string(line))
        }
    }
    return []string{}
}

func split(str string) []string {
    var l []string
    list := strings.Split(str, " ")
    for _, v := range list {
        if len(v) == 0 {
            continue
        }
        if strings.Contains(v, "    ") {

            list := strings.Split(v, "  ")
            for _, v := range list {
                if len(v) == 0 {
                    continue
                }
                l = append(l, v)
            }
            continue
        }
        l = append(l, v)
    }
    return l
}
