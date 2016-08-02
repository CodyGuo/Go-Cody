package main

import (
	"flag"
	"log"
	"os"
)

import (
	"github.com/secsy/goftp"
)

const (
	HELPEREXE = "hpidmNacSetup.exe"
)

var (
	ftpServer string
	ftpUser   string
	ftpPasswd string
)

func init() {
	flag.StringVar(&ftpServer, "ftpServer", "10.10.2.222", "set ftp server ip.")
	flag.StringVar(&ftpUser, "ftpUser", "iman", "set ftp login user.")
	flag.StringVar(&ftpPasswd, "ftpPasswd", "iman", "set ftp login passwd.")
}

type FTP struct {
	ip     string
	user   string
	passwd string
}

func NewFTP() *FTP {
	ftp := new(FTP)
	ftp.ip = ftpServer
	ftp.user = ftpUser
	ftp.passwd = ftpPasswd

	return ftp
}

func (f *FTP) Upload(fullfile, fileName string) error {
	config := goftp.Config{
		User:            f.user,
		Password:        f.passwd,
		ActiveTransfers: false, // 如果服务器不是主动模式可以不设置
	}
	ftpConn, err := goftp.DialConfig(config, f.ip)
	if err != nil {
		return err
	}
	defer ftpConn.Close()

	// 上传
	uploadFile, err := os.Open(fullfile)
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	return ftpConn.Store(fileName, uploadFile)
}

func main() {
	ftp := NewFTP()
	if err := ftp.Upload(`D:\iMan-SVN\helper\hpidmnac\out\hpidmNacSetup.exe`, HELPEREXE); err != nil {
		log.Fatal(err)
	}
}
