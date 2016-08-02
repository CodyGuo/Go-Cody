package main

import (
    "errors"
    "fmt"
    "log"
)

import (
    "github.com/lxn/walk"
)

func main() {
    if err := RunMaster(); err != nil {
        log.Fatalln(err)
    }

}

func RunMaster() (err error) {
    mw := new(MyWindow)
    err = mw.init()
    if err != nil {
        return err
    }

    mw.SetScreenCenter(true)
    mw.SetForegroundWindow()
    mw.SwitchToThisWindow(true)

    var runNum int
    mw.ui.CreateBtn.Clicked().Attach(func() {
        if mw.ui.SecTitileLe.Text() != "" {
            runNum += 1
            if runNum > 1 {
                // 只允许打开一个子窗体
                dlg.Dispose()
            }
            RunChild(mw, mw.ui.SecTitileLe.Text())
        } else {
            walk.MsgBox(mw, "提示", "请填写子窗体标题.", walk.MsgBoxIconInformation)
        }
    })

    // 主窗体关闭时,关闭子窗体
    mw.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
        if reason == walk.CloseReasonUnknown {
            if runNum >= 1 {
                dlg.Close(0)
            }
        } else {
            defer dlg.Close(1)
        }
    })

    mw.Show()
    fmt.Println("主窗体运行中...--> 【 " + mw.Title() + " 】.")
    ok := mw.Run()
    if ok != 0 {
        return errors.New("运行主窗体错误.")
    }

    return nil
}

func RunChild(owner walk.Form, title string) (err error) {
    dlg = new(DlgChild)
    err = dlg.init(owner, title)
    if err != nil {
        return err
    }
    dlg.SetScreenCenter(true)

    dlg.Show()
    fmt.Println("子窗体运行中...--> 【 " + dlg.Title() + " 】.")

    ok := dlg.Run()
    if ok != 0 {
        return errors.New("运行子窗体错误.")
    }

    return nil
}
