// iMan 升级程序界面版本

package main

import (
    "log"
)

import (
    "github.com/lxn/walk"
)

func main() {
    MS.LoadData()

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
    // 选择附件
    dlg.ui.browseBtn.Clicked().Attach(func() {
        dlg.browseFile()
        dlg.ui.mailAdjLe.SetText(dlg.adjFile)
    })

    // 开始发送
    dlg.ui.sendBtn.Clicked().Attach(func() {
        // 初始化数据
        MS.SendList = append(MS.SendList, dlg.ui.mailListLe.Text())
        MS.UserName = dlg.ui.userNameLe.Text()
        MS.Passwd = dlg.ui.passwdLe.Text()
        MS.Smtp = dlg.ui.smtpLe.Text()
        MS.Port = dlg.ui.portLe.Text()
        MS.Subject = dlg.ui.mailSubLe.Text()
        MS.Body = dlg.ui.mailBodyLe.Text()
        MS.Adjunct = dlg.ui.mailAdjLe.Text()

        err = MS.SendMail()
        if err != nil {
            dlg.myMsg("错误信息", "邮件群发失败."+err.Error(), walk.MsgBoxIconError)
        } else {
            dlg.myMsg("提示信息", "邮件群发成功.", walk.MsgBoxIconInformation)
            MS.SaveData()
        }

    })

    return dlg.Run(), nil

}
