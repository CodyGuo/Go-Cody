// iMan 升级程序界面版本

package main

import (
    "log"
    "regexp"
)

import (
    "github.com/lxn/walk"
    // "github.com/lxn/win"
)

const (
    ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

func main() {
    if _, err := RunMyDialog(nil); err != nil {
        log.Fatal(err)
    }
}

func RunMyDialog(owner walk.Form) (int, error) {
    dlg := new(MyDialog)
    err := dlg.init(owner)
    if err != nil {
        return 0, err
    }

    dlg.SetScreenCenter(true)

    // 增加托盘图标
    err = dlg.setMyNotify()
    dlg.checkError(err)
    defer dlg.ni.Dispose()

    // 托盘图标默认隐藏.
    err = dlg.ni.SetVisible(true)
    dlg.checkError(err)

    // 增加托盘图标退出菜单
    err = dlg.addMyNotifyAction()
    dlg.checkError(err)

    // 退出时隐藏到托盘图标
    err = dlg.setExitHide(true)
    dlg.checkError(err)

    // 鼠标左键显示主窗体
    dlg.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
        if button == walk.LeftButton {
            dlg.Show()
            dlg.BringToTop() // 窗口前置
            // ni.SetVisible(false)

            dlg.ni.ShowInfo("Hello", "显示主窗体成功.")
        }
    })

    // 浏览按钮
    dlg.ui.browseBtn.Clicked().Attach(func() {
        dlg.uploadFile()
        dlg.ui.fileLe.SetText(dlg.upgrideFile)
    })

    // 上传按钮
    dlg.ui.uploadBtn.Clicked().Attach(func() {
        dlg.ui.lv.Clean()
        // 判断输入的是否为正确的IP地址
        ipInputTrue, _ := regexp.MatchString(ipRegxp, dlg.ui.ipLe.Text())
        switch {
        case dlg.upgrideFile == "":
            dlg.myMsg("警告信息", "请选择上传文件...", walk.MsgBoxIconWarning)
        case dlg.ui.ipLe.Text() == "" || !ipInputTrue:
            dlg.myMsg("警告信息", "请输入正确的服务器IP...--> "+dlg.ui.ipLe.Text(), walk.MsgBoxIconWarning)
        default:
            dlg.ui.lv.AppendText("日志打印开始...\n")
            dlg.myMsg("提示信息", "服务器 ["+dlg.ui.ipLe.Text()+"] 上传文件成功...\n"+dlg.upgrideFile, walk.MsgBoxIconInformation)
            dlg.ui.lv.PostAppendText("日志打印结束...\n")
            dlg.ui.fileLe.SetText("")
            dlg.upgrideFile = ""
        }
    })

    return dlg.Run(), nil

}
