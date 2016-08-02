package main

import (
    // "errors"
    // "fmt"
    // "io"
    "log"
    // "os"
    // "path/filepath"
    // "strings"
    // "time"
)

import (
    "github.com/dynport/gossh"
    // "golang.org/x/crypto/ssh"

    // "github.com/pkg/sftp"
)

type SshConfig struct {
    Host   string
    User   string
    Passwd string
    Debug  bool
}

func NewSshConfig() *SshConfig {
    return new(SshConfig)
}

func (s *SshConfig) MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}

func (s *SshConfig) NewSeesion() {
    sshConf := NewSshConfig()
    sshConf.Host = "192.168.1.1"
}

// type UploadFile struct {
//     UpgradeInfo

//     timeout time.Duration

//     LocalFilePath  string
//     UploadFilePath string

//     UpgradeFilePath string
//     UpgradeFileName string
// }

// func NewUploadFile() *UploadFile {
//     return &UploadFile{}
// }

// func NewConn() *gossh.Client {
//     // client := &gossh.Client{}
//     return &gossh.Client{}
// }

// func (file *UploadFile) UploadiManFile() (err error) {
//     client := NewConn()
//     client.Host = file.Host
//     client.User = file.User
//     client.SetPassword(file.Passwd)
//     if file.Debug {
//         client.DebugWriter = file.MakeLogger("[DEBUG]")
//     }

//     client.InfoWriter = file.MakeLogger("[INFO]")
//     client.ErrorWriter = file.MakeLogger("[ERROR]")

//     // 建立一个conn
//     conn, err := client.Connection()
//     if err != nil {
//         client.ErrorWriter("连接SSH失败," + err.Error())
//         return
//     }
//     defer conn.Close()

//     // 创建 sftp 连接
//     sftpConn, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
//     if err != nil {
//         client.ErrorWriter("创建SFTP失败," + err.Error())
//         return
//     }
//     defer sftpConn.Close()

//     // 本地文件路径
//     f, err := os.Open(file.LocalFilePath)
//     if err != nil {
//         client.ErrorWriter("打开上传文件【" + file.LocalFilePath + "】失败" + err.Error())
//         return
//     }
//     defer f.Close()

//     // 要上传的文件信息
//     _, file.UpgradeFileName = filepath.Split(file.LocalFilePath)
//     info, err := f.Stat()
//     if err != nil {
//         client.ErrorWriter("获取上传文件【" + file.UpgradeFileName + "】信息失败" + err.Error())
//         return
//     }
//     // 判断文件是否问解压脚本
//     if file.UpgradeFileName == "decompress_pack.sh" {
//         return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】是非法文件,禁止上传.")
//         client.Close()
//     }

//     // 上传计时
//     timeNow := time.Now()

//     // 流的方式上传
//     client.InfoWriter("正在上传【" + file.UpgradeFileName + "】升级文件,请稍等...")

//     size := info.Size()
//     unit := "byte"
//     switch {
//     case size < 1024:
//     case size >= 1024 && size < 1024*1024:
//         size = size / 1024
//         unit = "KB"
//     case size >= 1024*1024 && size < 1024*1024*200:
//         size = size / 1024 / 1024
//         unit = "MB"
//     case size >= 1024*1024*200:
//         if size < 1024*1024*1024 {
//             size = size / 1024 / 1024
//             unit = "MB"
//         } else {
//             size = size / 1024 / 1024 / 1024
//             unit = "GB"

//         }

//         return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】总大小为 " + fmt.Sprint(size) + " " + unit +
//             ",超过上传文件的最大限制.")
//         client.Close()
//     }

//     // 服务器升级路径
//     destFile := file.UploadFilePath + "/" + file.UpgradeFileName

//     w, err := sftpConn.Create(destFile)
//     if err != nil {
//         client.ErrorWriter("创建上传文件【" + file.UpgradeFileName + "】失败" + err.Error())
//         return
//     }
//     defer w.Close()

//     // 开始通过流的方式上传
//     n, err := io.Copy(w, io.LimitReader(f, info.Size()))
//     if err != nil {
//         client.ErrorWriter("获取上传文件【" + file.UpgradeFileName + "】信息失败" + err.Error())
//         return
//     }
//     if n != info.Size() {
//         return errors.New("[ERROR] 升级文件【" + file.UpgradeFileName + "】总大小为 " + fmt.Sprint(size) + " " + unit +
//             ",已上传 " + fmt.Sprint(n) + " bytes,上传失败.")
//     }

//     file.timeout = time.Since(timeNow)
//     client.InfoWriter("升级文件【" + file.UpgradeFileName + "】上传成功, 总大小为 " + fmt.Sprint(size) + " " + unit +
//         ", 用时 " + fmt.Sprint(file.timeout) + ".")

//     return nil
// }

// func (cmd *UploadFile) except(sep string, r io.Reader, timeoutSec ...time.Duration) (result string, ok bool) {
//     var (
//         buf [65 * 1024]byte
//         t   int
//     )

//     if len(timeoutSec) == 1 {
//         if timeoutSec[0] != time.Nanosecond {
//             cmd.timeout = timeoutSec[0]
//         }
//     }

//     timer := time.NewTimer(time.Second * cmd.timeout)

//     for {
//         n, err := r.Read(buf[t:])
//         // log.Println("正在读取", n)
//         if err != nil {
//             // log.Println("退出读取", n)
//             log.Fatal(err)
//         }

//         t += n
//         tmpLogIn := string(buf[:t])
//         if strings.Contains(tmpLogIn, sep) {
//             result = string(buf[:t])
//             timer.Stop()
//             ok = true
//             break
//         }

//         // 超时器
//         go func() {
//             <-timer.C
//             ok = false
//             result = "SSH超时: " + fmt.Sprint(cmd.timeout) + " 秒,返回值中不存在: " + sep + "."
//         }()
//     }

//     return
// }

// func (cmd *UploadFile) sendIn(w io.Writer, command string) {
//     w.Write([]byte(command + "\n"))
// }
