package main

import (
	"fmt"
	// "os"
	"os/exec"
	"strings"
)

const (
	PREFIX = "[ERROR]"
	INSTAG = "============="
)

func main() {
	var flag bool
	// 验证go的安装
	if out := doCmd("cmd", "/c", "go", "version"); out != "cmdError" {
		v := strings.Split(out, "go")
		fmt.Println("已安装的go版本为:", v[2])
		flag = true
	} else {
		fmt.Println(PREFIX, "未安装go,请先安装go.")
		// return
	}
	// 验证git的安装
	if out := doCmd("cmd", "/c", "git", "version"); out != "cmdError" {
		v := strings.Split(out, " ")
		fmt.Println("已安装的git版本为:", v[2])
		flag = true
	} else {
		fmt.Println(PREFIX, "未安装git,请先安装git.")
		flag = false
	}

	var gitList []string
	gitList = []string{
		"github.com/akavel/rsrc",
		"github.com/nsf/gocode",
		"github.com/codyguo/Go-Cody",
		"github.com/codyguo/xcgui",
		"github.com/lxn/walk",
		"github.com/lxn/win",
		"github.com/golang/crypto",
		"github.com/golang/image",
		"github.com/golang/tools",
		"github.com/golang/text",
		"github.com/golang/net",
		"github.com/golang/exp",
		"github.com/golang/sys",
	}

	if flag {
		for _, git := range gitList {
			fmt.Println(INSTAG, "正在安装", git, INSTAG)
			if out := doCmd("cmd", "/c", "go", "get", "-u", git); out != "cmdError" {
				fmt.Println(git, "安装成功.")
			} else {
				fmt.Println(git, "安装失败.")
			}
			fmt.Print("\n")
		}
		fmt.Println("开发环境安装完成,可以开心的玩耍了.")
	} else {
		fmt.Println("开发环境验证失败,请重新检查是否安装go和git.")
	}

	var s string
	fmt.Scan(&s)
}

func doCmd(arg ...string) string {
	if len(arg) == 1 {
		out, err := exec.Command(arg[0]).Output()
		if err != nil {
			return "cmdError"
		}
		return string(out)
	}

	out, err := exec.Command(arg[0], arg[1:]...).Output()
	if err != nil {
		return "cmdError"
	}
	return string(out)
}
