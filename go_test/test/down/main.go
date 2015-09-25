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

    urlMap := strings.Split(*url, "/")
    fileName := urlMap[len(urlMap)-1]

    if strings.Contains(fileName, "=") {
        splitName := strings.Split(fileName, "=")
        fileName = splitName[len(splitName)-1]
    }

    resp, err := http.Get(*url)
    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode == http.StatusOK {
        log.Printf("[INFO] 正在下载: [%s]", fileName)
        fmt.Print("\n")

        downFile, err := os.Create(fileName)
        if err != nil {
            log.Fatal(err)
        }
        i, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
        sourceSiz := int64(i)
        source := resp.Body

        bar := pb.New(int(sourceSiz)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
        bar.Start()

        bar.ShowSpeed = true
        bar.ShowFinalTime = true
        bar.SetMaxWidth(80)

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
