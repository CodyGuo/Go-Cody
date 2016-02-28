package main

import (
// "log"
)

import (
	"github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
)

func main() {
	var tool *walk.Action
	var menutool *walk.Menu

	var mw *walk.MainWindow

	mw.SetMaximizeBox(false)
	mw.SetFixedSize(true)

	mw, _ = walk.NewMainWindowCody()
	mw.SetTitle("测试")
	mw.SetSize(walk.Size{300, 200})

	menutool, _ = walk.NewMenu()
	tool = walk.NewMenuAction(menutool)
	tool.SetText("文件")
	open := walk.NewAction()
	open.SetText("打开")
	exit := walk.NewAction()
	exit.SetText("退出")

	menutool.Actions().Add(open)
	menutool.Actions().Add(exit)

	men2, _ := walk.NewMenu()
	too2 := walk.NewMenuAction(men2)
	too2.SetText("工具")

	mw.Menu().Actions().Add(tool)
	mw.Menu().Actions().Add(too2)

	mw.Show()
	mw.Run()
}
