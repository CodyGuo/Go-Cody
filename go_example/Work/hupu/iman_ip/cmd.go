package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"syscall"
)

import (
	"github.com/axgle/mahonia"
)

const (
	ETH0 = "190.190.190.2"
	ETH1 = "191.191.191.2"
	DHCP = "dhcp"
)

type iManIP struct {
	Name    string
	IP      string
	Status  string
	pingPid int

	windows [2]string

	*exec.Cmd
	ctx context.Context
	out io.ReadCloser
}

func NewiManIP() (iman *iManIP, err error) {
	iman = new(iManIP)
	iman.windows = [2]string{"本地连接", "以太网"}

	result, err := iman.ShowStatus()
	if err != nil {
		return nil, fmt.Errorf("NewiManIP: %v", err)
	}

	connect := "已连接"
	unconnected := "已断开连接"
	if strings.Contains(result, connect) {
		iman.Status = connect
	} else {
		iman.Status = unconnected
	}

	return iman, nil
}

func (m *iManIP) ShowStatus() (result string, err error) {
	for _, win := range m.windows {
		cmd := fmt.Sprintf("netsh interface show interface name=%s", win)
		m.DoCmd(cmd)
		out, err := m.StdoutPipe()
		if err != nil {
			return "", err
		}
		m.Start()
		result, err = m.GetCmdResult(out)
		if err != nil && !strings.Contains(result, "此名称的接口未与路由器一起注册") {
			return "", fmt.Errorf("showNC: %v", err)
		}

		// 确定网卡成名，启用被禁用的网卡
		if strings.Contains(result, "管理状态") {
			m.Name = win
			if strings.Contains(result, "已禁用") {
				m.enableNC(true)
			}

			if err := m.getIP(); err != nil {
				return "", err
			}

			// 发现后正确的网卡后退出循环
			break
		}

		m.Wait()
	}

	return result, nil
}

func (m *iManIP) SetDHCP() error {
	cmd := fmt.Sprintf("netsh interface ip set address name=%s source=%s", m.Name, DHCP)
	if err := m.DoCmd(cmd); err != nil {
		return fmt.Errorf("SetDHCP: %v", err)
	}

	return m.Run()
}

func (m *iManIP) SetStatic(ip string) error {
	cmd := fmt.Sprintf("netsh interface ip set address name=%s source=static addr=%s mask=255.255.0.0",
		m.Name, ip)
	if err := m.DoCmd(cmd); err != nil {
		return fmt.Errorf("SetStatic: %v", err)
	}

	return m.Run()
}

func (m *iManIP) Ping(ip string) (err error) {
	cmd := fmt.Sprintf("ping %s -t", ip)

	return m.DoCmd(cmd)
}

func (m *iManIP) DoCmd(cmd string) error {
	if len(cmd) == 0 {
		return errors.New("The cmd can't be empty!")
	}

	if m.ctx == nil {
		m.ctx = context.Background()
	}
	command := strings.Split(cmd, " ")
	m.Cmd = exec.CommandContext(m.ctx, command[0], command[1:]...)
	m.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return nil
}

func (m *iManIP) GetCmdResult(out io.ReadCloser) (result string, err error) {
	if m.Cmd == nil {
		return "", errors.New("The cmd not create!")
	}

	gbk := mahonia.NewDecoder("gbk")
	reader := gbk.NewReader(out)
	buf := make([]byte, 1<<16)
	n, err := reader.Read(buf)

	result = string(buf[:n])

	return
}

func (m *iManIP) enableNC(ok bool) error {
	action := "disable"
	if ok {
		action = "enable"
	}
	cmd := fmt.Sprintf("netsh interface set interface  %s %s", m.Name, action)
	if err := m.DoCmd(cmd); err != nil {
		return fmt.Errorf("enableNC: %v", err)
	}

	return m.Run()
}

func (m *iManIP) getIP() error {
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
