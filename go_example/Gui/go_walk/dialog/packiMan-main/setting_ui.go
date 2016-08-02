package main

import (
    "log"
)

import (
    "github.com/lxn/walk"
)

type ServerUI struct {
    IpLb *walk.Label
    IpLe *walk.LineEdit

    UserLb *walk.Label
    UserLe *walk.LineEdit

    PasswdLb *walk.Label
    PasswdLe *walk.LineEdit

    AcceptPB *walk.PushButton
}

type DlgServer struct {
    *walk.Dialog

    ui  ServerUI
}

func (dlg *DlgServer) checkError(err error) {
    if err != nil {
        log.Println(err.Error())
    }
}

func (dlg *DlgServer) init(owner walk.Form) (err error) {
    dlg.Dialog, err = walk.NewDialogWithFixedSize(owner)
    dlg.checkError(err)

    // 读取配置文件
    ConfSer.Read()

    succeeded := false
    defer func() {
        if !succeeded {
            dlg.Dispose()
        }
    }()

    // 设置主窗体大小
    err = dlg.SetClientSize(walk.Size{250, 180})
    dlg.checkError(err)

    // 设置窗体标题
    err = dlg.SetTitle("服务器设置")
    dlg.checkError(err)

    // 设置字体和图标
    fountTitle, _ := walk.NewFont("幼圆", 10, walk.FontBold)
    fountOther, _ := walk.NewFont("幼圆", 10, 0)

    // IP标题
    dlg.ui.IpLb, err = walk.NewLabel(dlg)
    dlg.ui.IpLb.SetText("IP:")
    dlg.ui.IpLb.SetFont(fountTitle)
    dlg.ui.IpLb.SetBounds(walk.Rectangle{30, 20, 60, 20})

    // IP编辑框
    dlg.ui.IpLe, err = walk.NewLineEdit(dlg)
    dlg.ui.IpLe.SetBounds(walk.Rectangle{100, 20, 120, 20})
    dlg.ui.IpLe.SetFont(fountOther)
    dlg.ui.IpLe.SetText(ConfSer.Ip)
    dlg.ui.IpLe.SetMaxLength(15)

    // 用户名标题
    dlg.ui.UserLb, err = walk.NewLabel(dlg)
    dlg.ui.UserLb.SetText("用户名:")
    dlg.ui.UserLb.SetFont(fountTitle)

    dlg.ui.UserLb.SetBounds(walk.Rectangle{30, 60, 60, 20})

    // 用户名编辑框
    dlg.ui.UserLe, err = walk.NewLineEdit(dlg)
    dlg.ui.UserLe.SetBounds(walk.Rectangle{100, 60, 120, 20})
    dlg.ui.UserLe.SetFont(fountOther)
    dlg.ui.UserLe.SetText(ConfSer.User)

    // 密码标题
    dlg.ui.PasswdLb, err = walk.NewLabel(dlg)
    dlg.ui.PasswdLb.SetText("密码:")
    dlg.ui.PasswdLb.SetFont(fountTitle)

    dlg.ui.PasswdLb.SetBounds(walk.Rectangle{30, 100, 60, 20})

    // 密码编辑框
    dlg.ui.PasswdLe, err = walk.NewLineEdit(dlg)
    dlg.ui.PasswdLe.SetBounds(walk.Rectangle{100, 100, 120, 20})
    dlg.ui.PasswdLe.SetPasswordMode(true)
    dlg.ui.PasswdLe.SetText(ConfSer.Passwd)

    // 确定
    dlg.ui.AcceptPB, err = walk.NewPushButton(dlg)
    dlg.ui.AcceptPB.SetText("确定")

    dlg.ui.AcceptPB.SetBounds(walk.Rectangle{150, 140, 70, 25})

    // 设置图标
    icon, _ := walk.NewIconFromResourceId(3)
    dlg.SetIcon(icon)

    succeeded = true

    return nil
}
