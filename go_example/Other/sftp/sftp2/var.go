package main

import (
    "errors"
    "fmt"
    "github.com/dynport/gossh"
    "github.com/pkg/sftp"
    "io"
    "log"
    "os"
    "time"
)

var (
    Upload *UploadFile
)

type UpgradeInfo struct {
    Host   string
    User   string
    Passwd string
    Debug  bool

    client gossh.Client
}

// type UpgredeCmd struct {
//     UpgradeInfo
// }

type UploadFile struct {
    UpgradeInfo

    UpgradeFile string
    DestFile    string
}

func NewUploadFile() *UploadFile {
    return &UploadFile{}
}

func (info *UpgradeInfo) MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}

func (info *UpgradeInfo) CheckError(tip string, err error) {
    if err != nil {
        info.client.ErrorWriter(fmt.Sprintf("{%s} %s", tip, err.Error()))
        os.Exit(1)
    }
}

func (file *UploadFile) UploadiManFile() (err error) {
    file.client.Host = file.Host
    file.client.User = file.User
    file.client.SetPassword(file.Passwd)
    if file.Debug {
        file.client.DebugWriter = file.MakeLogger("[DEBUG]")
    }

    file.client.InfoWriter = file.MakeLogger("[INFO]")
    file.client.ErrorWriter = file.MakeLogger("[ERROR]")

    defer file.client.Close()

    // 建立一个conn
    conn, err := file.client.Connection()

    // 创建 sftp 连接
    sftpConn, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
    file.CheckError("unable to start sftp subsytem: 10.10.2.222", err)
    defer sftpConn.Close()

    // 本地文件路径
    f, err := os.Open(file.UpgradeFile)
    file.CheckError("打开上传文件["+file.UpgradeFile+"]失败", err)
    defer f.Close()

    // 上传到的位置
    w, err := sftpConn.Create(file.DestFile)
    file.CheckError("创建上传文件["+file.DestFile+"]失败", err)
    defer w.Close()

    // 要上传的文件信息
    info, err := f.Stat()
    file.CheckError("获取上传文件["+file.UpgradeFile+"]信息失败", err)

    // 上传计时
    timeNow := time.Now()

    // 流的方式上传
    n, err := io.Copy(w, io.LimitReader(f, info.Size()))
    file.CheckError("上传文件["+file.UpgradeFile+"]失败.", err)
    if n != info.Size() {
        return errors.New("上传文件[" + file.UpgradeFile + "]总大小为 " + fmt.Sprint(info.Size()) + " bytes, 已上传 " + fmt.Sprint(n) + " bytes,上传失败.")
    }

    file.client.InfoWriter("上传文件[" + file.UpgradeFile + "]成功, 总大小为 " + fmt.Sprint(info.Size()) + ", 用时 " + fmt.Sprint(time.Since(timeNow)))

    return nil
}
