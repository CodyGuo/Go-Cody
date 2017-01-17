package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/braintree/manners"
	"github.com/lxn/walk"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

type MyWindow struct {
	*walk.MainWindow
	ni *walk.NotifyIcon
}

func NewMyWindow() *MyWindow {
	mw := new(MyWindow)
	var err error
	mw.MainWindow, err = walk.NewMainWindow()
	checkError(err)
	return mw
}

func (mw *MyWindow) init() {
	http.HandleFunc("/", handler)
}

func (mw *MyWindow) RunHttpServer() error {
	return manners.ListenAndServe(":8080", http.DefaultServeMux)
}

func (mw *MyWindow) AddNotifyIcon() {
	var err error
	mw.ni, err = walk.NewNotifyIcon()
	checkError(err)
	mw.ni.SetVisible(true)

	icon, err := walk.NewIconFromResourceId(3)
	checkError(err)
	mw.SetIcon(icon)
	mw.ni.SetIcon(icon)

	startAction := mw.addAction(nil, "start")
	stopAction := mw.addAction(nil, "stop")
	stopAction.SetEnabled(false)
	startAction.Triggered().Attach(func() {
		go func() {
			err := mw.RunHttpServer()
			if err != nil {
				mw.msgbox("start", "start http server failed.", walk.MsgBoxIconError)
				return
			}
		}()
		startAction.SetChecked(true)
		startAction.SetEnabled(false)
		stopAction.SetEnabled(true)
		mw.msgbox("start", "start http server success.", walk.MsgBoxIconInformation)
	})

	stopAction.Triggered().Attach(func() {
		ok := manners.Close()
		if !ok {
			mw.msgbox("stop", "stop http server failed.", walk.MsgBoxIconError)
		} else {
			stopAction.SetEnabled(false)
			startAction.SetChecked(false)
			startAction.SetEnabled(true)
			mw.msgbox("stop", "stop http server success.", walk.MsgBoxIconInformation)
		}
	})

	helpMenu := mw.addMenu("help")
	mw.addAction(helpMenu, "help").Triggered().Attach(func() {
		walk.MsgBox(mw, "help", "http://127.0.0.1:8080", walk.MsgBoxIconInformation)
	})

	mw.addAction(helpMenu, "about").Triggered().Attach(func() {
		walk.MsgBox(mw, "about", "http server.", walk.MsgBoxIconInformation)
	})

	mw.addAction(nil, "exit").Triggered().Attach(func() {
		mw.ni.Dispose()
		mw.Dispose()
		walk.App().Exit(0)
	})

}

func (mw *MyWindow) addMenu(name string) *walk.Menu {
	helpMenu, err := walk.NewMenu()
	checkError(err)
	help, err := mw.ni.ContextMenu().Actions().AddMenu(helpMenu)
	checkError(err)
	help.SetText(name)

	return helpMenu
}

func (mw *MyWindow) addAction(menu *walk.Menu, name string) *walk.Action {
	action := walk.NewAction()
	action.SetText(name)
	if menu != nil {
		menu.Actions().Add(action)
	} else {
		mw.ni.ContextMenu().Actions().Add(action)
	}

	return action
}

func (mw *MyWindow) msgbox(title, message string, style walk.MsgBoxStyle) {
	mw.ni.ShowInfo(title, message)
	walk.MsgBox(mw, title, message, style)
}

func main() {
	mw := NewMyWindow()

	mw.init()
	mw.AddNotifyIcon()
	mw.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
