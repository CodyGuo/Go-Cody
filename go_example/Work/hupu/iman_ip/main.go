package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

import (
	"github.com/axgle/mahonia"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const (
	ETH0 = "190.190.190.2"
	ETH1 = "191.191.191.2"
	DHCP = "dhcp"
)

var system [2]string = [2]string{
	"本地连接",
	"以太网",
}

type MyWindow struct {
	*walk.MainWindow
	iman   *iManIP
	NCName *walk.Label
	IP     *walk.Label
	Status *walk.Label

	Eth0 *walk.PushButton
	Eth1 *walk.PushButton
	Dhcp *walk.PushButton

	lv *LogView

	ctx    context.Context
	cancel context.CancelFunc
}

func (mw *MyWindow) SetIP(value string) {
	mw.iman.SetIP(value)
	switch value {
	case ETH0:
		mw.Ping("190.190.190.190")
	case ETH1:
		mw.Ping("191.191.191.191")
	default:
		mw.Ping("114.114.114.114")
	}
}

func (mw *MyWindow) Ping(value string) {
	cmd := exec.CommandContext(mw.ctx, "ping", value, "-t")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Println("ping", err)
		return
	}

	var ok bool = true
	go func() {
		select {
		case <-mw.ctx.Done():
			// mw.cancel()
			cmd.Process.Kill()
			ok = false
		}
	}()

	for ok {
		gbk := mahonia.NewDecoder("gbk")
		buf := gbk.NewReader(out)
		data := make([]byte, 1<<16)
		n, _ := buf.Read(data)
		mw.lv.PostAppendText(string(data[:n]))
		// fmt.Println("正在输出", cmd.Process.Pid)
	}

	cmd.Wait()
}

func (mw *MyWindow) RunApp() {
	mw.SetFixedSize(true)
	mw.SetMaximizeBox(false)
	mw.ctx, mw.cancel = context.WithCancel(context.Background())

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "iMan初始化配置工具",
		MinSize:  Size{450, 320},
		Layout:   VBox{},
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
								mw.SetIP(ETH0)
							}()

							mw.lv.Clean()
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
								mw.SetIP(ETH1)
							}()
							mw.lv.Clean()
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
								mw.SetIP(DHCP)
							}()
							mw.lv.Clean()
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

	mw.Run()
}

type iManIP struct {
	Name   string
	IP     string
	Status string
}

func (m *iManIP) CheckNC() error {
	result, err := m.showStatu()
	if err != nil {
		return err
	}

	// fmt.Println("CheckNC", m.Name, result, err)
	if strings.Contains(result, "已连接") {
		m.Status = "已连接"
	} else {
		m.Status = "已断开连接"
	}

	if err := m.showIP(); err != nil {
		return err
	}

	return nil
}

func (m *iManIP) SetIP(ip string) error {
	switch ip {
	case DHCP:
		if _, err := runCMD("netsh interface ip set address name=" + m.Name + " source=" + DHCP); err != nil {
			return err
		}
	default:
		if _, err := runCMD("netsh interface ip set address name=" + m.Name + " source=static addr=" + ip + " mask=255.255.0.0"); err != nil {
			return err
		}
	}

	return nil
}

func (m *iManIP) enableNC(ok bool) error {
	var action string = " disable"
	if ok {
		action = " enable"
	}
	if _, err := runCMD("netsh interface set interface " + m.Name + action); err != nil {
		return fmt.Errorf("enableNC: %v", err)
	}

	return nil
}

func (m *iManIP) showStatu() (result string, err error) {
	for win := range system {
		data, err := runCMD("netsh interface show interface name=" + system[win])
		result = string(data)
		if err != nil && !strings.Contains(result, "此名称的接口未与路由器一起注册") {
			return "", fmt.Errorf("showNC: %v", err)
		}

		// fmt.Println("system", system[win], result)
		if strings.Contains(result, "管理状态") {
			m.Name = system[win]
			if strings.Contains(result, "已禁用") {
				m.enableNC(true)
			}
			return result, nil
		}
	}

	return result, nil
}

func (m *iManIP) showIP() error {
	inter, err := net.InterfaceByName(m.Name)
	if err != nil {
		return err
	}

	ip, err := inter.Addrs()
	if err != nil {
		return err
	}

	m.IP = ip[len(ip)-1].String()

	return nil
}

func runCMD(args string) (data []byte, err error) {
	cmd := exec.Command("cmd.exe", "/k", args)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	gbk := mahonia.NewDecoder("gbk")
	buf := gbk.NewReader(out)
	data = make([]byte, 1<<16)
	n, _ := buf.Read(data)
	err = cmd.Wait()

	// fmt.Printf("runCMD %s , %v, %s\n", args, err, data[:n])

	return data[:n], err
}

func main() {
	mw := new(MyWindow)
	mw.iman = new(iManIP)

	go func() {
		for {
			mw.iman.CheckNC()
			mw.NCName.SetText(mw.iman.Name)
			mw.IP.SetText(mw.iman.IP)
			mw.Status.SetText(mw.iman.Status)
			time.Sleep(3 * time.Second)
		}

	}()

	mw.RunApp()
}
