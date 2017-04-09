package main

import (
	"fmt"
	"net/http"

	"github.com/codyguo/logs"
)

func main() {
	requestURL := "http://dwz.cn/1E"
	req, errNewRequest := http.NewRequest(http.MethodGet, requestURL, nil)
	if errNewRequest != nil {
		logs.Error(errNewRequest)
		return
	}
	req.Header.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; WOW64; rv:52.0) Gecko/20100101 Firefox/52.0")

	// RoundTrip executes a single HTTP transaction, returning the Response for the request req. (RoundTrip 代表一个http事务，给一个请求返回一个响应)
	// 说白了，就是你给它一个request,它给你一个response
	res, errGetURLContent := http.DefaultTransport.RoundTrip(req)
	if errGetURLContent != nil {
		fmt.Println("AppService AnalysisURL request errGetURLContent:", errGetURLContent.Error())
	}

	fmt.Println("request status code:", res.StatusCode, " ", res.Status)

}
