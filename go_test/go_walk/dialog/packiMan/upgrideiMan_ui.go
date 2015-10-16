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

type myDialogUI struct {
    // 开始打包

    StartPacking *walk.PushButton

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
    LinuxRuby       *walk.CheckBox

    // 打包日志
    PackLogLb *walk.Label
    lv        *LogView
}

type MyDialog struct {
    *walk.Dialog

    ui          myDialogUI
    upgrideFile string
    ni          *walk.NotifyIcon
}

func (mw *MyDialog) checkError(err error) {
    if err != nil {
        log.Println(err.Error())
    }
}

func (mw *MyDialog) init(owner walk.Form) (err error) {
    // 设置最小化
    mw.SetMinimizeBox(true)
    // 禁用最大化
    mw.SetMaximizeBox(false)
    // 设置窗口固定
    mw.SetFixedSize(true)
    // // 设置窗口前置
    // mw.SetWindowPos(true)

    mw.Dialog, err = walk.NewDialog(owner)
    mw.checkError(err)

    succeeded := false
    defer func() {
        if !succeeded {
            mw.Dispose()
        }
    }()

    // 设置主窗体大小
    mw.SetClientSize(walk.Size{700, 560})
    mw.checkError(err)

    // 设置主窗体标题
    mw.SetTitle("iMan-打包工具   V【" + _VERSION_ + "】")
    mw.checkError(err)

    // 设置字体和图标
    fountTitle, _ := walk.NewFont("幼圆", 10, walk.FontBold)
    otherFont, _ := walk.NewFont("幼圆", 10, 0)

    // 开始打包
    mw.ui.StartPacking, err = walk.NewPushButton(mw)
    mw.checkError(err)

    mw.ui.StartPacking.SetText("开始打包")

    mw.ui.StartPacking.SetBounds(walk.Rectangle{310, 20, 75, 30})

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

    mw.ui.PackGb.SetBounds(walk.Rectangle{355, 60, 330, 260})

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

    mw.ui.CheckAllCb.SetBounds(walk.Rectangle{80, 20, 20, 25})

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

    mw.ui.JavaWebCb.SetBounds(walk.Rectangle{80, 50, 70, 25})

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
    mw.ui.LinuxRuby, err = walk.NewCheckBox(mw.ui.PackGb)
    mw.checkError(err)

    mw.ui.LinuxRuby.SetText("Linux Ruby")

    mw.ui.LinuxRuby.SetBounds(walk.Rectangle{80, 170, 90, 25})

    // 打包日志
    mw.ui.PackLogLb, err = walk.NewLabel(mw)
    mw.checkError(err)

    mw.ui.PackLogLb.SetText("打包日志")
    mw.ui.PackLogLb.SetFont(fountTitle)

    mw.ui.PackLogLb.SetBounds(walk.Rectangle{20, 330, 60, 20})

    // 日志输出
    mw.ui.lv, err = NewLogView(mw)
    mw.checkError(err)

    err = mw.ui.lv.SetBounds(walk.Rectangle{10, 360, 680, 190})
    mw.checkError(err)

    log.SetOutput(mw.ui.lv)

    // img, _ = walk.NewIconFromResourceId(7)
    // mw.ui.browseBtn.SetImage(img)
    // mw.ui.browseBtn.SetImageAboveText(false)
    // mw.ui.browseBtn.SetBackground(bg)

    // reader, _ := os.Open("../../img/folder_add.png")
    // add, _, _ := image.Decode(reader)
    // var img walk.Image
    // img, _ = walk.NewBitmapFromImage(add)

    // mw.ui.browseBtn.SetImage(img)

    // img, _ = walk.NewImageFromFile("../../img/arrow_divide.png")
    // mw.ui.uploadBtn.SetImage(img)

    succeeded = true

    return nil
}
