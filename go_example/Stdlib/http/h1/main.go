package main

import (
    "fmt"
    "net/http"
    "runtime"
    "time"
)

type respChan struct {
    url      string
    response *http.Response
    reptime  float32
}

var cur int
var worker = runtime.NumCPU()
var ch = make(chan *respChan, 1000)

func gethtml(url string) {
    timeout := time.Duration(10 * time.Second)
    client := http.Client{
        Timeout: timeout,
    }
    startTime := time.Now().UnixNano()
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()
    if true == resp.Close {
        fmt.Println("我被关闭了...")
        resp.Body.Close()
    }

    processed := float32(time.Now().UnixNano()-startTime) / 1e6 // 统计总耗时

    ch <- &respChan{url, resp, processed}

}

func main() {

    urllist := make([]string, 1000)

    for i := range urllist {
        urllist[i] = "http://www.baidu.com"
    }

    fmt.Print(len(urllist))
    runtime.GOMAXPROCS(runtime.NumCPU())

    for _, url := range urllist {
        go gethtml(url)
    }

    for i := 0; i < len(urllist); i++ {
        result := <-ch
        fmt.Printf("%s status: %d %.3fms\n", result.url, result.response.StatusCode, result.reptime)
    }

}
