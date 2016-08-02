package main

import (
    "fmt"
)

func main() {

    Upload := NewUploadFile()
    Upload.Host = "10.10.3.100"
    Upload.User = "nacupdate"
    Upload.Passwd = "imanupdate"
    Upload.Debug = true
    fmt.Println(Upload.Host, Upload.User, Upload.Passwd)

    Upload.UpgradeFile = "var.g"
    Upload.DestFile = "/home/nacupdate/" + Upload.UpgradeFile

    err := Upload.UploadiManFile()
    Upload.CheckError("上传文件失败.", err)
}
