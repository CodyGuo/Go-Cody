// 邮件群发

package main

import (
    // "fmt"
    // "image"
    // _ "image/png"
    "log"
    // "os"
)

import (
    "github.com/lxn/walk"
)

type myDialogUI struct {
    mailListGb, mailServerGb, mailContentGb *walk.GroupBox
    mailListLe, mailBodyLe                  *walk.TextEdit
    userNameLb, passwdLb, smtpLb            *walk.Label
    userNameLe, passwdLe, smtpLe, portLe    *walk.LineEdit

    mailSubLb, mailAdjLb, mailBodyLb *walk.Label
    mailSubLe, mailAdjLe             *walk.LineEdit
    browseBtn, sendBtn               *walk.PushButton

    logLb *walk.Label
    lv    *LogView
}

type MyDialog struct {
    *walk.Dialog
    ui  myDialogUI
    ni  *walk.NotifyIcon

    adjFile string
}

func (mw *MyDialog) checkError(err error) error {
    if err != nil {
        return err
        log.Fatalln(err)
    }
    return nil
}

func (mw *MyDialog) init(owner walk.Form) (err error) {
    // 设置最小化
    mw.SetMinimizeBox(true)
    // 禁用最大化
    mw.SetMaximizeBox(false)
    // 窗口屏幕居中
    mw.SetFixedSize(true)

    mw.Dialog, err = walk.NewDialog(owner)
    mw.checkError(err)

    succeeded := false
    defer func() {
        if !succeeded {
            mw.Dispose()
        }
    }()

    // 设置主窗体大小
    mw.SetClientSize(walk.Size{600, 450})
    mw.checkError(err)

    // 设置主窗体标题
    mw.SetTitle("邮件群发-by Cody.Guo")
    mw.checkError(err)

    // 设置上传组合窗体
    mw.ui.mailListGb, err = walk.NewGroupBox(mw)
    mw.checkError(err)

    err = mw.ui.mailListGb.SetTitle("发送邮件列表")
    mw.checkError(err)

    err = mw.ui.mailListGb.SetBounds(walk.Rectangle{7, 8, 165, 434})
    mw.checkError(err)

    // 邮件编辑框
    mw.ui.mailListLe, err = walk.NewTextEdit(mw.ui.mailListGb)
    mw.checkError(err)

    err = mw.ui.mailListLe.SetBounds(walk.Rectangle{7, 25, 155, 405})
    mw.checkError(err)

    // 邮件服务器配置
    mw.ui.mailServerGb, err = walk.NewGroupBox(mw)
    mw.checkError(err)
    err = mw.ui.mailServerGb.SetTitle("邮件服务器配置")
    mw.checkError(err)
    err = mw.ui.mailServerGb.SetBounds(walk.Rectangle{185, 8, 405, 113})
    mw.checkError(err)

    // 用户名
    mw.ui.userNameLb, err = walk.NewLabel(mw.ui.mailServerGb)
    mw.checkError(err)
    err = mw.ui.userNameLb.SetText("用户名:")
    mw.checkError(err)
    err = mw.ui.userNameLb.SetBounds(walk.Rectangle{8, 20, 80, 20})
    mw.checkError(err)

    // 用户名编辑框
    mw.ui.userNameLe, err = walk.NewLineEdit(mw.ui.mailServerGb)
    mw.checkError(err)
    mw.ui.userNameLe.SetText(MS.UserName)

    err = mw.ui.userNameLe.SetBounds(walk.Rectangle{110, 20, 150, 20})
    mw.checkError(err)

    // 密码
    mw.ui.passwdLb, err = walk.NewLabel(mw.ui.mailServerGb)
    mw.checkError(err)

    err = mw.ui.passwdLb.SetText("密码:")
    mw.checkError(err)

    err = mw.ui.passwdLb.SetBounds(walk.Rectangle{8, 50, 80, 20})
    mw.checkError(err)

    // 密码编辑框
    mw.ui.passwdLe, err = walk.NewLineEdit(mw.ui.mailServerGb)
    mw.checkError(err)
    mw.ui.passwdLe.SetText(MS.Passwd)

    err = mw.ui.passwdLe.SetBounds(walk.Rectangle{110, 50, 150, 20})
    mw.checkError(err)

    mw.ui.passwdLe.SetPasswordMode(true)

    // SMTP服务器
    mw.ui.smtpLb, err = walk.NewLabel(mw.ui.mailServerGb)
    mw.checkError(err)

    err = mw.ui.smtpLb.SetText("SMTP服务器:")
    mw.checkError(err)

    err = mw.ui.smtpLb.SetBounds(walk.Rectangle{8, 80, 90, 20})
    mw.checkError(err)

    // // SMTP服务器编辑框
    mw.ui.smtpLe, err = walk.NewLineEdit(mw.ui.mailServerGb)
    mw.checkError(err)
    mw.ui.smtpLe.SetText(MS.Smtp)

    err = mw.ui.smtpLe.SetBounds(walk.Rectangle{110, 80, 150, 20})
    mw.checkError(err)

    // PORT 25
    mw.ui.portLe, err = walk.NewLineEdit(mw.ui.mailServerGb)
    mw.checkError(err)

    mw.ui.portLe.SetBounds(walk.Rectangle{270, 80, 30, 20})
    mw.ui.portLe.SetText(MS.Port)

    // 发送配置
    mw.ui.mailContentGb, err = walk.NewGroupBox(mw)
    mw.checkError(err)

    mw.ui.mailContentGb.SetTitle("发送配置")
    err = mw.ui.mailContentGb.SetBounds(walk.Rectangle{185, 135, 405, 187})
    mw.checkError(err)

    // 邮件主题
    mw.ui.mailSubLb, err = walk.NewLabel(mw.ui.mailContentGb)
    mw.checkError(err)

    mw.ui.mailSubLb.SetText("邮件主题:")
    err = mw.ui.mailSubLb.SetBounds(walk.Rectangle{8, 20, 80, 20})
    mw.checkError(err)

    // 邮件主题编辑框
    mw.ui.mailSubLe, err = walk.NewLineEdit(mw.ui.mailContentGb)
    mw.checkError(err)
    mw.ui.mailSubLe.SetText(MS.Subject)

    err = mw.ui.mailSubLe.SetBounds(walk.Rectangle{110, 20, 150, 20})
    mw.checkError(err)

    // 邮件内容
    mw.ui.mailBodyLb, err = walk.NewLabel(mw.ui.mailContentGb)
    mw.checkError(err)

    mw.ui.mailBodyLb.SetText("邮件内容:")
    err = mw.ui.mailBodyLb.SetBounds(walk.Rectangle{8, 50, 80, 20})
    mw.checkError(err)

    // 邮件内容编辑框
    mw.ui.mailBodyLe, err = walk.NewTextEdit(mw.ui.mailContentGb)
    mw.checkError(err)

    mw.ui.mailBodyLe.SetText(MS.Body)

    err = mw.ui.mailBodyLe.SetBounds(walk.Rectangle{110, 50, 280, 100})
    mw.checkError(err)

    // 附件
    mw.ui.mailAdjLb, err = walk.NewLabel(mw.ui.mailContentGb)
    mw.checkError(err)

    mw.ui.mailAdjLb.SetText("附件:")
    err = mw.ui.mailAdjLb.SetBounds(walk.Rectangle{8, 160, 80, 20})
    mw.checkError(err)

    // 附件编辑框
    mw.ui.mailAdjLe, err = walk.NewLineEdit(mw.ui.mailContentGb)
    mw.checkError(err)

    err = mw.ui.mailAdjLe.SetBounds(walk.Rectangle{110, 160, 150, 20})
    mw.checkError(err)

    mw.ui.mailAdjLe.SetText(MS.Adjunct)

    mw.ui.mailAdjLe.SetReadOnly(true)

    // 浏览
    mw.ui.browseBtn, err = walk.NewPushButton(mw.ui.mailContentGb)
    mw.checkError(err)
    mw.ui.browseBtn.SetText("浏览")
    err = mw.ui.browseBtn.SetBounds(walk.Rectangle{270, 155, 35, 30})
    mw.checkError(err)

    // 日志
    mw.ui.logLb, err = walk.NewLabel(mw)
    mw.checkError(err)

    mw.ui.logLb.SetText("日志")

    err = mw.ui.logLb.SetBounds(walk.Rectangle{185, 325, 30, 20})
    mw.checkError(err)

    // 日志输出
    mw.ui.lv, err = NewLogView(mw)
    mw.checkError(err)

    err = mw.ui.lv.SetBounds(walk.Rectangle{185, 345, 405, 70})
    mw.checkError(err)

    log.SetOutput(mw.ui.lv)

    // 发送邮件
    mw.ui.sendBtn, err = walk.NewPushButton(mw)
    mw.checkError(err)

    mw.ui.sendBtn.SetText("开始发送")

    err = mw.ui.sendBtn.SetBounds(walk.Rectangle{350, 420, 80, 30})
    mw.checkError(err)

    // // 设置背景
    // color := walk.RGB(255, 0, 0)
    // bg, _ := walk.NewSolidColorBrush(color)
    // mw.SetBackground(bg)

    // 设置字体和图标
    fount, _ := walk.NewFont("宋体", 11, 0)
    mw.SetFont(fount)
    mw.ui.mailListGb.SetFont(fount)
    mw.ui.mailServerGb.SetFont(fount)
    mw.ui.mailContentGb.SetFont(fount)

    succeeded = true

    return nil
}

