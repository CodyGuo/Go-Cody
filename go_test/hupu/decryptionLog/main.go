package main

import (
	// "fmt"
	"log"
	"os/exec"
	"strings"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	mw := new(MyWindow)
	mw.RunApp()
}

type MyWindow struct {
	*walk.MainWindow
	textEdit *walk.TextEdit
	tv       *walk.TableView
}

func (mw *MyWindow) RunApp() {
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "iMan日志解密工具 1.0",
		Layout:   VBox{},
		MinSize:  Size{600, 450},
		Children: []Widget{
			Composite{
				Layout:  HBox{},
				MaxSize: Size{0, 50},
				Children: []Widget{
					// VSpacer{MinSize: Size{50, 0}},
					PushButton{
						Text: "开始解密日志",
						OnClicked: func() {
							walk.MsgBox(mw, "title", "message", walk.MsgBoxIconInformation)
						},
					},
					PushButton{
						Text: "打开输出文件夹",
						OnClicked: func() {
							exec.Command("cmd", "/c", "start", "logout").Run()
						},
					},
					Composite{
						Layout:  HBox{},
						MinSize: Size{60, 60},
						MaxSize: Size{60, 60},
						Children: []Widget{
							TextEdit{
								AssignTo: &mw.textEdit,
								Text:     "请将日志文件拖放到这里!",
								ReadOnly: true,
							},
						},
					},
				},
			},

			TableView{
				AssignTo: &mw.tv,
			},
		},
	}.CreateCody()); err != nil {
		log.Fatal(err)
	}

	mw.textEdit.DropFiles().Attach(func(files []string) {
		mw.textEdit.SetText(strings.Join(files, "\r\n"))
	})
	mw.Run()
}
