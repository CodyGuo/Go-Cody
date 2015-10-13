package main

import (
    "log"
    "os/exec"
    "strings"
)

type SvnUrl struct {
    Svn     string
    Version int
}

func (s SvnUrl) Get(svnUrl, svnName string) {
    log.Println("[INFO] begin get " + svnName + " for svn.")

    cmdList := "svn checkout " + svnUrl + " --username=cody.guo --password=iman2015 --no-auth-cache " + svnName
    list := strings.Split(cmdList, " ")
    cmd := exec.Command(list[0], list[1:]...)

    err := cmd.Run()
    if err != nil {
        log.Fatal("[ERROR] get svn err. ", svnUrl, list, err)
    } else {
        log.Println("[INFO] " + svnName + " 检出成功.")
    }

    log.Println("[INFO] end get " + svnName + " for svn.\n")
}

func (s SvnUrl) Update(svnUrl, svnName string) {
    log.Println("[INFO] begin update " + svnName + " for svn.")

    cmdList := "svn update --username=cody.guo --password=iman2015 --no-auth-cache " + svnName
    list := strings.Split(cmdList, " ")
    cmd := exec.Command(list[0], list[1:]...)

    err := cmd.Run()
    if err != nil {
        log.Fatal("[ERROR] update svn err. ", svnUrl, list, err)
    } else {
        log.Println("[INFO] " + svnName + " 更新成功.")
    }

    log.Println("[INFO] end update " + svnName + " for svn.\n")
}
