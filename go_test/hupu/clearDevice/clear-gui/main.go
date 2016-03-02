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
	ipRegxp  = `^((25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)\.){3}(25[0-5]|2[0-4]\d|1\d{2}|[1-9]?\d)$`
	macRegxp = `^([0-9a-zA-Z]{2}-){5}([0-9a-zA-Z]){2}$`
)

const (
	line        = 10    // '\n'
	enter       = 13    //'\r'
	space       = 32    // ' '
	asterisk    = 42    // '*'
	question    = 63    // "?"
	questionBig = 65311 // '？'
	percent     = 37    // '%'
	underline   = 95    // '_'
	commaEN     = 44    // ','
	commaZH     = 65292 // '，'
	semicolonEN = 59    // ';'
	semicolonZH = 65307 // '；'

	maxSize       = 600
	minSize       = 450
	title         = "iMan设备故障清理工具 1.1"
	sipTitle      = "服务器IP:"
	ipTitle       = "服务器IP:"
	macTitle      = "故障设备MAC:"
	clearTitle    = "开始清理故障设备"
	clearingTitle = "正在清理故障设备..."
	sCheckTitle   = "服务器故障检查"
	sipHelp       = "请输入iMan服务器IP地址，可输入多个，以分号、空格、逗号隔开。"
	ipHelp        = "请输入故障设备IP地址，可输入多个，以分号、空格、逗号或换行隔开。"
	macHelp       = "请输入故障设备MAC地址，可输入多个，以分号、空格、逗号或换行隔开。"
)

var (
	ipChan  chan bool
	macChan chan bool

	numZero bool
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
	serverIP     *walk.LineEdit
	ip           *walk.TextEdit
	mac          *walk.TextEdit
	clearButton  *walk.PushButton
	sCheckButton *walk.PushButton
}

