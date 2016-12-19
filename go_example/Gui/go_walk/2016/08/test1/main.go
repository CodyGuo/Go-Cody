package main

import (
	"log"
)

import (
	"github.com/CodyGuo/walk"
	. "github.com/CodyGuo/walk/declarative"
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

	mw.label.Clicked().Attach(func() {
		log.Println("我被点击了.")
	})

	mw.label.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		log.Println(x, y, button)
	})

	mw.label.MouseWheel().Attach(func(x, y int, button walk.MouseButton) {
		log.Println(x, y, button)
	})

	mw.label.MouseMove().Attach(func(x, y int, button walk.MouseButton) {
		log.Println(x, y, button)
	})

	mw.label.MouseUp().Attach(func(x, y int, button walk.MouseButton) {
		log.Println(x, y, button)
	})

	mw.Run()
}

func doSomething(owner walk.Form) {
	walk.MsgBox(owner, "title", "message", walk.MsgBoxIconInformation)
}