func (mw *MyDialog) setMyNotify() (err error) {
    // 托盘图标
    // icon, _ := walk.NewIconFromFile("../../img/main.ico")
    icon, _ := walk.NewIconFromResourceId(3)
    mw.ni, err = walk.NewNotifyIcon()
    mw.checkError(err)

    mw.SetIcon(icon)
    // Set the icon and a tool tip text.
    err = mw.ni.SetIcon(icon)
    mw.checkError(err)

    mw.ni.SetToolTip("测试程序")

    // The notify icon is hidden initially, so we have to make it visible.
    err = mw.ni.SetVisible(false)
    mw.checkError(err)

    return nil
}

func (mw *MyDialog) addMyNotifyAction() (err error) {
    // We put an exit action into the context menu.
    exitAction := walk.NewAction()
    err = exitAction.SetText("E&xit")
    mw.checkError(err)

    exitAction.Triggered().Attach(func() {
        mw.Dispose()    // 释放主程序
        mw.ni.Dispose() // 右下角图标退出
        walk.App().Exit(1)
    })
    // 增加快捷键
    exitAction.SetShortcut(walk.Shortcut{walk.ModShift, walk.KeyB})
    // 提示信息
    exitAction.SetToolTip("退出程序.")
    err = mw.ni.ContextMenu().Actions().Add(exitAction)
    mw.checkError(err)

    return nil
}

