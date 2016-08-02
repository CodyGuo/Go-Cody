package main

import (
    "errors"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    // "sync"
)

import (
    "github.com/dynport/gossh"
    "golang.org/x/crypto/ssh"

    "github.com/pkg/sftp"
)

type UpgradeInfo struct {
    Host   string
    User   string
    Passwd string
    Debug  bool
}

func (info *UpgradeInfo) MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}

// func (info *UpgradeInfo) CheckError(tip string, err error) {
//     if err != nil {
//         // return
//         info.client.ErrorWriter(fmt.Sprintf("{%s} %s", tip, err.Error()))
//     }
// }

type UploadFile struct {
    UpgradeInfo

    timeout time.Duration

    LocalFilePath  string
    UploadFilePath string

    UpgradeFilePath string
    UpgradeFileName string
}

func NewUploadFile() *UploadFile {
    return &UploadFile{}
}

func NewConn() *gossh.Client {
    // client := &gossh.Client{}
    return &gossh.Client{}
}

func (file *UploadFile) UploadiManFile() (err error) {
    client := NewConn()
    client.Host = file.Host
    client.User = file.User
    client.SetPassword(file.Passwd)
    if file.Debug {
        client.DebugWriter = file.MakeLogger("[DEBUG]")
    }

    client.InfoWriter = file.MakeLogger("[INFO]")
    client.ErrorWriter = file.MakeLogger("[ERROR]")

    // 建立一个conn
    conn, err := client.Connection()
    if err != nil {
        client.ErrorWriter("连接SSH失败," + err.Error())
        return
    }
    defer conn.Close()

    // 创建 sftp 连接
    sftpConn, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
    if err != nil {
        client.ErrorWriter("创建SFTP失败," + err.Error())
        return
    }
    defer sftpConn.Close()

    // 本地文件路径
    f, err := os.Open(file.LocalFilePath)
    if err != nil {
        client.ErrorWriter("打开上传文件【" + file.LocalFilePath + "】失败" + err.Error())
        return
    }
    defer f.Close()

    // 要上传的文件信息
    _, file.UpgradeFileName = filepath.Split(file.LocalFilePath)
    info, err := f.Stat()
    if err != nil {
        client.ErrorWriter("获取上传文件【" + file.UpgradeFileName + "】信息失败" + err.Error())
        return
    }
    // 判断文件是否问解压脚本
    if file.UpgradeFileName == "decompress_pack.sh" {
        return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】是非法文件,禁止上传.")
        client.Close()
    }

    // 上传计时
    timeNow := time.Now()

    // 流的方式上传
    client.InfoWriter("正在上传【" + file.UpgradeFileName + "】升级文件,请稍等...")

    size := info.Size()
    unit := "byte"
    switch {
    case size < 1024:
    case size >= 1024 && size < 1024*1024:
        size = size / 1024
        unit = "KB"
    case size >= 1024*1024 && size < 1024*1024*200:
        size = size / 1024 / 1024
        unit = "MB"
    case size >= 1024*1024*200:
        if size < 1024*1024*1024 {
            size = size / 1024 / 1024
            unit = "MB"
        } else {
            size = size / 1024 / 1024 / 1024
            unit = "GB"

        }

        return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】总大小为 " + fmt.Sprint(size) + " " + unit +
            ",超过上传文件的最大限制.")
        client.Close()
    }

    // 服务器升级路径
    destFile := file.UploadFilePath + "/" + file.UpgradeFileName

    w, err := sftpConn.Create(destFile)
    if err != nil {
        client.ErrorWriter("创建上传文件【" + file.UpgradeFileName + "】失败" + err.Error())
        return
    }
    defer w.Close()

    // 开始通过流的方式上传
    n, err := io.Copy(w, io.LimitReader(f, info.Size()))
    if err != nil {
        client.ErrorWriter("获取上传文件【" + file.UpgradeFileName + "】信息失败" + err.Error())
        return
    }
    if n != info.Size() {
        return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】总大小为 " + fmt.Sprint(size) + " " + unit +
            ",已上传 " + fmt.Sprint(n) + " bytes,上传失败.")
    }

    file.timeout = time.Since(timeNow)
    client.InfoWriter("升级文件【" + file.UpgradeFileName + "】上传成功, 总大小为 " + fmt.Sprint(size) + " " + unit +
        ", 用时 " + fmt.Sprint(file.timeout) + ".")

    return nil
}

