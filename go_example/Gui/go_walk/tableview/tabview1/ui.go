// iMan 升级程序 ui 控制

package main

import (
    "log"
)

import (
    "github.com/lxn/walk"
)

type myWindowUI struct {

    // 刷新
    PushButton1 *walk.PushButton

    // 获取
    PushButton2 *walk.PushButton

    // 日志框
    PackTabWidget *walk.TabWidget

    // 历史版本记录
    VersionPage *walk.TabPage

    // 历史版本记录 TableView
    VersionTableView *walk.TableView

    // 历史版本记录 - 序号
    VersionTabVieConIndex *walk.TableViewColumn

    // 历史版本记录 - 主版本
    VersionTabVieConMastVer *walk.TableViewColumn

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

// var (
//     _VERSION_ = "cody.guo"
// )

type MyWindow struct {
    *walk.MainWindow

    ui  myWindowUI

    ni  *walk.NotifyIcon
}

func (mw *MyWindow) checkError(err error) {
    if err != nil {
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
    err = mw.SetClientSize(walk.Size{800, 560})
    mw.checkError(err)

    // 设置主窗体标题
    err = mw.SetTitle("iMan-打包工具   V【" + _VERSION_ + "】")
    mw.checkError(err)

    mw.ui.PushButton1, err = walk.NewPushButton(mw)
    mw.checkError(err)
    mw.ui.PushButton1.SetText("刷新")
    mw.ui.PushButton1.SetBounds(walk.Rectangle{300, 10, 50, 30})

    // 获取
    mw.ui.PushButton2, err = walk.NewPushButton(mw)
    mw.checkError(err)
    mw.ui.PushButton2.SetText("获取")
    mw.ui.PushButton2.SetBounds(walk.Rectangle{400, 10, 50, 30})

    // 打包日志 TabWidget
    mw.ui.PackTabWidget, err = walk.NewTabWidget(mw)
    mw.checkError(err)
    mw.ui.PackTabWidget.SetBounds(walk.Rectangle{5, 50, 780, 500})

    // 历史版本记录 TabPage
    mw.ui.VersionPage, err = walk.NewTabPage()
    mw.ui.VersionPage.SetTitle("历史版本记录")

    // TabPage 添加到 TabWidget
    mw.ui.PackTabWidget.Pages().Add(mw.ui.VersionPage)

    // 历史版本记录 TableView
    mw.ui.VersionTableView, err = walk.NewTableView(mw.ui.VersionPage)
    mw.ui.VersionTableView.SetCheckBoxes(true)
    mw.ui.VersionTableView.SetBounds(walk.Rectangle{10, 10, 770, 460})

    mw.checkError(err)

    // 历史版本记录 - 序号
    mw.ui.VersionTabVieConIndex = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConIndex.SetTitle("序号")
    mw.ui.VersionTabVieConIndex.SetWidth(50)

    // 历史版本记录 - 主版本
    mw.ui.VersionTabVieConMastVer = walk.NewTableViewColumn()
    mw.ui.VersionTabVieConMastVer.SetTitle("主版本")
    mw.ui.VersionTabVieConMastVer.SetWidth(60)

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
    mw.ui.VersionTabVieConTime.SetWidth(150)
    mw.ui.VersionTabVieConTime.SetFormat(layoutTime)

    // TableViewColumn 添加到 TableView
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConIndex)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConMastVer)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConVer)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConPack)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTag)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTagPath)
    mw.ui.VersionTableView.Columns().Add(mw.ui.VersionTabVieConTime)

    succeeded = true

    return nil
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
