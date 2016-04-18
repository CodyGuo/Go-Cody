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
file = mac.csv
[Boom]
n = 50
c = 50
debug = false
`
)

type BoomDevce struct {
	URL   []string
	n     string
	c     string
	debug bool
}

func (b *BoomDevce) read() {
	file := "conf.ini"
	_, err := os.Stat(file)
	if err != nil {
		conf, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
		checkErr(err)
		defer conf.Close()
		conf.WriteString(iniFile)
	}

	conf, err := goconfig.LoadConfigFile("conf.ini")
	checkErr(err)
	url, err := conf.GetValue("URL", "url")
	checkErr(err)
	toDo, err := conf.GetValue("URL", "toDo")
	asmType, err := conf.GetValue("URL", "asmType")
	checkErr(err)
	natType, err := conf.GetValue("URL", "natType")
	checkErr(err)
	toUrl, err := conf.GetValue("URL", "toUrl")
	checkErr(err)
	eth0_mac, err := conf.GetValue("URL", "eth0_mac")
	checkErr(err)
	netapp, err := conf.GetValue("URL", "netapp")
	checkErr(err)
	userAgent, err := conf.GetValue("URL", "userAgent")
	checkErr(err)
	netApp_error, err := conf.GetValue("URL", "netApp_error")
	checkErr(err)
	// 获取并发配置
	b.n, err = conf.GetValue("Boom", "n")
	checkErr(err)
	b.c, err = conf.GetValue("Boom", "c")
	checkErr(err)
	debug, err := conf.GetValue("Boom", "debug")
	checkErr(err)
	if debug == "true" {
		b.debug = true
	} else {
		b.debug = false
	}

	csvName, err := conf.GetValue("File", "file")
	checkErr(err)

	csvFile, err := os.Open(csvName)
	checkErr(err)
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
			url, dev[0], dev[1], toDo, asmType,
			natType, toUrl, eth0_mac, netapp, userAgent,
			netApp_error)
		b.URL = append(b.URL, urlStr)
	}

}

func (b *BoomDevce) Request(index int, url string, ok chan bool) {
	writeLog(fmt.Sprintf("[%d] 正在并发[%s]次请求%s\n", index, b.c, url))
	cmd := exec.Command("boom.exe", "-n", b.n, "-c", b.c, url)
	outPut, err := cmd.Output()
	checkErr(err)
	cmd.Wait()
	if b.debug {
		writeLog(fmt.Sprintf("[%d] 请求URL=%s，并发[%s]次结果为：\n%s\n", index, url, b.c, string(outPut)))
	}
	ok <- true

}

func main() {
	checkBoom()
	now := time.Now()
	boom := new(BoomDevce)
	boom.read()
	ok := make(chan bool)
	for index, url := range boom.URL {
		go boom.Request(index, url, ok)
	}

	for range boom.URL {
		<-ok
	}

	tag := "-------------------"
	writeLog(fmt.Sprintf("[结果]\n%s并发设备数[%d]个，并发请求次数[%s]次。%s\n", tag, len(boom.URL), boom.c, tag))
	writeLog(fmt.Sprintf("%s本次并发测试结束，用时%s。%s\n", tag, fmt.Sprint(time.Since(now)), tag))
}

func writeLog(info string) {
	fmt.Println(info)
	logFile, err := os.OpenFile("boom_iman.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	checkErr(err)
	defer logFile.Close()
	logFile.WriteString(info)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
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
