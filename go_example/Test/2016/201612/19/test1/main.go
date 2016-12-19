package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data`
}

type IP struct {
	Country string `json:"country"`
	Area    string `json:"area"`
	Region  string `json:"region"`
	City    string `json:"city"`
	Isp     string `json:"isp"`
}

func main() {
	ip := tabaoAPI("180.173.144.112")
	fmt.Println(ip)
}

func tabaoAPI(ip string) *IPInfo {
	resp, err := http.Get(fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", ip))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return nil
	}

	return &result
}
