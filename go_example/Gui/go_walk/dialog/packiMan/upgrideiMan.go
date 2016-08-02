// iMan 升级程序界面版本

package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strconv"
)

import (
    "github.com/lxn/walk"
)

const (
    ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

func ProcExsit(tmpDir string) (err error) {
    iManPidFile, err := os.Open(tmpDir + "\\imanPack.pid")
    defer iManPidFile.Close()

    if err == nil {
        filePid, err := ioutil.ReadAll(iManPidFile)
        if err == nil {
            pidStr := fmt.Sprintf("%s", filePid)
            pid, _ := strconv.Atoi(pidStr)
            _, err := os.FindProcess(pid)
            if err == nil {
                return errors.New("[ERROR] iMan升级工具已启动.")
            }
        }
    }

    return nil
}

func init() {
    iManPid := fmt.Sprint(os.Getpid())
    tmpDir := os.TempDir()

    if err := ProcExsit(tmpDir); err == nil {
        pidFile, _ := os.Create(tmpDir + "\\imanPack.pid")
        defer pidFile.Close()

        pidFile.WriteString(iManPid)
    } else {
        os.Exit(1)
    }
}

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

    // 居中
    dlg.SetScreenCenter(true)

    // 设置主窗体在所有窗体之前
    dlg.SetForegroundWindow()
    dlg.SwitchToThisWindow(true)

    // // 设置
    // dlg.ui.SettingMenu, _ = walk.NewMenu()
    // dlg.ui.SettingAction = walk.NewMenuAction(dlg.ui.SettingMenu)
    // dlg.ui.SettingAction.SetText("设置")

    // dlg.ui.ServerAction = walk.NewAction()
    // dlg.ui.ServerAction.SetText("服务器")

    // dlg.ui.SettingMenu.Actions().Add(dlg.ui.ServerAction)

    // // 帮助
    // dlg.ui.HelpMenu, _ = walk.NewMenu()
    // dlg.ui.HelpAction = walk.NewMenuAction(dlg.ui.HelpMenu)
    // dlg.ui.HelpAction.SetText("帮助")

    // dlg.ui.AboutAction = walk.NewAction()
    // dlg.ui.AboutAction.SetText("关于")

    // dlg.ui.HelpMenu.Actions().Add(dlg.ui.AboutAction)

    // // 菜单配置
    // dlg.Menu().Actions().Add(dlg.ui.SettingMenu)

    // 增加托盘图标
    err = dlg.SetMyNotify()
    dlg.checkError(err)
    defer dlg.ni.Dispose()

    // 托盘图标默认隐藏.
    err = dlg.ni.SetVisible(true)
    dlg.checkError(err)

    // 增加托盘图标退出菜单
    err = dlg.AddMyNotifyAction()
    dlg.checkError(err)

    // 退出时隐藏到托盘图标
    err = dlg.SetExitHide(false)
    dlg.checkError(err)

    // 最小化时隐藏窗体到托盘图标
    dlg.SizeChanged().Attach(func() {
        if dlg.X() == -32000 && dlg.Y() == -32000 {
            dlg.Hide()
        }
    })

    // 鼠标左键显示主窗体
    dlg.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
        if button == walk.LeftButton {
            dlg.SwitchToThisWindow(true)
            dlg.SetForegroundWindow()
        }
    })

    // 开始打包
    dlg.ui.StartPackingBtn.Clicked().Attach(func() {
        dlg.MyMsg("打包程序", "即将实现打包...", walk.MsgBoxIconInformation)
    })

    // 测试版本控制
    dlg.DisablePackGb(true)
    dlg.ui.VersionTestRadio.CheckedChanged().Attach(func() {
        if dlg.ui.VersionTestRadio.Checked() {
            dlg.DisablePackGb(true)
        }
    })

    // 正式版本控制
    dlg.ui.VersionOffRadio.CheckedChanged().Attach(func() {
        if dlg.ui.VersionOffRadio.Checked() {
            dlg.DisablePackGb(false)
        }
    })

    // 全选控制
    dlg.ui.CheckAllCb.CheckedChanged().Attach(func() {
        if dlg.ui.CheckAllCb.Checked() {
            dlg.CheckAll(true)
        } else {
            dlg.CheckAll(false)
        }

    })

    // PC上传按钮
    dlg.ui.PcUploadBtn.Clicked().Attach(func() {
        pcFile := dlg.UploadFile("PC助手")
        dlg.ui.PcHelperLe.SetText(pcFile)
    })

    // Android 助手
    dlg.ui.AndUploadBtn.Clicked().Attach(func() {
        androidFile := dlg.UploadFile("Android助手")
        dlg.ui.AndroidHelperLe.SetText(androidFile)
    })

    // Web 数据库
    dlg.ui.WebSqlBtn.Clicked().Attach(func() {
        webSqlFile := dlg.UploadFile("Web 数据库")
        dlg.ui.WebSqlLe.SetText(webSqlFile)
    })

    return dlg.Run(), nil
}
