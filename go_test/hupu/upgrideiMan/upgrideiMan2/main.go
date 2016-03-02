package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	maxSize = 380
	minSize = 280
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

func main() {
	mw := new(MyWindow)
	mw.RunApp()
}

type MyWindow struct {
	*walk.MainWindow
	uploadGroup  *walk.GroupBox
	uploadFileLb *walk.Label
	uploadFileLe *walk.LineEdit
	browserBtn   *walk.PushButton
	sipLb        *walk.Label
	sipLe        *walk.LineEdit
	uploadBtn    *walk.PushButton
}

func (mw *MyWindow) RunApp() {
	mw.SetMaximizeBox(false)
	// mw.SetFixedSize(true)
	font := Font{Family: "幼圆", PointSize: 10, Bold: true}

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    title,
		MinSize:  Size{maxSize, minSize},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					GroupBox{
						Title:         uploadGroupTitle,
						Font:          font,
						StretchFactor: 20,
						Layout:        Grid{Columns: 3},
						Children: []Widget{
							Label{AssignTo: &mw.uploadFileLb, Text: uploadFileTitle},
							LineEdit{AssignTo: &mw.uploadFileLe, ReadOnly: true},
							PushButton{AssignTo: &mw.browserBtn, Text: browserTitle},

							VSpacer{RowSpan: 3, ColumnSpan: 3},

							Label{AssignTo: &mw.sipLb, Text: sipTitle},
							LineEdit{AssignTo: &mw.sipLe, CueBanner: "请输入iMan服务器IP."},
							PushButton{AssignTo: &mw.uploadBtn, Text: uploadTitle},
						},
					},
				},
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}
	mw.setHeight(25)

	mw.Run()
}

func (mw *MyWindow) setHeight(value int) {
	mw.uploadFileLb.SetHeight(value)
	mw.uploadFileLe.SetHeight(value)
	mw.browserBtn.SetHeight(value)

	mw.sipLb.SetHeight(value)
	mw.sipLe.SetHeight(value)
	mw.uploadBtn.SetHeight(value)
}
