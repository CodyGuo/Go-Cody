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
)

type Proxy struct {
	ip         string
	port       string
	pTtype     string
	speed      string
	verifyTime string
}

type ProxyModel struct {
	proxy []Proxy
	debug bool
	*goquery.Document
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
		checkError(err)
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

func (p *ProxyModel) ReadINI() {
	conf, err := goconfig.LoadConfigFile(proxyFile)
	checkError(err)
	p.proxy = make([]Proxy, 100)
	ipList, err := conf.GetSection("HTTP")
	checkError(err)
	for ip, port := range ipList {
		i := 0
		p.proxy[i].ip = ip
		p.proxy[i].port = port
		i++
	}
}

func (p *ProxyModel) writeINI() {
	_, err := os.Stat(proxyFile)
	if err != nil {
		file, err := os.OpenFile(proxyFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		checkError(err)
		file.Close()
		file.WriteString("[HTTP]\n[HTTPS]\n")
	}

	conf, err := goconfig.LoadConfigFile(proxyFile)
	checkError(err)
	for _, v := range p.proxy {
		switch v.pTtype {
		case "HTTP":
			conf.SetValue("HTTP", v.ip, v.port)
		case "HTTPS":
			conf.SetValue("HTTPS", v.ip, v.port)
		}
	}
	goconfig.SaveConfigFile(conf, proxyFile)
}

func (p *ProxyModel) updateINI() {
	conf, err := goconfig.LoadConfigFile(proxyFile)
	checkError(err)
	proxyIP, err := conf.GetSection("HTTP")
	checkError(err)
	ok := make(chan bool)
	for ip, port := range proxyIP {
		ip := ip
		go func() {
			_, err := net.DialTimeout("tcp", ip+":"+port, 10*time.Second)
			if err != nil {
				fmt.Printf("代理IP[%-15s]不能访问，正在删除.\n", ip)
				conf.DeleteKey("HTTP", ip)
			}
			ok <- true
		}()
	}

	for range proxyIP {
		<-ok
	}
	goconfig.SaveConfigFile(conf, proxyFile)
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

func (p *ProxyModel) getTransportFileURL(addr *string) (transport *http.Transport) {
	URL := url.URL{}
	URLProxy, _ := URL.Parse(*addr)
	transport = &http.Transport{Proxy: http.ProxyURL(URLProxy)}
	return
}

func (p *ProxyModel) Request(url, addr *string) {
	transport := p.getTransportFileURL(addr)
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", *url, nil)
	checkError(err)

	resp, err := client.Do(req)
	checkError(err)
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("代理[%s] 访问url[%s]成功.\n", *addr, *url)
	} else {
		fmt.Printf("代理[%s] 访问url[%s]失败.\n", *addr, *url)
	}
}

func main() {
	now := time.Now()
	url := "http://www.xicidaili.com/nn"
	proxy := GetProxy(url, 10)
	proxy.writeINI()
	proxy.updateINI()
	fmt.Printf("获取100页代理用时%v\n", fmt.Sprint(time.Since(now)))

	// proxy.show()

	proxy.ReadINI()
	urlQuest := "http://studygolang.com/articles/6551"
	ok := make(chan bool)
	for _, pro := range proxy.proxy {
		addr := fmt.Sprintf("http://%s:%s", pro.ip, pro.port)
		go func() {
			proxy.Request(&urlQuest, &addr)
			fmt.Printf("使用代理[%s]访问URL[%s]完成.", addr, urlQuest)
			ok <- true
		}()
	}

	for range proxy.proxy {
		<-ok
	}

	var tmp string
	fmt.Scan(&tmp)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
