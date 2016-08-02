// FTP主动模式上传下载实例
package main

import (
	"log"
	"os"
)

import (
	"github.com/secsy/goftp"
)

func main() {
	config := goftp.Config{
		User:            "iman",
		Password:        "iman",
		ActiveTransfers: true, // 如果服务器不是主动模式可以不设置
	}
	ftpConn, err := goftp.DialConfig(config, "10.10.2.222")
	checkError(err)
	defer ftpConn.Close()

	// 下载
	readme, err := os.Create("eth2.cap")
	checkError(err)
	defer readme.Close()

	err = ftpConn.Retrieve("eth2.cap", readme)
	checkError(err)

	// 上传
	uploadFile, err := os.Open("main.go")
	checkError(err)
	defer uploadFile.Close()

	err = ftpConn.Store("main.go", uploadFile)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
