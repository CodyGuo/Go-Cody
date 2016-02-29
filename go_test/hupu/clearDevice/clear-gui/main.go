package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mw := new(MyWindow)

	mw.RunApp()

}

type MyWindow struct {
	*walk.MainWindow
	serverIP *walk.LineEdit
	ip       *walk.TextEdit
	mac      *walk.TextEdit
}

func (mw *MyWindow) RunApp() {
	mw.SetMaximizeBox(false)
	mw.SetFixedSize(true)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Layout:   VBox{},
		MinSize:  Size{450, 380},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: "服务器IP:"},
					LineEdit{AssignTo: &mw.serverIP},

					Label{Text: "故障设备IP:"},
					TextEdit{AssignTo: &mw.ip},

					Label{Text: "故障设备MAC:"},
					TextEdit{AssignTo: &mw.mac},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					PushButton{
						Text: "开始清理故障设备",
						OnClicked: func() {
							mw.DoClear()
						},
					},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	lv, err := NewLogView(mw)
	checkError(err)
	log.SetOutput(lv)

	mw.SetScreenCenter(true)
	mw.Run()
}

func (mw *MyWindow) DoClear() {
	log.Println("i clicked.", mw.ip.Text(), mw.mac.Text())
	mw.checkSIP()

	log.Println("清理结束.")
}

func (mw *MyWindow) checkSIP() {
	serverIP := mw.serverIP.Text()
	servers := mw.list(serverIP)
	fmt.Println("servers: ", servers)
	for _, ip := range servers {
		if mw.checkResgexp(ipRegxp, ip) {
			log.Printf("[INFO] 服务器IP检查正确. %s\n", ip)
		} else {
			log.Printf("[ERROR] 服务器IP检查错误. %s\n", ip)
		}
	}
}

func (mw *MyWindow) checkIP() {

}

func (mw *MyWindow) checkMAC() {

}

func (mw *MyWindow) list(str string) []string {
	f := func(c rune) bool {
		return string(c) == "；" || string(c) == ";"
	}
	str = strings.Replace(str, "\r\n", ";", -1)
	return strings.FieldsFunc(str, f)
}

func (mw *MyWindow) checkResgexp(expr, s string) bool {
	re := regexp.MustCompile(expr)
	return re.MatchString(s)
}
