package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	title            = "iMan升级工具 2.0"
	uploadGroupTitle = "升级包上传"
	uploadFileTitle  = "升级包:"
	browserTitle     = "浏览"
	sipTitle         = "服务器IP:"
	uploadTitle      = "上传"
	logTitle         = "日志"
)

const (
	INFO  = "[INFO] "
	DEBUG = "[DEBUG] "
)

func main() {
	mw := new(MyWindow)
	mw.RunApp()

	defer mw.notifyIcon.Dispose()
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
				MaxSize: Size{0, maxWidth},
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
									go mw.SetUploadFile()
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
									go mw.Upload()
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					Label{AssignTo: &mw.logLb, Font: font, Text: logTitle},
					HSpacer{MinSize: Size{80, 0}},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	mw.setHeight(25)
	mw.setWidth(200)

	mw.lv, _ = NewLogView(mw)
	log.SetOutput(mw.lv)

	mw.addNotyfyAction()
	mw.setIcon(3)

	// 解决Ctrl + C之后托盘显示问题
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		s := <-c
		switch s.String() {
		case "interrupt":
			signal.Stop(c)
			fmt.Println("Got sinal top:", s)
			mw.exit()
		default:
			fmt.Println("Got sinal:", s)
		}
	}()

	mw.Run()
}

// 浏览
func (mw *MyWindow) SetUploadFile() {
	mw.browser()
	mw.uploadFileLe.SetText(mw.file)
	mw.uploadFileLe.SetTextSelection(len(mw.file), len(mw.file))
}

// 上传
func (mw *MyWindow) Upload() {
	mw.lv.Clean()
	sip := mw.sipTe.Text()

	okSIP, errSIP := checkSIP(stringToList(sip))
	switch {
	case mw.file == "":
		mw.msg(DEBUG, "请选择升级包！")
	case len(okSIP) == 0:
		mw.msg(DEBUG, fmt.Sprintf("请输入正确的服务器IP: %s\n", errSIP))
	default:
		mw.msg(INFO, fmt.Sprintf("正确的服务器IP: %s\n错误的服务器IP：%s\n", okSIP, errSIP))
	}
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

// 设置宽度
func (mw *MyWindow) setWidth(value int) {
	mw.uploadFileLe.SetWidth(value)
	mw.sipTe.SetWidth(value)
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
		mw.exit()
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

// 退出程序
func (mw *MyWindow) exit() {
	mw.Dispose()
	mw.notifyIcon.Dispose()
	walk.App().Exit(0)
}

func (mw *MyWindow) msg(level string, message string) {
	switch level {
	case INFO:
		log.SetPrefix(INFO)
		title := "提示信息"
		log.Printf("%s\n", message)
		mw.notifyIcon.ShowInfo(title, message)
		walk.MsgBox(mw, title, message, walk.MsgBoxIconInformation)
	case DEBUG:
		log.SetPrefix(DEBUG)
		title := "警告信息"
		log.Printf("%s\n", message)
		mw.notifyIcon.ShowWarning(title, message)
		walk.MsgBox(mw, title, message, walk.MsgBoxIconWarning)
	}
}
