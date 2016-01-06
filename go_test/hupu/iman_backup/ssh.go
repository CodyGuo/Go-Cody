package main

import (
	"fmt"
	"github.com/dynport/gossh"
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
)

const (
	host     = "10.10.2.222"
	user     = "root"
	password = "hpiman"
)

func sshConn(bak *Backup) {
	fmt.Println("--------------------------------------------------------")

	client := gossh.New(host, user)
	defer client.Close()
	client.SetPassword(password)

	// client.DebugWriter = MakeLogger("[DEBUG]")
	client.InfoWriter = MakeLogger("[INFO]")
	client.ErrorWriter = MakeLogger("[ERROR]")

	client.Info("正在执行命令: 备份【" + bak.fileName + "】,请耐心等待.")
	rsp, err := client.Execute(bak.cmd)
	if rsp.Success() {
		client.Info("执行命令: 备份【" + bak.fileName + "】成功.")
	} else {
		checkError(err)
	}

	conn, err := client.Connection()
	checkError(err)
	defer conn.Close()

	sftpConn, err := sftp.NewClient(conn, sftp.MaxPacket(1<<15))
	checkError(err)

	var filePath string
	if bak.fileName == "zentaopms.tar.gz" {
		srcSqlPath = "/var/www"
		filePath = bak.pathName + "/"
	} else {
		filePath = MkBakDir(bak.pathName + "/SQL/")
	}

	srcFile, err := sftpConn.Open(srcSqlPath + "/" + bak.fileName)
	defer srcFile.Close()
	checkError(err)

	info, _ := srcFile.Stat()
	dstFile, err := os.Create(filePath + bak.fileName)
	defer dstFile.Close()
	checkError(err)

	client.Info("正在下载备份的文件【" + bak.fileName + "】,请耐心等待.")
	_, err = io.Copy(dstFile, io.LimitReader(srcFile, info.Size()))
	checkError(err)

	client.Info("下载备份文件【" + bak.fileName + "】完成.\n")
}

func MakeLogger(prefix string) gossh.Writer {
	return func(args ...interface{}) {
		log.Println((append([]interface{}{prefix}, args...))...)
	}
}