func (cmd *UploadFile) UpgradeiManCmd(upgrade string) (err error) {
    client := NewConn()
    client.Host = cmd.Host
    client.User = "root"
    client.SetPassword("123456")
    if cmd.Debug {
        client.DebugWriter = cmd.MakeLogger("[DEBUG]")
    }

    client.InfoWriter = cmd.MakeLogger("[INFO]")
    client.ErrorWriter = cmd.MakeLogger("[ERROR]")

    // 建立一个conn
    conn, err := client.Connection()
    if err != nil {
        client.ErrorWriter("连接SSH失败," + err.Error())
        return
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        client.ErrorWriter("创建seesion 失败," + err.Error())
        return
    }
    defer session.Close()

    modes := ssh.TerminalModes{
        ssh.ECHO:          0,     // disable echoing
        ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
        ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
    }
    err = session.RequestPty("vt100", 80, 40, modes)
    if err != nil {
        client.ErrorWriter("session请求VT100失败," + err.Error())
        return
    }

    w, err := session.StdinPipe()
    if err != nil {
        client.ErrorWriter("创建写入通道 失败," + err.Error())
        return
    }

    r, err := session.StdoutPipe()
    if err != nil {
        client.ErrorWriter("创建读取通道 失败," + err.Error())
        return
    }

    err = session.Shell()
    if err != nil {
        client.ErrorWriter("创建seesion shell 失败," + err.Error())
        return
    }

    // 默认超时时间是上传升级包的时间.
    // cmd.timeout = 3

    cmd.sendIn(w, "admin")

    // result, _ := cmd.except("NAC>", r)
    // log.Println(result)

    cmd.sendIn(w, "admin2014")

    // result = except("Password:", r)
    // log.Println(result)

    // sendIn(w, "echo cody.guo")
    _, ok := cmd.except("#", r)
    if ok {
        client.InfoWriter("连接服务器【" + cmd.Host + "】成功.")
    } else {
        return errors.New("[ERROR] 连接服务器【" + cmd.Host + "】失败.")
    }

    switch upgrade {
    case "upgrade":
        // cp 文件
        srcfile := cmd.UploadFilePath + "/" + cmd.UpgradeFileName
        desfile := cmd.UpgradeFilePath + "/" + cmd.UpgradeFileName
        cmd.sendIn(w, "/bin/mv -f "+srcfile+" "+desfile)
        // fmt.Println("默认的超时时间为", cmd.timeout)
        time.Sleep(cmd.timeout)

        // 解压升级包,验证版本
        cmd.sendIn(w, "cd "+cmd.UpgradeFilePath)

        cmd.sendIn(w, "chown root:root "+cmd.UpgradeFileName)

        client.InfoWriter("正在升级【" + cmd.Host + "】请稍等...")
        cmd.sendIn(w, "./decompress_pack.sh "+cmd.UpgradeFileName)

    END:
        for {
            result, ok := cmd.except("#", r)
            if ok {
                switch {
                case strings.Contains(result, "please upload the nac_pack") || strings.Contains(result, "decompress the pack fail") || strings.Contains(result, "You input wrong,Please retry"):
                    return errors.New("[ERROR] 升级【文件】验证失败,请重新上传.")
                    break END
                case strings.Contains(result, "version_check_err"):
                    return errors.New("[ERROR] 升级【文件版本】验证失败,请重新上传.")
                    break END
                case strings.Contains(result, "version_check_ok"):
                    break END
                    // default:
                    //     break END
                }

            } else {
                return errors.New(result)
            }
        }

    case "reboot":
        cmd.sendIn(w, "init 6")
    default:
        return errors.New("[ERROR] 升级指令错误,请联系技术支持.")
    }

    cmd.sendIn(w, "exit")

    cmd.sendIn(w, "exit")

    session.Wait()

    return nil
    // log.Fatalf("本次SSH结束 [%s:%d].", "10.10.3.100", 22)
}

func (cmd *UploadFile) except(sep string, r io.Reader, timeoutSec ...time.Duration) (result string, ok bool) {
    var (
        buf [65 * 1024]byte
        t   int
    )

    if len(timeoutSec) == 1 {
        if timeoutSec[0] != time.Nanosecond {
            cmd.timeout = timeoutSec[0]
        }
    }

    timer := time.NewTimer(time.Second * cmd.timeout)

    for {
        n, err := r.Read(buf[t:])
        // log.Println("正在读取", n)
        if err != nil {
            // log.Println("退出读取", n)
            log.Fatal(err)
        }

        t += n
        tmpLogIn := string(buf[:t])
        if strings.Contains(tmpLogIn, sep) {
            result = string(buf[:t])
            timer.Stop()
            ok = true
            break
        }

        // 超时器
        go func() {
            <-timer.C
            ok = false
            result = "SSH超时: " + fmt.Sprint(cmd.timeout) + " 秒,返回值中不存在: " + sep + "."
        }()
    }

    return
}

func (cmd *UploadFile) sendIn(w io.Writer, command string) {
    w.Write([]byte(command + "\n"))
}
