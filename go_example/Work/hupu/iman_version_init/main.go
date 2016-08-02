package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	layout  = "20060102"
	version = "升级版本"
	tools   = "升级工具"
	file    = "升级文件"
	tag     = "============"
)

type VersionDir struct {
	Data string
	Dir  Dir
}

func (v *VersionDir) Init() {
	err := v.mkDir(v.Dir.version)
	checkError(err)
	err = v.mkDir(v.Dir.tools)
	checkError(err)
	err = v.mkDir(v.Dir.file)
	checkError(err)
}

func (v *VersionDir) mkDir(path string) error {
	return os.MkdirAll(v.Data+"/"+path, 0777)
}

type Dir struct {
	version string
	tools   string
	file    string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	iterFlag string
)

func init() {
	flag.StringVar(&iterFlag, "iter", "", os.Args[0]+" 3.30-36")
}

func main() {
	flag.Parse()
	log.Printf("%s %s %[1]s", tag, "开始初始化升级文件目录结构.")

	day := time.Now()
	nowDay := day.Format(layout)
	if flag.NFlag() == 0 {
		fmt.Print("请输入文件夹名字，例如3.30-36：")
		fmt.Scan(&iterFlag)
	}

	update := new(VersionDir)
	update.Data = nowDay + " " + iterFlag
	update.Dir.version = version
	update.Dir.tools = tools
	update.Dir.file = file

	update.Init()
	log.Printf("%s %s %[1]s", tag, "升级文件目录结构初始化完成.")
	var pause string
	fmt.Print("请手动关闭窗口.\n")
	fmt.Scan(&pause)
}
