package main

import (
	"log"
)

import (
	"github.com/CodyGuo/walk"
	. "github.com/CodyGuo/walk/declarative"
)

func main() {
	var text *walk.TextEdit
	var lineE *walk.LineEdit

	mw, _ := walk.NewMainWindow()

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "测试程序",
		MinSize:  Size{300, 300},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "FmtLinesTrue",
				OnClicked: func() {
					text.FmtLines(true)
					lineE.Undo()
					text.Undo()
				},
			},
			PushButton{
				Text: "FmtLinesFalse",
				OnClicked: func() {
					text.FmtLines(false)
					lineE.SetCueBanner("banner")
				},
			},
			PushButton{
				Text:      "GetLine",
				OnClicked: func() { log.Println(text.GetLine(0)) },
			},
			TextEdit{
				AssignTo:    &text,
				ToolTipText: "hello golang.",
			},
			LineEdit{
				AssignTo:  &lineE,
				CueBanner: "test",
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	mw.Run()

}
