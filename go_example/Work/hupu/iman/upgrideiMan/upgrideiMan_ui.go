// iMan 升级程序 ui 控制

package main

import (
	"log"
)

import (
	"github.com/lxn/walk"
)

var _VERSION_ = "cody.guo"

type myDialogUI struct {
	uploadBtn           *walk.PushButton
	fileLe, ipLe        *walk.LineEdit
	uploadGb            *walk.GroupBox
	fileLb, ipLb, logLb *walk.Label
	browseBtn           *walk.PushButton
	lv                  *LogView
}

type MyDialog struct {
	*walk.Dialog
	ui          myDialogUI
	upgrideFile string
	ni          *walk.NotifyIcon
}

func (mw *MyDialog) checkError(err error) {
	if err != nil {
		mw.myMsg("失败信息", "上传过程出错."+err.Error(), walk.MsgBoxIconError)
	}
}

func (mw *MyDialog) init(owner walk.Form) (err error) {
	// 设置最小化
	mw.SetMinimizeBox(true)
	// 禁用最大化
	mw.SetMaximizeBox(false)
	// 设置窗口固定
	mw.SetFixedSize(true)

	mw.Dialog, err = walk.NewDialogCody(owner)
	mw.checkError(err)

	succeeded := false
	defer func() {
		if !succeeded {
			mw.Dispose()
		}
	}()

	// 设置主窗体大小
	mw.SetClientSize(walk.Size{376, 287})
	mw.checkError(err)

	// 居中
	mw.SetScreenCenter(true)

	// 设置主窗体在所有窗体之前
	mw.SetForegroundWindow()
	mw.SwitchToThisWindow(true)

	// 设置主窗体标题
	mw.SetTitle("iMan-升级工具   V【" + _VERSION_ + "】")
	mw.checkError(err)

	// 设置上传组合窗体
	mw.ui.uploadGb, err = walk.NewGroupBox(mw)
	mw.checkError(err)

	err = mw.ui.uploadGb.SetTitle("升级包上传")
	mw.checkError(err)

	err = mw.ui.uploadGb.SetBounds(walk.Rectangle{3, 7, 368, 138})
	mw.checkError(err)

	// 升级包
	mw.ui.fileLb, err = walk.NewLabel(mw.ui.uploadGb)
	mw.checkError(err)

	err = mw.ui.fileLb.SetText("升级包:")
	mw.checkError(err)

	err = mw.ui.fileLb.SetBounds(walk.Rectangle{10, 33, 70, 25})
	mw.checkError(err)

	// 上传路径
	mw.ui.fileLe, err = walk.NewLineEdit(mw.ui.uploadGb)
	mw.checkError(err)

	err = mw.ui.fileLe.SetBounds(walk.Rectangle{96, 33, 166, 25})
	mw.checkError(err)

	err = mw.ui.fileLe.SetReadOnly(true)
	mw.checkError(err)

	// 浏览按钮
	mw.ui.browseBtn, err = walk.NewPushButton(mw.ui.uploadGb)
	mw.checkError(err)

	err = mw.ui.browseBtn.SetText("浏览")
	mw.checkError(err)

	err = mw.ui.browseBtn.SetBounds(walk.Rectangle{288, 34, 55, 25})
	mw.checkError(err)

	mw.ui.browseBtn.SetCursor(walk.CursorHand())

	// 服务器IP lb
	mw.ui.ipLb, err = walk.NewLabel(mw.ui.uploadGb)
	mw.checkError(err)

	err = mw.ui.ipLb.SetText("服务器IP:")
	mw.checkError(err)

	err = mw.ui.ipLb.SetBounds(walk.Rectangle{10, 94, 70, 25})
	mw.checkError(err)

	// 服务器IP le
	mw.ui.ipLe, err = walk.NewLineEdit(mw.ui.uploadGb)
	mw.checkError(err)

	err = mw.ui.ipLe.SetBounds(walk.Rectangle{96, 92, 166, 25})
	mw.checkError(err)

	// 控制IP输入长度为15
	mw.ui.ipLe.SetMaxLength(15)

	// 上传按钮
	mw.ui.uploadBtn, err = walk.NewPushButton(mw.ui.uploadGb)
	mw.checkError(err)

	// // // 设置背景
	// color := walk.RGB(0, 255, 0)
	// bg, _ := walk.NewSolidColorBrush(color)
	// mw.SetBackground(bg)
	// mw.ui.uploadBtn.SetBackground(bg)
	// mw.ui.uploadBtn.ImageAboveText()

	err = mw.ui.uploadBtn.SetText("上传")
	mw.checkError(err)

	err = mw.ui.uploadBtn.SetBounds(walk.Rectangle{288, 92, 55, 25})
	mw.checkError(err)

	mw.ui.uploadBtn.SetCursor(walk.CursorHand())

	// 日志
	mw.ui.logLb, err = walk.NewLabel(mw)
	mw.checkError(err)

	err = mw.ui.logLb.SetText("日志")
	mw.checkError(err)

	err = mw.ui.logLb.SetBounds(walk.Rectangle{5, 152, 29, 13})
	mw.checkError(err)

	// 日志输出
	mw.ui.lv, err = NewLogView(mw)
	mw.checkError(err)

	err = mw.ui.lv.SetBounds(walk.Rectangle{6, 172, 365, 106})
	mw.checkError(err)

	log.SetOutput(mw.ui.lv)

	// 设置字体和图标
	fount, _ := walk.NewFont("幼圆", 10, walk.FontBold)
	mw.ui.logLb.SetFont(fount)
	mw.ui.uploadGb.SetFont(fount)

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

	return nil
}

func (mw *MyDialog) addMyNotifyAction() (err error) {
	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	err = exitAction.SetText("退出程序")
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
	mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		reason = walk.CloseReasonUnknown
		var closingPublisher walk.CloseEventPublisher
		// 不关闭程序
		*canceled = exit
		closingPublisher.Publish(canceled, reason)
		// 隐藏程序,显示托盘
		mw.Hide()
		// mw.ni.SetVisible(true)
	})

	return nil
}

func (mw *MyDialog) uploadFile() {
	if err := mw.openFile(); err != nil {
		log.Print(err)
	}
}

func (mw *MyDialog) openFile() error {
	dlgFile := new(walk.FileDialog)

	dlgFile.FilePath = mw.upgrideFile
	dlgFile.Filter = "iMan Files 最大200MB|*"
	dlgFile.Title = "选择iMan升级包"

	if ok, err := dlgFile.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.upgrideFile = dlgFile.FilePath

	return nil
}

func (mw *MyDialog) myMsg(title, message string, style walk.MsgBoxStyle) (result int) {
	switch style {
	case walk.MsgBoxIconInformation, walk.MsgBoxOKCancel + walk.MsgBoxIconInformation:
		mw.ni.ShowInfo(title, message)
	case walk.MsgBoxIconWarning:
		mw.ni.ShowWarning(title, message)
	case walk.MsgBoxIconError:
		mw.ni.ShowError(title, message)
	}
	// log.Println(message)
	result = walk.MsgBox(mw, title, message, style)

	return
}