func (mw *MyDialog) setExitHide(exit bool) (err error) {
    if exit {
        mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
            reason = walk.CloseReasonUnknown
            var closingPublisher walk.CloseEventPublisher
            // 不关闭程序
            *canceled = true
            closingPublisher.Publish(canceled, reason)
            // 隐藏程序,显示托盘
            mw.Hide()
            mw.ni.SetVisible(true)
        })
    }
    return nil
}

func (mw *MyDialog) browseFile() {
    if err := mw.openFile(); err != nil {
        log.Print(err)
    }
}

func (mw *MyDialog) openFile() error {
    dlgFile := new(walk.FileDialog)

    dlgFile.FilePath = mw.adjFile
    dlgFile.Filter = "附件 (*)|*"
    dlgFile.Title = "选择附件"

    if ok, err := dlgFile.ShowOpen(mw); err != nil {
        return err
    } else if !ok {
        return nil
    }

    mw.adjFile = dlgFile.FilePath

    return nil
}

func (mw *MyDialog) myMsg(title, message string, style walk.MsgBoxStyle) {
    switch style {
    case walk.MsgBoxIconInformation:
        mw.ni.ShowInfo(title, message)
    case walk.MsgBoxIconWarning:
        mw.ni.ShowWarning(title, message)
    }
    log.Println(message)
    walk.MsgBox(mw, title, message, style)
}
