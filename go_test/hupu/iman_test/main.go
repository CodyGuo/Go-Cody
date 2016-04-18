package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

import (
	"github.com/Unknwon/goconfig"
)

const (
	iniFile = `[URL]
url = http://10.10.3.227/index.jsp
toDo = 1
asmType = 1
natType = 0
toUrl = union.click.jd.com
eth0_mac = b0-51-8e-03-9c-93
netapp = 1
userAgent = 2
netApp_error = 0
[File]
csvFile = mac.csv
# c为并发数，如果mysql为true优先使用服务器上的设备.
[Boom]
c = 50
mysql = true
debug = false
`
)

type Conf struct {
	url          string
	toDo         string
	asmType      string
	natType      string
	toUrl        string
	eth0_mac     string
	netapp       string
	userAgent    string
	netApp_error string
	csvFile      string
}

type BoomDevce struct {
	URL   []string
	conf  Conf
	c     string
	mysql bool
	debug bool
}

func (b *BoomDevce) readINI() {
	file := "conf.ini"
	info := "readINI: "
	// 初始化配置文件
	_, err := os.Stat(file)
	if err != nil {
		conf, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
		checkErr(info, err)
		defer conf.Close()
		conf.WriteString(iniFile)
	}

	// 读取配置文件
	conf, err := goconfig.LoadConfigFile(file)
	checkErr(info, err)
	b.conf.url, err = conf.GetValue("URL", "url")
	checkErr(info, err)
	b.conf.toDo, err = conf.GetValue("URL", "toDo")
	checkErr(info, err)
	b.conf.asmType, err = conf.GetValue("URL", "asmType")
	checkErr(info, err)
	b.conf.natType, err = conf.GetValue("URL", "natType")
	checkErr(info, err)
	b.conf.toUrl, err = conf.GetValue("URL", "toUrl")
	checkErr(info, err)
	b.conf.eth0_mac, err = conf.GetValue("URL", "eth0_mac")
	checkErr(info, err)
	b.conf.netapp, err = conf.GetValue("URL", "netapp")
	checkErr(info, err)
	b.conf.userAgent, err = conf.GetValue("URL", "userAgent")
	checkErr(info, err)
	b.conf.netApp_error, err = conf.GetValue("URL", "netApp_error")
	checkErr(info, err)
	// 获取并发配置
	b.c, err = conf.GetValue("Boom", "c")
	checkErr(info, err)
	mysql, err := conf.GetValue("Boom", "mysql")
	checkErr(info, err)
	if mysql == "true" {
		b.mysql = true
	} else {
		b.mysql = false
	}
	debug, err := conf.GetValue("Boom", "debug")
	checkErr(info, err)
	if debug == "true" {
		b.debug = true
	} else {
		b.debug = false
	}

	if b.mysql != true {
		b.conf.csvFile, err = conf.GetValue("File", "csvFile")
		checkErr(info, err)
	}
}

func (b *BoomDevce) getDeviceFromCSV() {
	info := "getDeviceFromCSV："
	csvFile, err := os.Open(b.conf.csvFile)
	checkErr(info, err)
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for {
		dev, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		urlStr := fmt.Sprintf("%s?ip=%s&mac=%s&toDo=%s&asmType=%s&natType=%s&toUrl=%s&eth0_mac=%s&netapp=%s&userAgent=%s&netApp_error=%s",
			b.conf.url, dev[0], dev[1], b.conf.toDo, b.conf.asmType,
			b.conf.natType, b.conf.toUrl, b.conf.eth0_mac, b.conf.netapp,
			b.conf.userAgent, b.conf.netApp_error)
		b.URL = append(b.URL, urlStr)
	}

}

func (b *BoomDevce) getDeviceFromMysql() {
	info := "getDeviceFromMysql"
	fmt.Println(info, "正在获取mysql设备.")
}

func (b *BoomDevce) Request(index int, url string, ok chan bool) {
	info := "Request："
	writeLog(fmt.Sprintf("[%d] 正在并发[%s]次请求%s\n", index, b.c, url))
	cmd := exec.Command("boom", "-n", b.c, "-c", b.c, url)
	outPut, err := cmd.Output()
	checkErr(info, err)
	cmd.Wait()

	if b.debug {
		writeLog(fmt.Sprintf("[%d] 请求URL=%s，并发[%s]次结果为：\n%s\n", index, url, b.c, string(outPut)))
	}
	ok <- true

}

func main() {
	// 检查boom程序是否存在.
	checkBoom()
	now := time.Now()

	boom := new(BoomDevce)
	tag := "-------------------"
	writeLog(fmt.Sprintf("%s[开始]%[1]s\n", tag))

	// 读取配置文件
	boom.readINI()
	if boom.mysql == true {
		boom.getDeviceFromMysql()
	} else {
		boom.getDeviceFromCSV()
	}

	ok := make(chan bool)
	for index, url := range boom.URL {
		go boom.Request(index, url, ok)
	}

	for range boom.URL {
		<-ok
	}

	writeLog(fmt.Sprintf("%s[结束]%[1]s\n%[1]s并发设备数[%d]个，并发请求次数[%s]次。%[1]s\n", tag, len(boom.URL), boom.c))
	writeLog(fmt.Sprintf("%s本次并发测试结束，用时%s。%[1]s\n", tag, fmt.Sprint(time.Since(now))))
}

func writeLog(info string) {
	infoPre := "writeLog："
	fmt.Println(info)
	logFile, err := os.OpenFile("boom_iman.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	checkErr(infoPre, err)
	defer logFile.Close()
	logFile.WriteString(info)
}

func checkErr(info string, err error) {
	if err != nil {
		log.Fatal(info, err)
	}
}

func checkBoom() {
	err := exec.Command("boom").Run()
	if strings.Contains(err.Error(), "exit status 1") {
		return
	}
	if err != nil || strings.Contains(err.Error(), "executable file not found in %PATH%") {
		fmt.Println("并发测试程序已退出,原因：boom.exe程序不存在.\n解决办法：请把boom.exe程序放到当前目录或者放置系统环境变量中.")
		os.Exit(1)
	}
}
