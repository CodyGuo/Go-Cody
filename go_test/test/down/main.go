package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/cheggaaa/pb"
)

var (
    url = flag.String("url", "", "The download file URL.")
)

func Usage() {
    fmt.Printf("Usage of %s:\n  -url=\"http://www.xxx.com/file.exe\": The download file URL.\n", os.Args[0])
    os.Exit(1)
}

func main() {
    flag.Parse()

    if os.Args == nil || *url == "" {
        Usage()
    }

    // 解析/后的文件名字
    urlMap := strings.Split(*url, "/")
    fileName := urlMap[len(urlMap)-1]

    // 解析带? = fileName 的文件名字
    if strings.Contains(fileName, "=") {
        splitName := strings.Split(fileName, "=")
        fileName = splitName[len(splitName)-1]
    }

    resp, err := http.Get(*url)
    if err != nil {
        log.Fatal(err)
    }

    // 判断get url的状态码, StatusOK = 200
    if resp.StatusCode == http.StatusOK {
        log.Printf("[INFO] 正在下载: [%s]", fileName)
        fmt.Print("\n")

        downFile, err := os.Create(fileName)
        if err != nil {
            log.Fatal(err)
        }
        // 不要忘记关闭打开的文件.
        defer downFile.Close()

        // 获取下载文件的大小
        i, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
        sourceSiz := int64(i)
        source := resp.Body

        // 创建一个进度条
        bar := pb.New(int(sourceSiz)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)

        // 显示下载速度
        bar.ShowSpeed = true

        // 显示剩余时间
        bar.ShowTimeLeft = true

        // 显示完成时间
        bar.ShowFinalTime = true

        bar.SetMaxWidth(80)

        bar.Start()

        writer := io.MultiWriter(downFile, bar)
        io.Copy(writer, source)
        bar.Finish()

        fmt.Print("\n")
        log.Printf("[INFO] [%s]下载成功.", fileName)
    } else {
        fmt.Print("\n")
        log.Printf("[ERROR] [%s]下载失败,%s.", fileName, resp.Status)
    }

}
