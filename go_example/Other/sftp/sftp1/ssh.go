package main

import (
    // "fmt"
    "github.com/dynport/gossh"
    "github.com/pkg/sftp"
    "io"
    "log"
    "os"
    "runtime"
    // "syscall"
)

func sshCmd(sqlcmd, srcdir, desdir string, debug bool) {

    client := gossh.New("10.10.2.222", "root")
    // my default agent authentication is used. use
    client.SetPassword("hpiman")
    // for password authentication
    if debug {
        client.DebugWriter = MakeLogger("[DEBUG]")
    }

    client.InfoWriter = MakeLogger("[INFO]")
    client.ErrorWriter = MakeLogger("[ERROR]")

    defer client.Close()

    // 获取备份数据库的名字
    pc, _, _, _ := runtime.Caller(1)

    funcname := runtime.FuncForPC(pc).Name()

    switch funcname {
    case "main.JavaBackup":
        backName = "hupunac.sql"
    case "main.RegisterBackup":
        backName = "licensemanager.sql"
    case "main.WebSiteBackup":
        backName = "hupuwebsite.sql"
    case "main.BusinessBackup":
        backName = "hupuerp.sql"
    case "main.ProcutBackup":
        backName = "hupu.sql"
    case "main.BbsBackup":
        backName = "hupubbs.sql"
    }

    rsp, err := client.Execute(sqlcmd)
    if rsp.Success() {
        client.InfoWriter("执行命令: [备份" + backName + "] 成功.")
    } else {
        checkError(sqlcmd, err)
    }
    defer client.Close()

    conn, err := client.Connection()
    defer conn.Close()

    c, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
    checkError("unable to start sftp subsytem: 10.10.2.222", err)
    defer c.Close()

    wr, err := c.Open(srcdir + backName)
    checkError("unable open the linux file: "+srcdir+backName, err)
    defer wr.Close()

    info, _ := wr.Stat()

    filelocal, err := os.Create(desdir + backName)
    checkError("unable down the file: "+desdir+backName, err)
    defer filelocal.Close()

    client.InfoWriter("正在下载备份文件[" + backName + "]")
    _, err = io.Copy(filelocal, io.LimitReader(wr, info.Size()))
    checkError("unable copy the file: "+desdir+backName, err)

    client.InfoWriter("下载备份文件[" + backName + "]完成\n")
}

func checkError(info string, err error) {
    if err != nil {
        log.Printf("[ERROR] %s [%s] %v", info, err)
    }

}

func MakeLogger(prefix string) gossh.Writer {
    return func(args ...interface{}) {
        log.Println((append([]interface{}{prefix}, args...))...)
    }
}