// 主界面
func (mw *MyWindow) RunApp() {
	mw.SetMaximizeBox(false)
	mw.SetFixedSize(true)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    title,
		Layout:   VBox{},
		MinSize:  Size{maxSize, minSize},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{Text: sipTitle},
					LineEdit{
						AssignTo:    &mw.serverIP,
						ToolTipText: sipHelp,
					},

					Label{Text: ipTitle},
					TextEdit{
						AssignTo:    &mw.ip,
						ToolTipText: ipHelp,
					},

					Label{Text: macTitle},
					TextEdit{
						AssignTo:    &mw.mac,
						ToolTipText: macHelp,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &mw.clearButton,
						Text:     clearTitle,
						OnClicked: func() {
							go mw.DoClear()
						},
					},
					PushButton{
						AssignTo: &mw.sCheckButton,
						Text:     sCheckTitle,
						OnClicked: func() {
							mw.msg("INFO", "hello 服务器故障检查.")
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

// 执行清理动作
func (mw *MyWindow) DoClear() {
	mw.clearButton.SetText(clearingTitle)
	mw.disable()

	log.Println("[INFO] 开始清理故障设备...")
	ok, sipList := mw.checkSIP()
	okIP, errIP := mw.checkIP()
	okMAC, errMAC := mw.checkMAC()
	switch ok {
	case false:
		numZero = true
		mw.msg("ERROR", fmt.Sprintf("[ERROR] 服务器IP检查错误. %s", sipList))
	case len(okIP) == 0 && len(okMAC) == 0:
		numZero = true
		mw.msg("ERROR", "[ERROR] 请填写正确的设备IP或者MAC.")
	default:
		if len(okIP) != 0 {
			ipChan = make(chan bool)
			go mw.clearIP(sipList, okIP)
			<-ipChan
		}
		if len(okMAC) != 0 {
			macChan = make(chan bool)
			go mw.clearMAC(sipList, okMAC)
			<-macChan
		}

		if len(errIP) != 0 || len(errMAC) != 0 {
			mw.msg("ERROR", fmt.Sprintf("[ERROR] 检查并清理失败的设备IP: %s\n[ERROR] 检查并清理失败的设备MAC: %s", errIP, errMAC))
		}
	}

	if numZero {
		okIP, okMAC = []string{}, []string{}
	}
	mw.msg("INFO", fmt.Sprintf("[INFO] 清理故障设备结束. 服务器%s 清理设备IP [%d] 个, MAC [%d] 个.", sipList, len(okIP), len(okMAC)))
	tag := "====================================="
	log.Printf("%s\n", tag)
	mw.clearButton.SetText(clearTitle)
	mw.enable()

}

// 启用
func (mw *MyWindow) enable() {
	enable := true
	mw.serverIP.SetEnabled(enable)
	mw.ip.SetEnabled(enable)
	mw.mac.SetEnabled(enable)
	mw.clearButton.SetEnabled(enable)
	mw.sCheckButton.SetEnabled(enable)
}

// 禁用
func (mw *MyWindow) disable() {
	disable := false
	mw.serverIP.SetEnabled(disable)
	mw.ip.SetEnabled(disable)
	mw.mac.SetEnabled(disable)
	mw.clearButton.SetEnabled(disable)
	mw.sCheckButton.SetEnabled(disable)
}

// 清理IP
func (mw *MyWindow) clearIP(sipList, okIP []string) {
	for _, sIP := range sipList {
		for _, ip := range okIP {
			err := doSQL(sIP, 0, ip)
			if err != nil {
				numZero = true
				mw.msg("DEBUG", fmt.Sprintf("[DEBUG] 服务器[%s] 连接错误: %s", sIP, err))
			} else {
				log.Printf("[INFO] 服务器[%s] 正在清理的设备IP: [%s].\n", sIP, ip)
			}
		}
	}
	ipChan <- true
}

// 清理MAC
func (mw *MyWindow) clearMAC(sipList, okMAC []string) {
	for _, sIP := range sipList {
		for _, mac := range okMAC {
			err := doSQL(sIP, 1, mac)
			if err != nil {
				numZero = true
				mw.msg("DEBUG", fmt.Sprintf("[DEBUG] 服务器[%s] 连接错误: %s", sIP, err))
			} else {
				log.Printf("[INFO] 服务器[%s] 正在清理的设备MAC: [%s].\n", sIP, mac)
			}

		}
	}
	macChan <- true
}

// 检查服务器IP正确性
func (mw *MyWindow) checkSIP() (bool, []string) {
	serverIP := mw.serverIP.Text()
	sipList := mw.list(serverIP)
	var errSIP []string
	for _, ip := range sipList {
		if !mw.Resgexp(ipRegxp, ip) {
			errSIP = append(errSIP, ip)
		}
	}
	if len(sipList) == 0 || len(errSIP) != 0 {
		return false, errSIP
	}
	return true, sipList
}

// 检查设备IP正确性，并返回正确的IP和错误的IP地址。
func (mw *MyWindow) checkIP() ([]string, []string) {
	ip := mw.ip.Text()
	ipList := mw.list(ip)
	var okIP, errIP []string
	for _, ip := range ipList {
		ok, ip := mw.replaceSQL(ip)
		if mw.Resgexp(ipRegxp, ip) || ok == true {
			okIP = append(okIP, ip)
		} else {
			errIP = append(errIP, ip)
		}
	}
	return okIP, errIP
}

// 检查设备MAC正确性，并返回正确的MAC和错误的MAC地址。
func (mw *MyWindow) checkMAC() ([]string, []string) {
	mac := mw.mac.Text()
	macList := mw.list(mac)
	var okMAC, errMAC []string
	for _, mac := range macList {
		ok, mac := mw.replaceSQL(mac)
		if mw.Resgexp(macRegxp, mac) || ok == true {
			okMAC = append(okMAC, mac)
		} else {
			errMAC = append(errMAC, mac)
		}
	}
	return okMAC, errMAC
}

// 字符串转换为列表
func (mw *MyWindow) list(str string) []string {
	f := func(c rune) bool {
		return c == commaEN || c == commaZH ||
			c == semicolonEN || c == semicolonZH
	}
	m := func(r rune) rune {
		switch r {
		case line, enter, space:
			return commaEN
		}
		return r
	}
	str = strings.Map(m, str)
	return strings.FieldsFunc(str, f)
}

// 数据库特殊字符匹配转换
func (mw *MyWindow) replaceSQL(str string) (bool, string) {
	m := func(r rune) rune {
		switch r {
		case asterisk, percent:
			return percent
		case question, questionBig, underline:
			return underline
		}
		return r
	}
	ok := strings.ContainsAny(str, "* ? % _")
	if ok {
		return ok, strings.Map(m, str)
	}
	return false, str
}

// 正则验证
func (mw *MyWindow) Resgexp(expr, s string) bool {
	re := regexp.MustCompile(expr)
	return re.MatchString(s)
}

func (mw *MyWindow) msg(level, message string) {
	switch level {
	case "INFO":
		walk.MsgBox(mw, "提示信息", message, walk.MsgBoxIconInformation)
		log.Printf("%s\n", message)
	case "DEBUG":
		walk.MsgBox(mw, "警告信息", message, walk.MsgBoxIconWarning)
		log.Printf("%s\n", message)
	case "ERROR":
		walk.MsgBox(mw, "错误信息", message, walk.MsgBoxIconError)
		log.Printf("%s\n", message)
	}
}
