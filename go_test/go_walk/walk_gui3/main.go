package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
}

func main() {
	mw := new(MyWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		MinSize:  Size{800, 600},
		Layout:   HBox{},
		Children: []Widget{

			TabWidget{
				MinSize: Size{700, 500},
				Pages: []TabPage{
					TabPage{
						Title:   "测试1",
						MinSize: Size{500, 300},
						Layout:  HBox{},
						Children: []Widget{
							TableView{
								AlternatingRowBGColor: walk.RGB(255, 255, 224),
								CheckBoxes:            true,
								ColumnsOrderable:      true,
								Columns: []TableViewColumn{
									{Title: "#"},
									{Title: "Bar"},
									{Title: "Baz", Format: "%.2f", Alignment: AlignFar},
									{Title: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
								},
							},
						},
					},
					TabPage{
						Title:   "测试2",
						MinSize: Size{500, 300},
						Layout:  HBox{},
						Children: []Widget{
							Label{
								Text: "测试",
							},
							PushButton{
								Text: "提交",
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}
