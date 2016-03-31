package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Unknwon/goconfig"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"
	"time"
)

const (
	proxyFile = "proxy.ini"
	HTTP      = "HTTP"
	HTTPS     = "HTTPS"
)

type Proxy struct {
	ip         string
	port       string
	pTtype     string
	speed      string
	verifyTime string
}

type ProxyConf struct {
	*goconfig.ConfigFile
	file  string
	proxy []Proxy
}

func (p *ProxyConf) load() {
	var err error
	_, err = os.Stat(p.file)
	if err != nil {
		file, err := os.OpenFile(p.file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
		checkError("load create ini", err)
		file.WriteString(fmt.Sprintf("[%s]\n[%s]\n", HTTP, HTTPS))
		defer file.Close()

	}
	p.ConfigFile, err = goconfig.LoadConfigFile(p.file)
	checkError("load ini file", err)
}

func (p *ProxyConf) Read() {
	p.load()
	http, err := p.GetSection(HTTP)
	checkError("read ini http", err)
	p.proxy = []Proxy{}
	for ip, port := range http {
		p.proxy = append(p.proxy, Proxy{ip: ip, port: port, pTtype: HTTP})
	}
	https, err := p.GetSection(HTTPS)
	checkError("read ini https", err)
	for ip, port := range https {
		p.proxy = append(p.proxy, Proxy{ip: ip, port: port, pTtype: HTTPS})
	}
}

func (p *ProxyConf) Write(proxy []Proxy) {
	if len(proxy) == 0 {
		return
	}
	p.load()
	for _, pro := range proxy {
		switch pro.pTtype {
		case HTTP:
			p.SetValue(HTTP, pro.ip, pro.port)
		case HTTPS:
			p.SetValue(HTTPS, pro.ip, pro.port)
		}
	}
	goconfig.SaveConfigFile(p.ConfigFile, proxyFile)
}

func (p *ProxyConf) Update() {
	p.Read()
	ok := make(chan bool)
	for _, proxy := range p.proxy {
		ip := proxy.ip
		prot := proxy.port
		go func() {
			_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ip, prot), 10*time.Second)
			if err != nil {
				fmt.Printf("代理服务器 [%s:%s] 不可用，已删除.\n", ip, prot)
				p.DeleteKey(HTTP, ip)
				p.DeleteKey(HTTPS, ip)
			}
			ok <- true
		}()
	}
	for range p.proxy {
		<-ok
	}
	goconfig.SaveConfigFile(p.ConfigFile, proxyFile)
	p.Reload()
}

type ProxyModel struct {
	*goquery.Document
	proxy []Proxy
	ProxyConf
}

func GetProxy(url string, page int) *ProxyModel {
	urlList := []string{}
	if page > 0 {
		for i := 1; i <= page; i++ {
			urlList = append(urlList, fmt.Sprintf("%s/%d", url, i))
		}
	}

	proxy := new(ProxyModel)
	ok := make(chan bool)
	for index, url := range urlList {
		fmt.Printf("正在获取代理地址的URL[%d]: %s\n", index, url)
		var err error
		proxy.Document, err = goquery.NewDocument(url)
		checkError("newdocument", err)
		go func() {
			proxy.getProxy()
			ok <- true
		}()
	}

	for _, url := range urlList {
		<-ok
		urlRec := url
		fmt.Printf("URL[%-30s] 获取代理完成.\n", urlRec)
	}

	return proxy
}

func (p *ProxyModel) show() {
	title := new(Proxy)
	p.Find("#ip_list > tbody > tr").Each(func(i int, t *goquery.Selection) {
		if i == 0 {
			title.ip = t.Find("th").Eq(2).Text()
			title.port = t.Find("th").Eq(3).Text()
			title.pTtype = t.Find("th").Eq(6).Text()
			title.speed = t.Find("th").Eq(7).Text()
			title.verifyTime = t.Find("th").Eq(9).Text()
		}
	})

	const format = "%v\t %v\t %v\t %v\t %v\t \n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 10, 3, ' ', 0)
	fmt.Fprintf(tw, format, title.ip, title.port, title.pTtype, title.speed, title.verifyTime)
	fmt.Fprintf(tw, format, "--------", "--------", "--------", "--------", "--------")

	for _, v := range p.proxy {
		fmt.Fprintf(tw, format, v.ip, v.port, v.pTtype, v.speed, fmt.Sprint(v.verifyTime))
	}

	fmt.Fprintf(tw, "已获取的代理地址数：%d个\n", len(p.proxy))
	tw.Flush()
}

func (p *ProxyModel) getProxy() {
	p.Find("#ip_list > tbody > tr").Each(func(i int, t *goquery.Selection) {
		if i > 0 {
			speed, _ := t.Find("td > div").Attr("title")
			p.proxy = append(p.proxy, Proxy{
				t.Find("td").Eq(2).Text(),
				t.Find("td").Eq(3).Text(),
				t.Find("td").Eq(6).Text(),
				speed,
				t.Find("td").Eq(9).Text(),
			})
		}
	})

}

type TransportGet struct {
	url  string
	addr string
}

func (t TransportGet) GetByProxy() {
	req, err := http.NewRequest("GET", t.url, nil)
	checkError("NewRequest", err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 5.1; it; rv:1.8.1.11) Gecko/20071127 Firefox/2.0.0.11")

	proxyURL, err := url.Parse(t.addr)
	checkError("Parse proxy", err)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		// log.Printf("Request do: %v\n", err)
		fmt.Printf("代理服务器[%s] 访问url[%s]失败.\n", t.addr, t.url)
		return
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("代理服务器[%s] 访问url[%s]成功.\n", t.addr, t.url)
	} else {
		fmt.Printf("代理服务器[%s] 访问url[%s]失败.\n", t.addr, t.url)
	}

	defer resp.Body.Close()
	time.Sleep(10 * time.Second)
}

func main() {
	now := time.Now()
	// url := "http://www.xicidaili.com/nn"
	// proxy := GetProxy(url, 10)

	conf := ProxyConf{
		file: proxyFile,
	}
	// conf.Write(proxy.proxy)
	conf.Read()
	conf.Update()

	ok := make(chan bool)
	for _, proxy := range conf.proxy {
		transportGet := TransportGet{
			url: "http://www.golangtc.com/t/56fc7ab1b09ecc66b90001ea",
		}
		ip := proxy.ip
		port := proxy.port
		var addr string
		switch proxy.pTtype {
		case HTTP:
			addr = fmt.Sprintf("http://%s:%s", ip, port)
		case HTTPS:
			addr = fmt.Sprintf("https://%s:%s", ip, port)
		}
		go func() {
			transportGet.addr = addr
			fmt.Printf("正在使用代理[%s] 请求.\n", addr)
			transportGet.GetByProxy()
			ok <- true
		}()

	}

	for range conf.proxy {
		<-ok
	}
	fmt.Printf("获取100页代理用时%v\n", fmt.Sprint(time.Since(now)))

	var tmp string
	fmt.Scan(&tmp)
}

func checkError(info string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", info, err)
	}
}
