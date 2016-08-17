package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	button *walk.PushButton
	label  *walk.Label
}

func main() {
	mw := new(MyMainWindow)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "测试程序",
		Size:     Size{300, 200},
		Layout:   HBox{},
		Children: []Widget{
			PushButton{AssignTo: &mw.button, Text: "打开", OnClicked: func() { doSomething(mw) }},
			// OnMouseDown: func(x, y int, button walk.MouseButton) {
			// 	log.Println("down", x, y, button)
			// 	if button == walk.LeftButton {
			// 		doSomething(mw)
			// 	}
			// },
			// OnMouseUp: func(x, y int, button walk.MouseButton) {
			// 	log.Println("up", x, y, button)
			// 	if button == walk.LeftButton {
			// 		doSomething(mw)
			// 	}
			// }},
			Label{AssignTo: &mw.label, Text: "|"},
			Label{AssignTo: &mw.label, Text: "打开",
				OnMouseDown: func(x, y int, button walk.MouseButton) {
					log.Println("down", x, y, button)
					if button == walk.LeftButton {
						doSomething(mw)
					}
				},
				OnMouseUp: func(x, y int, button walk.MouseButton) {
					log.Println("up", x, y, button)
					if button == walk.LeftButton {
						doSomething(mw)
					}
				},
				OnMouseMove: func(x, y int, button walk.MouseButton) {
					log.Println("move", x, y, button)
					if button == walk.LeftButton {
						doSomething(mw)
					}

				}},
		},
	}).Create(); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func doSomething(owner walk.Form) {
	walk.MsgBox(owner, "title", "message", walk.MsgBoxIconInformation)
}
