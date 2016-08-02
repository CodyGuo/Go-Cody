/*walk自由布局
* 用Composite,不设置Layout ,所有组件SetBound
 */
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
	surface *walk.Composite
}

func main() {
	mw := new(MyMainWindow)
	width := 300
	percent := int(0.4 * float32(width))
	wnd := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Drawing Example",
		Size:     Size{width, 0},
		Layout:   Grid{Columns: 2},
		Children: []Widget{
			Label{
				MinSize: Size{percent, 0},
				Text:    "姓名：",
			},

			LineEdit{},
			HSpacer{ColumnSpan: 2, MinSize: Size{0, 20}},

			Label{
				Text: "帐号：",
			},

			LineEdit{},
			HSpacer{ColumnSpan: 2, MinSize: Size{0, 20}},

			Label{
				Text: "密码：",
			},

			LineEdit{},
			HSpacer{ColumnSpan: 2, MinSize: Size{0, 20}},
		},
	}
	if err := wnd.Create(); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}
