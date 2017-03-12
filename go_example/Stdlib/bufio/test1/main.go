package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/codyguo/logs"
)

func main() {
	f, err := os.Open("url.txt")
	if err != nil {
		logs.Fatal(err)
	}
	urlList := RetrunBurstURL(f, "http://wwww.")
	logs.Noticef("urlList --> %v", urlList)
}

func RetrunBurstURL(fURL *os.File, baseurl string) (urlList []string) {
	allURLTxt := bufio.NewReader(fURL)
	for {
		urlpath, err := allURLTxt.ReadString('\n')
		if err == io.EOF {
			logs.Info("Reading urlfile done...")
			return urlList
		}

		urlpath = strings.TrimSpace(urlpath)
		urlList = append(urlList, baseurl+urlpath)
	}
}
