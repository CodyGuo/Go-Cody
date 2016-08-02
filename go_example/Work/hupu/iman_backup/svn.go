package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Svn struct {
	svnUrl      string
	svnName     string
	pathName    string
	loginUser   string
	loginPasswd string
}

func (s *Svn) Get() {
	fmt.Println("--------------------------------------------------------")

	log.Println("[INFO] begin get 【" + s.svnName + "】for svn.")
	cmdList := "svn checkout " + s.svnUrl + " --username=" + s.loginUser + " --password=" +
		s.loginPasswd + " --no-auth-cache " + s.pathName
	list := strings.Split(cmdList, " ")
	cmd := exec.Command(list[0], list[1:]...)

	err := cmd.Run()
	if err != nil {
		if s.svnName != "licenseManager" {
			log.Fatal("[ERROR] get svn err. ", s.svnUrl, list, err)
		} else {
			log.Println("[INFO] 【" + s.svnName + "】检出成功.")
		}
	} else {
		log.Println("[INFO] 【" + s.svnName + "】检出成功.")
	}

	cmd.Wait()

	log.Println("[INFO] end get 【" + s.svnName + "】for svn.\n")
}

func (s *Svn) Update() {
	fmt.Println("--------------------------------------------------------")

	log.Println("[INFO] begin update 【" + s.svnName + "】for svn.")

	cmdList := "svn update --username=" + s.loginUser + " --password=" +
		s.loginPasswd + " --no-auth-cache " + s.pathName
	list := strings.Split(cmdList, " ")
	cmd := exec.Command(list[0], list[1:]...)

	err := cmd.Run()
	if err != nil {
		if s.svnName != "licenseManager" {
			log.Fatal("[ERROR] update svn err. ", s.svnUrl, list, err)
		} else {
			log.Println("[INFO] 【" + s.svnName + "】 更新成功.")
		}
	} else {
		log.Println("[INFO] 【" + s.svnName + "】 更新成功.")
	}

	cmd.Wait()
	log.Println("[INFO] end update 【" + s.svnName + "】for svn.\n")
}
