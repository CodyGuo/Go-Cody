// iMan 升级程序 ui 控制

package main

import (
    // "image"
    // _ "image/png"
    "log"
    // "os"
)

import (
    "github.com/lxn/walk"
)

var _VERSION_ = "cody.guo"

type myWindowUI struct {
    // 设置
    SettingMenu   *walk.Menu
    SettingAction *walk.Action
    ServerAction  *walk.Action

    // 帮助
    HelpMenu    *walk.Menu
    HelpAction  *walk.Action
    AboutAction *walk.Action

    // // 编译服务器
    BuildServerLb *walk.Label

    // 开始打包
    StartPackingBtn *walk.PushButton

    //  版本配置
    VersionGb *walk.GroupBox

    VersionTypeLb   *walk.Label
    MasterVersionLb *walk.Label
    PackVersionLb   *walk.Label
    PcHelperLb      *walk.Label
    AndroidHelperLb *walk.Label
    WebSqlLb        *walk.Label

    VersionTestRadio *walk.RadioButton
    VersionOffRadio  *walk.RadioButton

    MasterVersionLe *walk.LineEdit
    PackVersionLe   *walk.LineEdit
    PcHelperLe      *walk.LineEdit
    AndroidHelperLe *walk.LineEdit
    WebSqlLe        *walk.LineEdit

    PcUploadBtn  *walk.PushButton
    AndUploadBtn *walk.PushButton
    WebSqlBtn    *walk.PushButton

    // 打包内容配置
    PackGb *walk.GroupBox

    CheckAllLb *walk.Label
    ServerLb   *walk.Label
    ControlLb  *walk.Label

    CheckAllCb      *walk.CheckBox
    JavaWebCb       *walk.CheckBox
    WebSqlCb        *walk.CheckBox
    PcHelperCb      *walk.CheckBox
    AndroidHelperCb *walk.CheckBox
    LinuxKnlCb      *walk.CheckBox
    LinuxAppCb      *walk.CheckBox
    LinuxRubyCb     *walk.CheckBox

    // 日志框
    PackTabWidget *walk.TabWidget

    // 打包日志
    PackTabPage *walk.TabPage

    // 历史版本记录
    VersionPage *walk.TabPage

    // 日志输出
    lv  *LogView

    // 历史版本记录 TableView
    VersionTableView *walk.TableView

    // 历史版本记录 - 序号
    VersionTabVieConIndex *walk.TableViewColumn

    // 历史版本记录 - 版本
    VersionTabVieConVer *walk.TableViewColumn

    // 历史版本记录 - 打包内容
    VersionTabVieConPack *walk.TableViewColumn

    // 历史版本记录 - 是否打Tag
    VersionTabVieConTag *walk.TableViewColumn

    // 历史版本记录 - Tag 路径
    VersionTabVieConTagPath *walk.TableViewColumn

    // 历史版本记录 - 打包时间
    VersionTabVieConTime *walk.TableViewColumn
}

type MyWindow struct {
    *walk.MainWindow

    ui  myWindowUI

    ni  *walk.NotifyIcon
}

func (mw *MyWindow) checkError(err error) {
    if err != nil {
        mw.MyMsg("失败信息", "上传过程出错."+err.Error(), walk.MsgBoxIconError)
        log.Println(err.Error())
    }
}

