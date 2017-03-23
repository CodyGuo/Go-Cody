package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
)

func CopyFile(src, des string) (err error) {
    fmt.Println(src)
    srcFile, srcerr := os.Open(src)
    if srcerr != nil {
        log.Fatal(srcerr)
    }
    defer srcFile.Close()

    dstFile, dsterr := os.Create(des)
    if dsterr != nil {
        log.Fatal(dsterr)
    }

    defer dstFile.Close()

    _, ioerr := io.Copy(dstFile, srcFile)

    return ioerr
}

func GetRpmInfo(rpm string) (err error) {
    tftpInfo := exec.Command("rpm", "-q", rpm)
    infoerr := tftpInfo.Run()
    return infoerr
}

func InstallRpm(rpm string) (err error) {
    rmpInstall := exec.Command("rpm", "-i", rpm)
    installerr := rmpInstall.Run()
    return installerr
}

func StartServers(s string) {
    Server := exec.Command("service", s, "restart")
    serr := Server.Run()
    if serr != nil {
        log.Fatal(serr)
    }
}
