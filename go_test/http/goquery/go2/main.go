package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"text/tabwriter"
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

	var proxy *ProxyModel = &ProxyModel{}
	fmt.Println("url:", urlList)
	ok := make(chan bool)
	for index, url := range urlList {
		fmt.Printf("正在获取代理地址的URL [%d]: %s\n", index, url)
		var err error
		proxy.Document, err = goquery.NewDocument(url)
		checkError(err)
		go func() {
			proxy.getProxy()
			ok <- true
			fmt.Printf("URL[%d] 已发送退出消息.\n", index)
		}()
	}

	for index, _ := range urlList {
		<-ok
		fmt.Printf("URL[%d] 已收到退出消息.\n", index)
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

	fmt.Fprintf(tw, "已获取的代理地址数：%d\n", len(p.proxy))
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

func main() {
	url := "http://www.xicidaili.com/nn"
	proxy := GetProxy(url, 3)
	proxy.show()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
