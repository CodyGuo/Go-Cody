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
	line        = 10    // '\n'
	enter       = 13    //'\r'
	space       = 32    // ' '
	commaEN     = 44    // ','
	commaZH     = 65292 // '，'
	semicolonEN = 59    // ';'
	semicolonZH = 65307 // '；'

	width  = 450
	height = 380
)

const (
	ipRegxp = `^((25[0-5]|2[0-4]\d|1\d{2}|\d?\d)\.){3}(25[0-5]|2[0-4]\d|1\d{2}|\d?\d)$`
)

const (
	title            = "iMan升级工具 1.0"
	uploadGroupTitle = "升级包上传"
	uploadFileTitle  = "升级包:"
	browserTitle     = "浏览"
	sipTitle         = "服务器IP:"
	uploadTitle      = "上传"
	logTitle         = "日志"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func regex(s string) bool {
	re, err := regexp.Compile(ipRegxp)
	checkError(err)
	return re.MatchString(s)
}

func main() {
	mw := new(MyWindow)

	mw.RunApp()
}

type MyWindow struct {
	*walk.MainWindow
	notifyIcon   *walk.NotifyIcon
	uploadGroup  *walk.GroupBox
	uploadFileLb *walk.Label
	uploadFileLe *walk.LineEdit
	browserBtn   *walk.PushButton
	sipLb        *walk.Label
	sipTe        *walk.TextEdit
	uploadBtn    *walk.PushButton
	logLb        *walk.Label
	lv           *LogView

	file string
}

func (mw *MyWindow) RunApp() {
	mw.SetMaximizeBox(false)
	mw.SetFixedSize(true)
	font := Font{Family: "幼圆", PointSize: 10, Bold: true}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    title,
		MinSize:  Size{width, height},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout:  HBox{MarginsZero: true, Spacing: 5},
				MaxSize: Size{0, 180},
				Children: []Widget{
					GroupBox{
						Title:  uploadGroupTitle,
						Font:   font,
						Layout: Grid{Columns: 3},
						Children: []Widget{
							Label{AssignTo: &mw.uploadFileLb, Text: uploadFileTitle},
							LineEdit{AssignTo: &mw.uploadFileLe, ReadOnly: true},
							PushButton{
								AssignTo: &mw.browserBtn,
								Text:     browserTitle,
								OnClicked: func() {
									go mw.setUploadFile()
								},
							},

							Label{AssignTo: &mw.sipLb, Text: sipTitle},
							TextEdit{
								AssignTo:    &mw.sipTe,
								ToolTipText: "请输入iMan服务器IP地址，可输入多个，以分号、空格、逗号或换行隔开。",
								RowSpan:     5,
							},
							PushButton{
								AssignTo: &mw.uploadBtn,
								Text:     uploadTitle,
								OnClicked: func() {
									go mw.upload()
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: VBox{Margins: Margins{2, 2, width - 60, 2}},
				Children: []Widget{
					Label{AssignTo: &mw.logLb, Font: font, Text: logTitle},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	mw.setHeight(25)

	mw.lv, _ = NewLogView(mw)
	log.SetOutput(mw.lv)

	mw.addNotyfyAction()
	mw.setIcon(3)

	mw.Run()
}

// 设置高度
func (mw *MyWindow) setHeight(value int) {
	mw.uploadFileLb.SetHeight(value)
	mw.uploadFileLe.SetHeight(value)
	mw.browserBtn.SetHeight(value)

	mw.sipLb.SetHeight(value)
	// mw.sipLe.SetHeight(value)
	mw.uploadBtn.SetHeight(value)
}

// 托盘图标
func (mw *MyWindow) addNotyfyAction() {
	var err error
	mw.notifyIcon, err = walk.NewNotifyIcon()
	checkError(err)
	mw.notifyIcon.SetVisible(true)
	exitAction := walk.NewAction()
	exitAction.SetText("退出程序")
	exitAction.Triggered().Attach(func() {
		mw.Dispose()
		mw.notifyIcon.Dispose()
		walk.App().Exit(0)
	})
	mw.notifyIcon.ContextMenu().Actions().Add(exitAction)
}

// 设置程序图标
func (mw *MyWindow) setIcon(value int) {
	icon, err := walk.NewIconFromResourceId(uintptr(value))
	checkError(err)
	mw.SetIcon(icon)
	mw.notifyIcon.SetIcon(icon)
}

// 浏览升级包
func (mw *MyWindow) browser() {
	fd := new(walk.FileDialog)
	fd.Title = "选择iMan升级包"
	fd.Filter = "iMan Files 最大200MB|*"
	fd.FilePath = mw.file

	if _, err := fd.ShowOpen(mw); err != nil {
		log.Fatal(err)
	}

	mw.file = fd.FilePath
}

// 浏览
func (mw *MyWindow) setUploadFile() {
	mw.browser()
	mw.uploadFileLe.SetText(mw.file)
	mw.uploadFileLe.SetTextSelection(len(mw.file), len(mw.file))
}

// 上传
func (mw *MyWindow) upload() {
	mw.lv.Clean()
	okSIP, errSIP := mw.checkSIP()
	switch {
	case mw.file == "":
		mw.msg("DEBUG", "请选择升级包！")
	case len(okSIP) == 0:
		mw.msg("DEBUG", fmt.Sprintf("请输入正确的服务器IP: %s\n", errSIP))
	default:
		mw.msg("INFO", fmt.Sprintf("正确的服务器IP: %s\n错误的服务器IP：%s\n", okSIP, errSIP))
	}
}

// 转换为list
func (mw *MyWindow) stringToList() []string {
	sip := mw.sipTe.Text()

	m := func(c rune) rune {
		switch c {
		case line, enter:
			return commaEN
		}
		return c
	}
	sip = strings.Map(m, sip)

	f := func(c rune) bool {
		return c == space || c == commaEN || c == commaZH || c == semicolonEN || c == semicolonZH
	}
	return strings.FieldsFunc(sip, f)
}

// IP地址验证
func (mw *MyWindow) checkSIP() ([]string, []string) {
	var okSIP, errSIP []string

	sipList := mw.stringToList()
	if len(sipList) != 0 {
		for _, sip := range sipList {
			if regex(sip) {
				okSIP = append(okSIP, sip)
			} else {
				errSIP = append(errSIP, sip)
			}
		}
		return okSIP, errSIP
	}
	return nil, nil
}

func (mw *MyWindow) msg(level string, message string) {
	switch level {
	case "INFO":
		log.SetPrefix("[INFO] ")
		title := "提示信息"
		log.Printf("%s\n", message)
		mw.notifyIcon.ShowInfo(title, message)
		walk.MsgBox(mw, title, message, walk.MsgBoxIconInformation)
	case "DEBUG":
		log.SetPrefix("[DEBUG] ")
		title := "警告信息"
		log.Printf("%s\n", message)
		mw.notifyIcon.ShowWarning(title, message)
		walk.MsgBox(mw, title, message, walk.MsgBoxIconWarning)
	}
}
