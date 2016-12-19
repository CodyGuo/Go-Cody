package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

import (
	// "github.com/codyguo/win"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var _VERSION_ = "unknown"

const (
	INFO  = "[ INFO] "
	DEUBG = "[DEBUG] "
)

type MyWindow struct {
	*walk.MainWindow
	iman   *iManIP
	NCName *walk.Label
	IP     *walk.Label
	Status *walk.Label

	Open *walk.PushButton
	Eth0 *walk.PushButton
	Eth1 *walk.PushButton
	Dhcp *walk.PushButton

	lv *LogView

	ctx    context.Context
	cancel context.CancelFunc

	urlIP string
	ok    bool
}

func (mw *MyWindow) Ping(ip string) {
	mw.iman.ctx = mw.ctx
	mw.iman.Ping(ip)
	mw.urlIP = ip
	out, _ := mw.iman.StdoutPipe()
	mw.iman.Start()

	mw.ok = true
	go func() {
		select {
		case <-mw.ctx.Done():
			fmt.Printf("%s结束ping: %s, PID: %d\n", INFO, ip, mw.iman.pingPid)
			mw.iman.Process.Kill()
			mw.ok = false
		}
	}()

	fmt.Printf("%s开始ping: %s, PID: %d\n", INFO, ip, mw.iman.Process.Pid)
	for mw.ok {
		result, _ := mw.iman.GetCmdResult(out)
		mw.lv.PostAppendText(result)
	}

	mw.iman.Wait()
	defer mw.lv.Clean()
}

func (mw *MyWindow) OpenAdmin() {
	if mw.urlIP == "" {
		walk.MsgBox(mw, "提示信息", "请先点击切换IP，再点击打开后台管理页面按钮。", walk.MsgBoxIconInformation)
		return
	}

	url := fmt.Sprintf("http://%s/admin", mw.urlIP)
	if mw.urlIP == "114.114.114.114" {
		url = "http://www.baidu.com"
	}
	walk.MsgBox(mw, "提示信息", fmt.Sprintf("正在尝试打开管理后台：%s，请耐心等待！", url), walk.MsgBoxIconInformation)
	for mw.ok {
		url = fmt.Sprintf("http://%s/admin", mw.urlIP)
		if mw.urlIP == "114.114.114.114" {
			url = "http://www.baidu.com"
		}
		info := fmt.Sprintf("%s正在尝试打开管理后台: %s\n", INFO, url)
		mw.lv.PostAppendText(info)
		resp, err := http.Get(url)
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		if resp.StatusCode == http.StatusOK {
			cmd := fmt.Sprintf("cmd /k start %s", url)
			mw.iman.DoCmd(cmd)
			mw.iman.Run()
			resp.Body.Close()
			walk.MsgBox(mw, "提示信息", "服务器连接正常，可通过浏览器进行管理！", walk.MsgBoxIconWarning)
			return
		}
		resp.Body.Close()
	}

}

func (mw *MyWindow) RunApp() {
	mw.SetFixedSize(true)
	mw.SetMaximizeBox(false)
	mw.ctx, mw.cancel = context.WithCancel(context.Background())

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "iMan初始化配置工具 " + _VERSION_,
		MinSize:  Size{500, 430},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Label{
								AssignTo: &mw.NCName,
								Text:     "本地连接",
							},
							Label{
								AssignTo: &mw.IP,
								Text:     "190.190.190.190",
							},
							Label{
								AssignTo: &mw.Status,
								Text:     "已连接",
							},
						},
					},
					Composite{
						Layout:  HBox{},
						MinSize: Size{100, 20},
						Children: []Widget{
							PushButton{
								AssignTo: &mw.Open,
								MinSize:  Size{50, 20},
								Text:     "打开管理后台",
								OnClicked: func() {
									go mw.OpenAdmin()
								},
							},
						},
					},
				},
			},

			GroupBox{
				Title:   "切换IP",
				Layout:  HBox{},
				MinSize: Size{200, 100},
				Children: []Widget{
					PushButton{
						AssignTo: &mw.Eth0,
						Text:     "ETH0",
						OnClicked: func() {
							go func() {
								mw.cancel()
								mw.ctx, mw.cancel = context.WithCancel(context.Background())
								mw.iman.SetStatic(ETH0)
								mw.Ping("190.190.190.190")
							}()

							mw.IP.SetText(ETH0 + "/16")
						},
					},
					PushButton{
						AssignTo: &mw.Eth1,
						Text:     "ETH1",
						OnClicked: func() {
							go func() {
								mw.cancel()
								mw.ctx, mw.cancel = context.WithCancel(context.Background())
								mw.iman.SetStatic(ETH1)
								mw.Ping("191.191.191.191")
							}()
							mw.IP.SetText(ETH1 + "/16")
						},
					},
					PushButton{
						AssignTo: &mw.Dhcp,
						Text:     "DHCP",
						OnClicked: func() {
							go func() {
								mw.cancel()
								mw.ctx, mw.cancel = context.WithCancel(context.Background())
								mw.iman.SetDHCP()
								mw.Ping("114.114.114.114")
							}()
							mw.IP.SetText("169.254.130.60/16")
						},
					},
				},
			},
		},
	}.CreateCody()); err != nil {
		fmt.Println("RunApp", err)
		return
	}

	mw.lv, _ = NewLogView(mw)
	log.SetOutput(mw.lv)

	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		mw.iman.Process.Kill()
		cmd := "taskkill /F /IM ping.exe"
		mw.iman.DoCmd(cmd)
		mw.iman.Run()
	})

	// var rect win.RECT

	// hWnd := win.HWND(mw.Handle())
	// win.GetWindowRect(hWnd, &rect)

	// hRgn := win.CreateRoundRectRgn(0, 0, rect.Right-rect.Left, rect.Bottom-rect.Top, 4, 4)
	// win.SetWindowRgn(hWnd, hRgn, true)
	// win.DeleteObject(win.HGDIOBJ(hRgn))

	mw.Run()
}

func main() {
	var err error
	mw := new(MyWindow)
	mw.iman, err = NewiManIP()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			mw.iman.ShowStatus()
			fmt.Printf("%s自动更新网卡信息: [%s] => [%s]\n", DEUBG, mw.iman.Name, mw.iman.IP)
			mw.NCName.SetText(mw.iman.Name)
			mw.IP.SetText(mw.iman.IP)
			mw.Status.SetText(mw.iman.Status)
			time.Sleep(3 * time.Second)
		}

	}()

	mw.RunApp()
}