func (mw *MyWindow) init() (err error) {
    // 设置最小化
    mw.SetMinimizeBox(true)
    // 禁用最大化
    mw.SetMaximizeBox(false)
    // 设置窗口固定
    mw.SetFixedSize(true)
    // // 设置窗口前置
    // mw.SetWindowPos(true)

    mw.MainWindow, err = walk.NewMainWindow()
    mw.checkError(err)

    succeeded := false
    defer func() {
        if !succeeded {
            mw.Dispose()
        }
    }()

    // 设置主窗体大小
    err = mw.SetClientSize(walk.Size{700, 560})
    mw.checkError(err)

    // 设置主窗体标题
    err = mw.SetTitle("iMan-打包工具   V【" + _VERSION_ + "】")
    mw.checkError(err)

    // 设置
    mw.ui.SettingMenu, _ = walk.NewMenu()
    mw.ui.SettingAction = walk.NewMenuAction(mw.ui.SettingMenu)
    mw.ui.SettingAction.SetText("设置")

    mw.ui.ServerAction = walk.NewAction()
    mw.ui.ServerAction.SetText("服务器")

    mw.ui.SettingMenu.Actions().Add(mw.ui.ServerAction)

    // 帮助
    mw.ui.HelpMenu, _ = walk.NewMenu()
    mw.ui.HelpAction = walk.NewMenuAction(mw.ui.HelpMenu)
    mw.ui.HelpAction.SetText("帮助")

    mw.ui.AboutAction = walk.NewAction()
    mw.ui.AboutAction.SetText("关于")

    mw.ui.HelpMenu.Actions().Add(mw.ui.AboutAction)

    // 菜单配置
    mw.Menu().Actions().Add(mw.ui.SettingAction)
    mw.Menu().Actions().Add(mw.ui.HelpAction)

    // 设置字体和图标
    fountTitle, _ := walk.NewFont("幼圆", 10, walk.FontBold)
    otherFont, _ := walk.NewFont("幼圆", 10, 0)

    // 编译服务器IP
    mw.ui.BuildServerLb, err = walk.NewLabel(mw)
    mw.checkError(err)

    mw.ui.BuildServerLb.SetText("编译服务器IP:")
    mw.ui.BuildServerLb.SetFont(otherFont)

    mw.ui.BuildServerLb.SetBounds(walk.Rectangle{480, 10, 220, 20})

    // 开始打包
    mw.ui.StartPackingBtn, err = walk.NewPushButton(mw)
    mw.checkError(err)

    mw.ui.StartPackingBtn.SetText("开始打包")

    mw.ui.StartPackingBtn.SetBounds(walk.Rectangle{310, 20, 75, 30})

    // 版本配置
    mw.ui.VersionGb, err = walk.NewGroupBox(mw)
    mw.checkError(err)
    mw.ui.VersionGb.SetTitle("版本配置")
    mw.ui.VersionGb.SetFont(otherFont)

    err = mw.ui.VersionGb.SetBounds(walk.Rectangle{10, 60, 330, 260})
    mw.checkError(err)

    // 版本类型
    mw.ui.VersionTypeLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.VersionTypeLb.SetText("版本类型:")
    mw.ui.VersionTypeLb.SetFont(fountTitle)

    mw.ui.VersionTypeLb.SetBounds(walk.Rectangle{10, 20, 70, 25})

    // 测试版
    mw.ui.VersionTestRadio, err = walk.NewRadioButton(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.VersionTestRadio.SetText("测试版")

    mw.ui.VersionTestRadio.SetBounds(walk.Rectangle{110, 20, 60, 25})
    mw.ui.VersionTestRadio.SetChecked(true)

    // 正式版
    mw.ui.VersionOffRadio, err = walk.NewRadioButton(mw.ui.VersionGb)
    mw.checkError(err)
    mw.ui.VersionOffRadio.SetText("正式版")

    mw.ui.VersionOffRadio.SetBounds(walk.Rectangle{180, 20, 70, 25})

    // 主版本号
    mw.ui.MasterVersionLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.MasterVersionLb.SetText("主版本号:")
    mw.ui.MasterVersionLb.SetFont(fountTitle)

    mw.ui.MasterVersionLb.SetBounds(walk.Rectangle{10, 60, 70, 25})

    // 主版本号内容
    mw.ui.MasterVersionLe, err = walk.NewLineEdit(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.MasterVersionLe.SetBounds(walk.Rectangle{110, 60, 60, 25})

    // 生成版本号
    mw.ui.PackVersionLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.PackVersionLb.SetText("生成版本:")
    mw.ui.PackVersionLb.SetFont(fountTitle)

    mw.ui.PackVersionLb.SetBounds(walk.Rectangle{10, 100, 70, 25})

    // 生成版本号内容
    mw.ui.PackVersionLe, err = walk.NewLineEdit(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.PackVersionLe.SetEnabled(false)

    mw.ui.PackVersionLe.SetBounds(walk.Rectangle{110, 100, 140, 25})

    // PC助手
    mw.ui.PcHelperLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.PcHelperLb.SetText("PC助手:")
    mw.ui.PcHelperLb.SetFont(fountTitle)

    mw.ui.PcHelperLb.SetBounds(walk.Rectangle{10, 140, 70, 25})

    // PC助手上传路径
    mw.ui.PcHelperLe, err = walk.NewLineEdit(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.PcHelperLe.SetBounds(walk.Rectangle{110, 140, 150, 25})

    // PC助手上传按钮
    mw.ui.PcUploadBtn, err = walk.NewPushButton(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.PcUploadBtn.SetText("上传")
    mw.ui.PcUploadBtn.SetBounds(walk.Rectangle{270, 140, 50, 25})

    // Android助手
    mw.ui.AndroidHelperLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.AndroidHelperLb.SetText("Android助手:")
    mw.ui.AndroidHelperLb.SetFont(fountTitle)

    mw.ui.AndroidHelperLb.SetBounds(walk.Rectangle{10, 180, 100, 25})

    // Android助手上传路径
    mw.ui.AndroidHelperLe, err = walk.NewLineEdit(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.AndroidHelperLe.SetBounds(walk.Rectangle{110, 180, 150, 25})

    // Android助手上传按钮
    mw.ui.AndUploadBtn, err = walk.NewPushButton(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.AndUploadBtn.SetText("上传")

    mw.ui.AndUploadBtn.SetBounds(walk.Rectangle{270, 180, 50, 25})

    // Web 数据库
    mw.ui.WebSqlLb, err = walk.NewLabel(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.WebSqlLb.SetText("Web 数据库:")
    mw.ui.WebSqlLb.SetFont(fountTitle)

    mw.ui.WebSqlLb.SetBounds(walk.Rectangle{10, 220, 90, 25})

    // Web 数据库上传路径
    mw.ui.WebSqlLe, err = walk.NewLineEdit(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.WebSqlLe.SetBounds(walk.Rectangle{110, 220, 150, 25})

    // Web 数据库上传按钮
    mw.ui.WebSqlBtn, err = walk.NewPushButton(mw.ui.VersionGb)
    mw.checkError(err)

    mw.ui.WebSqlBtn.SetText("上传")

    mw.ui.WebSqlBtn.SetBounds(walk.Rectangle{270, 220, 50, 25})

    // 打包内容配置
    mw.ui.PackGb, err = walk.NewGroupBox(mw)
    mw.checkError(err)

    mw.ui.PackGb.SetTitle("打包内容配置")
    mw.ui.PackGb.SetFont(otherFont)

    mw.ui.PackGb.SetBounds(walk.Rectangle{355, 60, 335, 260})

    // 全选
    mw.ui.CheckAllLb, err = walk.NewLabel(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.CheckAllLb.SetText("全选:")
    mw.ui.CheckAllLb.SetFont(fountTitle)

    mw.ui.CheckAllLb.SetBounds(walk.Rectangle{10, 20, 55, 25})

    // 全选框
    mw.ui.CheckAllCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)
    mw.ui.CheckAllCb.SetVisible(true)

    mw.ui.CheckAllCb.SetBounds(walk.Rectangle{80, 20, 70, 25})

    // 服务器
    mw.ui.ServerLb, err = walk.NewLabel(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.ServerLb.SetText("服务器:")
    mw.ui.ServerLb.SetFont(fountTitle)

    mw.ui.ServerLb.SetBounds(walk.Rectangle{10, 50, 55, 25})

    // Java Web复选框
    mw.ui.JavaWebCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.JavaWebCb.SetText("Java Web")

    mw.ui.JavaWebCb.SetBounds(walk.Rectangle{80, 50, 80, 25})

    // Web 数据库复选框
    mw.ui.WebSqlCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.WebSqlCb.SetText("Web 数据库")

    mw.ui.WebSqlCb.SetBounds(walk.Rectangle{170, 50, 90, 25})

    // PC助手
    mw.ui.PcHelperCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.PcHelperCb.SetText("PC助手")

    mw.ui.PcHelperCb.SetBounds(walk.Rectangle{80, 90, 60, 25})

    // Android助手
    mw.ui.AndroidHelperCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.AndroidHelperCb.SetText("Android助手")

    mw.ui.AndroidHelperCb.SetBounds(walk.Rectangle{170, 90, 90, 25})

    // 控制器
    mw.ui.ControlLb, err = walk.NewLabel(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.ControlLb.SetText("控制器:")
    mw.ui.ControlLb.SetFont(fountTitle)

    mw.ui.ControlLb.SetBounds(walk.Rectangle{10, 130, 55, 25})

    // Linux knl
    mw.ui.LinuxKnlCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.LinuxKnlCb.SetText("Linux Knl")

    mw.ui.LinuxKnlCb.SetBounds(walk.Rectangle{80, 130, 80, 25})

    // Linux App
    mw.ui.LinuxAppCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.LinuxAppCb.SetText("Linux App")

    mw.ui.LinuxAppCb.SetBounds(walk.Rectangle{170, 130, 80, 25})

    // Linux Ruby
    mw.ui.LinuxRubyCb, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.LinuxRubyCb.SetText("Linux Ruby")

    mw.ui.LinuxRubyCb.SetBounds(walk.Rectangle{80, 170, 90, 25})

    // 打包日志 TabWidget
    mw.ui.PackTabWidget, err = walk.NewTabWidget(mw)
    mw.checkError(err)
    mw.ui.PackTabWidget.SetBounds(walk.Rectangle{10, 330, 680, 200})

    // 打包日志 TabPage
    mw.ui.PackTabPage, err = walk.NewTabPage()
    mw.ui.PackTabPage.SetTitle("打包日志")

    // 历史版本记录 TabPage
    mw.ui.VersionPage, err = walk.NewTabPage()
    mw.ui.VersionPage.SetTitle("历史版本记录")

    // TabPage 添加到 TabWidget
    mw.ui.PackTabWidget.Pages().Add(mw.ui.PackTabPage)
    mw.ui.PackTabWidget.Pages().Add(mw.ui.VersionPage)

    // 打包日志 输出记录
    mw.ui.lv, err = NewLogView(mw.ui.PackTabPage)
    mw.checkError(err)
    err = mw.ui.lv.SetBounds(walk.Rectangle{-1, -1, 680, 180})
    mw.checkError(err)

    log.SetOutput(mw.ui.lv)

    // 历史版本记录 TableView
    mw.ui.VersionTableView, err = walk.NewTableView(mw.ui.VersionPage)
    mw.checkError(err)

    mw.ui.VersionTableView.SetBounds(walk.Rectangle{-1, -1, 680, 180})

    // 历史版本记录 - 序号
    mw.ui.VersionTabVieConIndex = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConIndex.SetTitle("序号")
    mw.ui.VersionTabVieConIndex.SetWidth(50)

    // 历史版本记录 - 版本
    mw.ui.VersionTabVieConVer = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConVer.SetTitle("版本")
    mw.ui.VersionTabVieConVer.SetWidth(100)

    // 历史版本记录 - 打包内容
    mw.ui.VersionTabVieConPack = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConPack.SetTitle("打包内容")
    mw.ui.VersionTabVieConPack.SetWidth(120)

    // 历史版本记录 - 是否打Tag
    mw.ui.VersionTabVieConTag = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConTag.SetTitle("Tag")
    mw.ui.VersionTabVieConTag.SetWidth(120)

    // 历史版本记录 - Tag 路径
    mw.ui.VersionTabVieConTagPath = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConTagPath.SetTitle("Tag路径")
    mw.ui.VersionTabVieConTagPath.SetWidth(150)

    // 历史版本记录 - 打包时间
    mw.ui.VersionTabVieConTime = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConTime.SetTitle("打包时间")
    mw.ui.VersionTabVieConTime.SetWidth(100)

    // TableViewColumn 添加到 TableView
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConIndex)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConVer)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConPack)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTag)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTagPath)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTime)

    succeeded = true

    return nil
}

func (mw *MyWindow) CheckAll(check bool) {
    mw.ui.JavaWebCb.SetChecked(check)
    mw.ui.WebSqlCb.SetChecked(check)
    mw.ui.PcHelperCb.SetChecked(check)
    mw.ui.AndroidHelperCb.SetChecked(check)
    mw.ui.LinuxKnlCb.SetChecked(check)
    mw.ui.LinuxAppCb.SetChecked(check)
    mw.ui.LinuxRubyCb.SetChecked(check)
}

func (mw *MyWindow) DisablePackGb(disable bool) {
    if disable {
        mw.ui.CheckAllCb.SetChecked(true)
        mw.ui.PackGb.SetEnabled(false)
    } else {
        mw.ui.CheckAllCb.SetChecked(false)
        mw.ui.PackGb.SetEnabled(true)
    }

}

func (mw *MyWindow) SetMyNotify() (err error) {
    // 托盘图标
    icon, _ := walk.NewIconFromResourceId(3)
    mw.ni, err = walk.NewNotifyIcon()
    mw.checkError(err)

    mw.SetIcon(icon)
    // Set the icon and a tool tip text.
    err = mw.ni.SetIcon(icon)
    mw.checkError(err)

    return nil
}

func (mw *MyWindow) AddMyNotifyAction() (err error) {
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

func (mw *MyWindow) SetExitHide(exit bool) (err error) {
    mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
        reason = walk.CloseReasonUnknown
        var closingPublisher walk.CloseEventPublisher
        // 不关闭程序
        *canceled = exit
        closingPublisher.Publish(canceled, reason)
        // 隐藏程序,显示托盘
        mw.Hide()
    })

    return nil
}

func (mw *MyWindow) UploadFile(filetype string) (file string) {
    file, err := mw.openFile(filetype)
    mw.checkError(err)

    return file
}

func (mw *MyWindow) openFile(filetype string) (file string, err error) {
    dlgFile := new(walk.FileDialog)

    switch filetype {
    case "PC助手":
        dlgFile.Filter = "PC助手(*.exe)|*.exe"
        dlgFile.Title = "选择PC助手"
    case "Android助手":
        dlgFile.Filter = "Android助手(*.apk)|*.apk"
        dlgFile.Title = "选择Android助手"
    case "Web 数据库":
        dlgFile.Filter = "Web 数据库(*.sql)|*.sql"
        dlgFile.Title = "选择Web 数据库"
    }

    if ok, err := dlgFile.ShowOpen(mw); err != nil {
        return dlgFile.FilePath, err
    } else if !ok {
        return dlgFile.FilePath, err
    }

    file = dlgFile.FilePath

    return file, nil
}

func (mw *MyWindow) MyMsg(title, message string, style walk.MsgBoxStyle) (result int) {
    switch style {
    case walk.MsgBoxIconInformation, walk.MsgBoxOKCancel + walk.MsgBoxIconInformation:
        mw.ni.ShowInfo(title, message)
    case walk.MsgBoxIconWarning:
        mw.ni.ShowWarning(title, message)
    case walk.MsgBoxIconError:
        mw.ni.ShowError(title, message)
    }

    result = walk.MsgBox(mw, title, message, style)

    return
}
