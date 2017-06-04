package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	URL = `http://10.10.2.227/approve/Users-login`
)

var (
	data   = `&idomainId=&scompanycode=10000000&rememberusername=true&rememberpwd=true&asmType=1&sdeviceip=&sdevicemac=&ifnat=0&deviceip=10.10.2.162&devicemac=00-00-00-00-00-00`
	user   string
	passwd string
)
var limit int

func init() {
	flag.IntVar(&limit, "limit", 200, "limit request.")
}

func main() {
	flag.Parse()
	var n int
	var wg sync.WaitGroup
	user = base64.StdEncoding.EncodeToString([]byte("hupu@hupu.net"))
	passwd = base64.StdEncoding.EncodeToString([]byte("1"))
	data = fmt.Sprintf("susername=%s&suserpwd=%s%s", user, passwd, data)

	for {
		for i := 1; i <= limit; i++ {
			n++
			wg.Add(i)
			time.Sleep(200 * time.Millisecond)
			log.Printf("request --> [%d]\n", n)
			go func(n int) {
				res, err := Req(URL)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("response --> [%d], %s\n", n, res)
				wg.Done()
			}(n)
		}
		time.Sleep(10 * time.Second)
	}
}

func Req(url_ string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// 请求账号
	body := strings.NewReader(data)
	req, _ := http.NewRequest(http.MethodPost, url_, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Set("IMan-Language", "zh-CN")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	result, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(result))
	resp.Body.Close()
	return result, nil
}
