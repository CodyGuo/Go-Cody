package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	LAYOUT = "2006-01-02"
)

var (
	svn *Svn
)

var (
	linuxBak      *Backup
	iManBak       *Backup
	pcHelper      *Backup
	androidHelper *Backup
	nacupgrade    *Backup
	register      *Backup
	autoTesting   *Backup
	webSite       *Backup
	erp           *Backup
	erpApp        *Backup

	bbs         *Backup
	product     *Backup
	productFile *Backup
)

type Backup struct {
	fileName string
	cmd      string
	Svn
}

func (b *Backup) bakCode(ok chan bool) {
	if !isDirExists(b.pathName) {
		b.Get()
	} else {
		b.Update()
	}

	ok <- true
}

func (b *Backup) bakFile(ok chan bool) {
	if b.cmd != "" {
		sshConn(b)
		tarFile(b)
	}

	ok <- true
}

func main() {
	timeNow := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("---------开始备份，大概用时20分钟，请耐心等待.-----------\n")
	ok := make(chan bool)
	// // linux bak
	go linuxBak.bakCode(ok)
	<-ok

	// // iMan bak
	go iManBak.bakCode(ok)
	<-ok
	go iManBak.bakFile(ok)
	<-ok

	// // PC助手
	go pcHelper.bakCode(ok)
	<-ok

	// // android助手
	go androidHelper.bakCode(ok)
	<-ok

	// // 升级程序代码
	go nacupgrade.bakCode(ok)
	<-ok

	// // 注册服务器
	go register.bakCode(ok)
	<-ok
	go register.bakFile(ok)
	<-ok

	// // 自动化测试
	go autoTesting.bakCode(ok)
	<-ok

	// // 新官网
	go webSite.bakCode(ok)
	<-ok
	go webSite.bakFile(ok)
	<-ok

	// // 商机系统Web代码
	go erp.bakCode(ok)
	<-ok
	go erp.bakFile(ok)
	<-ok

	// 商机系统App代码
	go erpApp.bakCode(ok)
	<-ok

	// // Bbs备份
	go bbs.bakFile(ok)
	<-ok

	// // 禅道备份
	go product.bakFile(ok)
	<-ok
	go productFile.bakFile(ok)
	<-ok

	sumTime := time.Since(timeNow)
	log.Println("------------备份结束，用时" + fmt.Sprint(sumTime) + "---------------")

	var tmp string
	fmt.Println("请手动关闭窗口.")
	fmt.Scan(&tmp)

}

func MkBakDir(name string) string {
	err := os.MkdirAll(name, 0755)
	checkError(err)

	return name
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
		var errStop string
		fmt.Scan(&errStop)
	}
}

func isDirExists(name string) bool {
	file, err := os.Stat(name)
	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}

	panic(name + " not exist.")
}
