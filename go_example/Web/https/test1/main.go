package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CodyGuo/logs"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://10.10.3.227/GetVersionServlet?type=1")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	logs.Noticef("StatusCode: %d\n", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		fmt.Println(string(body))
	}
}
