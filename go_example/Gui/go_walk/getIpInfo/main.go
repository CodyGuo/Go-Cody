package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	ip      *walk.LineEdit
	country *walk.LineEdit
	area    *walk.LineEdit
	region  *walk.LineEdit
	city    *walk.LineEdit
	isp     *walk.LineEdit

	query *walk.PushButton
}

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
	mw := new(MyWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "IP查询",
		MinSize:  Size{350, 300},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					Label{Text: "IP: "},
					LineEdit{AssignTo: &mw.ip},
					PushButton{
						AssignTo: &mw.query,
						Text:     "查询",
					},
				},
			},
			Composite{
				MinSize: Size{0, 100},
				Layout:  HBox{},
				Children: []Widget{
					GroupBox{
						Title:  "查询结果",
						Layout: Grid{Columns: 2},
						Children: []Widget{
							Label{Text: "国家"},
							LineEdit{AssignTo: &mw.country, ReadOnly: true},
							Label{Text: "地区"},
							LineEdit{AssignTo: &mw.area, ReadOnly: true},
							Label{Text: "省"},
							LineEdit{AssignTo: &mw.region, ReadOnly: true},
							Label{Text: "市"},
							LineEdit{AssignTo: &mw.city, ReadOnly: true},
							Label{Text: "运营商"},
							LineEdit{AssignTo: &mw.isp, ReadOnly: true},
						},
					},
				},
			},
		},
	}).Create(); err != nil {
		log.Fatalln(err)
	}

	mw.query.Clicked().Attach(func() {
		go func() {
			mw.query.SetText("查询中...")
			mw.query.SetEnabled(false)
			mw.GetIpInfo()
			mw.query.SetText("查询")
			mw.query.SetEnabled(true)

		}()
	})

	mw.Run()
}

func (mw *MyWindow) GetIpInfo() {
	mw.clearInfo()
	ip := net.ParseIP(mw.ip.Text())
	if ip == nil {
		walk.MsgBox(mw, "查询地址", "您输入的不是有效的IP地址，请重新输入！", walk.MsgBoxIconWarning)
		return
	}

	ipResult := tabaoAPI(ip.String())

	mw.country.SetText(ipResult.Data.Country)
	mw.area.SetText(ipResult.Data.Area)
	mw.region.SetText(ipResult.Data.Region)
	mw.city.SetText(ipResult.Data.City)
	mw.isp.SetText(ipResult.Data.Isp)

	walk.MsgBox(mw, "查询地址", "查询结束!", walk.MsgBoxIconInformation)
}

func (mw *MyWindow) clearInfo() {
	mw.country.SetText("")
	mw.area.SetText("")
	mw.region.SetText("")
	mw.city.SetText("")
	mw.isp.SetText("")
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
