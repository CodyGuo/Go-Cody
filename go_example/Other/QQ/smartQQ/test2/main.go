package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url_ := "https://ssl.ptlogin2.qq.com/ptqrlogin?ptqrtoken=535989365&webqq_type=10&remember_uin=1&login2qq=1&aid=501004106&u1=http%3A%2F%2Fw.qq.com%2Fproxy.html%3Flogin2qq%3D1%26webqq_type%3D10&ptredirect=0&ptlang=2052&daid=164&from_ui=1&pttype=1&dumy=&fp=loginerroralert&action=0-0-14699&mibao_css=m_webqq&t=undefined&g=1&js_type=0&js_ver=10197&login_sig=&pt_randsalt=0"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url_, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
	// req.Header.Set("Content-Type", "application/x-javascript; charset=utf-8")

	req.Header.Set("Cookie", "qrsig=XdmyS5bTv5X6OsFgmgtQD6JLULeDF46R1hcZt2D7M4wfXpKInhNjgUGZMBOUk--D")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	rd, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body -> %s\n", rd)
}
