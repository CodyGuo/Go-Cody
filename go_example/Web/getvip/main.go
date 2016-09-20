package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"text/template"
)

import (
	"github.com/CodyGuo/win"
	"github.com/lxn/walk"
)

const (
	port                 = ":8888"
	STARTF_USESHOWWINDOW = 0x00000001
	STARTF_USESTDHANDLES = 0x00000100

	engine1 = "'http://jxapi.nepian.com/ckparse/?url='"
	engine2 = "'http://api.aikantv.cc/?url='"
	engine3 = "'http://www.tuhao13.com/ckflv/?url='"
)

var (
	err        error
	requestUrl = engine1
)

type MyWindows struct {
	*walk.MainWindow
	ni *walk.NotifyIcon
}

func NewMainWindow() *MyWindows {
	mw := new(MyWindows)
	mw.MainWindow, _ = walk.NewMainWindow()
	mw.ni, err = walk.NewNotifyIcon()
	checkErr(err)

	icon, _ := walk.NewIconFromResourceId(3)
	mw.setVipIcon(icon)

	mw.addAction()
	mw.ni.SetVisible(true)

	return mw
}

func (mw *MyWindows) OpenVip() {
	log.Printf("打开浏览器输入：http://127.0.0.1%s\n", port)
	mw.showInfo()
	execCmd()
	// cmd()
	// exec.Command("cmd", "/c", "start", "/b", fmt.Sprintf("http://127.0.0.1%s", port)).Run()
}

func execCmd() {
	lpCmdLine := win.StringToBytePtr("cmd /c start http://127.0.0.1:8888")
	ret := win.WinExec(lpCmdLine, win.SW_HIDE)
	fmt.Println("返回值:", ret)
}

func (mw *MyWindows) setVipIcon(icon *walk.Icon) {
	mw.SetIcon(icon)
	mw.ni.SetIcon(icon)
}

func (mw *MyWindows) addAction() {
	checkErr(mw.openAction())
	checkErr(mw.remoteAction())
	checkErr(mw.switchEngine())
	checkErr(mw.remoteClipboard())
	checkErr(mw.exitAction())
}

func (mw *MyWindows) showInfo() {
	if err := mw.ni.ShowInfo("VIP视频", "正在打开VIP视频，请稍等..."); err != nil {
		log.Fatal(err)
	}
}

func (mw *MyWindows) showMsg(title, msg string) {
	walk.MsgBox(mw, title, msg, walk.MsgBoxIconInformation)
}

func (mw *MyWindows) openAction() error {
	openAction := walk.NewAction()
	if err := openAction.SetText("打开VIP"); err != nil {
		return err
	}
	openAction.Triggered().Attach(func() {
		mw.OpenVip()
	})
	if err := mw.ni.ContextMenu().Actions().Add(openAction); err != nil {
		return err
	}

	return nil
}

func (mw *MyWindows) switchEngine() error {
	menuEngine, _ := walk.NewMenu()
	engine, _ := mw.ni.ContextMenu().Actions().AddMenu(menuEngine)
	engine.SetText("切换视频源")

	engineAction1 := walk.NewAction()
	engineAction1.SetText("资源一")
	engineAction2 := walk.NewAction()
	engineAction2.SetText("资源二")
	engineAction3 := walk.NewAction()
	engineAction3.SetText("资源三")

	engineAction1.SetChecked(true)
	engineAction2.SetChecked(false)
	engineAction3.SetChecked(false)

	engineAction1.Triggered().Attach(func() {
		requestUrl = engine1
		engineAction1.SetChecked(true)
		engineAction2.SetChecked(false)
		engineAction3.SetChecked(false)

		mw.OpenVip()
	})

	engineAction2.Triggered().Attach(func() {
		requestUrl = engine2
		engineAction1.SetChecked(false)
		engineAction2.SetChecked(true)
		engineAction3.SetChecked(false)

		mw.OpenVip()
	})

	engineAction3.Triggered().Attach(func() {
		requestUrl = engine3
		engineAction1.SetChecked(false)
		engineAction2.SetChecked(false)
		engineAction3.SetChecked(true)

		mw.OpenVip()
	})

	menuEngine.Actions().Add(engineAction1)
	menuEngine.Actions().Add(engineAction2)
	menuEngine.Actions().Add(engineAction3)

	return nil
}

func (mw *MyWindows) remoteAction() error {
	remoteAction := walk.NewAction()
	if err := remoteAction.SetText("远程访问"); err != nil {
		return err
	}
	remoteAction.Triggered().Attach(func() {
		ip, err := getIP()
		if err != nil {
			log.Fatal(err)
		}
		mw.showMsg("远程访问", fmt.Sprintf("其他电脑在浏览器地址中输入 http://%s%s 进行访问。", ip, port))
	})
	if err := mw.ni.ContextMenu().Actions().Add(remoteAction); err != nil {
		return err
	}

	return nil
}

func (mw *MyWindows) remoteClipboard() error {
	remoteAction := walk.NewAction()
	if err := remoteAction.SetText("复制远程访问地址"); err != nil {
		return err
	}
	remoteAction.Triggered().Attach(func() {
		ip, err := getIP()
		if err != nil {
			log.Fatal(err)
		}
		// 先清空粘贴板
		walk.Clipboard().Clear()
		err = walk.Clipboard().SetText(fmt.Sprintf("http://%s%s", ip, port))
		checkErr(err)
	})
	if err := mw.ni.ContextMenu().Actions().Add(remoteAction); err != nil {
		return err
	}

	return nil

}

func (mw *MyWindows) exitAction() error {
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出VIP"); err != nil {
		return err
	}
	exitAction.Triggered().Attach(func() {
		mw.ni.Dispose()
		walk.App().Exit(0)
	})
	if err := mw.ni.ContextMenu().Actions().Add(exitAction); err != nil {
		return err
	}

	return nil
}

func GetVip(w http.ResponseWriter, r *http.Request) {
	index, _ := viewsIndexTplBytes()
	t, err := template.New("index").Parse(string(index)) //解析模板文件
	checkErr(err)

	data := map[string]interface{}{
		"requestUrl": requestUrl,
	}

	t.Execute(w, data)
}

func getIP() (ip string, err error) {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}

func main() {
	var err error
	mw := NewMainWindow()

	http.HandleFunc("/", GetVip)
	go func() {
		err = http.ListenAndServe(port, nil)
		checkErr(err)
	}()

	mw.OpenVip()
	mw.Run()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
