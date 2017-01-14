package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/codyguo/walk"

	. "github.com/codyguo/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	lv               *LogView
	url              *walk.LineEdit
	cookies          *walk.TextEdit
	beginUID, endUID *walk.LineEdit
	clearBtn         *walk.PushButton
}

var (
	bbsURL     = "http://bbs.hupu.cn"
	bbsCookies = ""
)

func main() {
	title := "BBS 垃圾账号清理"
	mw := new(MyWindow)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    title,
		MinSize:  Size{700, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			Composite{
				Layout:  VBox{},
				MaxSize: Size{0, 200},
				Children: []Widget{
					Label{Text: "BBS_URL："},
					LineEdit{AssignTo: &mw.url, Text: bbsURL},
					Label{Text: "Cookies："},
					TextEdit{AssignTo: &mw.cookies},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{Text: "开始UID："},
					LineEdit{AssignTo: &mw.beginUID},
					Label{Text: "结束UID："},
					LineEdit{AssignTo: &mw.endUID},
					PushButton{AssignTo: &mw.clearBtn, Text: "开始清理", OnClicked: func() {
						mw.lv.Clean()
						url := mw.url.Text()
						cookies := mw.cookies.Text()
						beg := mw.beginUID.Text()
						end := mw.endUID.Text()
						if len(beg) == 0 || len(end) == 0 ||
							len(url) == 0 || len(cookies) == 0 {
							walk.MsgBox(mw, title, "请输入BBS URL，登录cookies，开始UID和结束UID！", walk.MsgBoxIconError)
						} else {
							bbsURL = url
							bbsCookies = cookies
							go func() {
								start := time.Now()
								mw.clearBtn.SetEnabled(false)
								clearUsers(beg, end)
								mw.clearBtn.SetEnabled(true)
								log.Printf("清理垃圾用户结束！！！，用时 %v\n", time.Since(start))
								begInt, _ := strconv.Atoi(beg)
								mw.beginUID.SetText(fmt.Sprint(begInt - 20))
							}()
						}
					}},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	mw.beginUID.TextChanged().Attach(func() {
		beg, _ := strconv.Atoi(mw.beginUID.Text())
		mw.endUID.SetText(fmt.Sprint(beg + 19))
	})

	var err error
	mw.lv, err = NewLogView(mw)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(mw.lv)

	mw.Run()
}

func clearUsers(begStr string, endStr string) {
	beg, _ := strconv.Atoi(begStr)
	end, _ := strconv.Atoi(endStr)
	var wg sync.WaitGroup
	for i := beg; i <= end; i++ {
		wg.Add(1)
		go func(i int) {
			req, err := creatRequest(i)
			if err != nil {
				log.Fatal(err)
			}
			clinet := &http.Client{}
			res, err := clinet.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode == http.StatusOK {
				bodyBytes, _ := ioutil.ReadAll(res.Body)
				if bytes.Contains(bodyBytes, []byte("清除数据成功")) {
					log.Printf("uid: [%d] -> 清除数据成功\n", i)
				} else {
					log.Printf("uid: [%d] -> 清除数失败\n", i)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func creatRequest(uid int) (*http.Request, error) {
	urlStr := bbsURL + "/admin.php?m=u&c=manage&a=doClear"
	tokens := strings.Split(bbsCookies, ";")[0]
	token := strings.Split(tokens, "=")[1]
	data := url.Values{
		"uid":        []string{fmt.Sprint(uid)},
		"clear[]":    []string{"topic", "post", "message"},
		"csrf_token": []string{token},
	}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Cookie":       bbsCookies,
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36",
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

/*
if len(os.Args) < 2 {
	fmt.Printf("Usage: %s minUid maxUid\n", os.Args)
	return
}
min, _ := strconv.Atoi(os.Args[1])
max := min
if len(os.Args) == 3 {
    max, _ = strconv.Atoi(os.Args[2])
}
*/
