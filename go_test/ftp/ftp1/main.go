package main

import "fmt"
import "os"
import "io/ioutil"
import "github.com/codyguo/ftp"

func main() {
    // new ftp
    ftpC := new(ftp.FTP)
    // set debug, default false
    ftpC.Debug = true
    // connect
    ftpC.Connect("10.10.3.100", 21)
    // login
    ftpC.Login("nac_ftp", "qaz!@#")
    // login failure
    if ftpC.Code == 530 {
        fmt.Println("error: login failure")
        os.Exit(-1)
    }
    // pwd
    ftpC.Pwd()

    fmt.Println("code:", ftpC.Code, ", message:", ftpC.Message)

    ftpC.Request(cmd)
    // mkdir new dir
    ftpC.Mkd("/smallfish")
    b, _ := ioutil.ReadFile("main.go")
    ftpC.Stor("/smallfish/a.txt", b)
    // quit
    ftpC.Quit()
}

//该代码片段来自于: http://www.sharejs.com/codes/go/8663
