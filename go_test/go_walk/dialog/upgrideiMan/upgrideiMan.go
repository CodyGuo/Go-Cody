// iMan 升级程序界面版本

package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "regexp"
    "strconv"
)

import (
    "github.com/lxn/walk"
)

const (
    ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

func ProcEsxit(tmpDir string) (err error) {
    iManPidFile, err := os.Open(tmpDir + "\\iman.pid")
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

    if err := ProcEsxit(tmpDir); err == nil {
        pidFile, _ := os.Create(tmpDir + "\\iman.pid")
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
    err = dlg.setExitHide(false)
    dlg.checkError(err)

    // 最小化时隐藏窗体到托盘图标
    dlg.SizeChanged().Attach(func() {
        if dlg.X() == -32000 && dlg.Y() == -32000 {
            dlg.Hide()
        }
    })

    // 鼠标左键显示主窗体
    // var num int
    dlg.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
        if button == walk.LeftButton {
            dlg.SwitchToThisWindow(true)
            dlg.SetForegroundWindow()
            // num += 1
        }

        // 双击的老土方法
        // if num == 2 {
        //     // dlg.Show()
        //     // 最小化时，弹出主窗体
        //     dlg.SwitchToThisWindow(true)
        //     num = 0
        //     // dlg.SetForegroundWindow()

        //     // dlg.ni.ShowInfo("Hello", "显示主窗体成功.")
        // }
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
            dlg.myMsg("警告信息", "请选择上传文件!", walk.MsgBoxIconWarning)
        case dlg.ui.ipLe.Text() == "" || !ipInputTrue:
            dlg.myMsg("警告信息", "请输入正确的服务器IP【 "+dlg.ui.ipLe.Text()+" 】.", walk.MsgBoxIconWarning)
        default:
            go func() {
                dlg.ui.uploadBtn.SetEnabled(false)
                dlg.ui.browseBtn.SetEnabled(false)

                dlg.ui.lv.AppendText("-------------------------iMan 开始升级...-------------------------\n\n")
                // 升级过程中退出时隐藏到托盘图标
                err = dlg.setExitHide(true)
                dlg.checkError(err)

                ok := make(chan bool)
                go upload(dlg, ok)

                <-ok
                dlg.ui.lv.PostAppendText("\n-------------------------iMan 升级结束...-------------------------\n")
                dlg.upgrideFile = ""
                dlg.ui.fileLe.SetText("")

                dlg.ui.uploadBtn.SetEnabled(true)
                dlg.ui.browseBtn.SetEnabled(true)

                // 升级完恢复关闭退出功能
                err = dlg.setExitHide(false)
                dlg.checkError(err)
            }()

        }
    })

    return dlg.Run(), nil

}

func upload(dlg *MyDialog, ok chan bool) {
    Upload := NewUploadFile()
    Upload.Host = dlg.ui.ipLe.Text()
    Upload.User = "nacupdate"
    Upload.Passwd = "imanupdate"
    Upload.Debug = true

    Upload.LocalFilePath = dlg.ui.fileLe.Text()

    Upload.UploadFilePath = "/home/nacupdate"

    Upload.UpgradeFilePath = "/bak/upgrade"

    // 上传升级包
    err := Upload.UploadiManFile()
    if err != nil {
        log.Println(err.Error())
        log.Println("[ERROR] 服务器【" + dlg.ui.ipLe.Text() + "】升级包上传失败.")
        dlg.myMsg("错误信息", "服务器【"+dlg.ui.ipLe.Text()+"】升级包上传失败.", walk.MsgBoxIconError)
    } else {
        // 执行升级命令
        err := Upload.UpgradeiManCmd("upgrade")
        dlg.SwitchToThisWindow(true)
        dlg.SetForegroundWindow()

        if err != nil {
            log.Println(err.Error())
            dlg.myMsg("错误信息", "服务器【"+dlg.ui.ipLe.Text()+"】升级失败.\n", walk.MsgBoxIconError)
        } else {
            log.Println("[INFO] 服务器【" + dlg.ui.ipLe.Text() + "】解压系统升级包成功,请重启设备完成升级!")
            result := dlg.myMsg("提示信息", "服务器【"+dlg.ui.ipLe.Text()+"】解压系统升级包成功,请重启设备完成升级!\n确定是否要重启设备?", walk.MsgBoxOKCancel+walk.MsgBoxIconInformation)
            switch result {
            case 1:
                err := Upload.UpgradeiManCmd("reboot")
                if err != nil {
                    log.Println("[ERROR] 服务器【" + dlg.ui.ipLe.Text() + "】重启失败." + err.Error())
                    dlg.myMsg("错误信息", "服务器【"+dlg.ui.ipLe.Text()+"】重启失败.", walk.MsgBoxIconError)
                } else {
                    log.Println("[INFO] 服务器【" + dlg.ui.ipLe.Text() + "】重启成功,请稍后重连服务器.")
                    dlg.myMsg("提示信息", "服务器【"+dlg.ui.ipLe.Text()+"】重启成功,请稍后重连服务器.", walk.MsgBoxIconInformation)
                }
            case 2:
                log.Println("[INFO] 解压系统升级包成功,请手动重启服务器【" + dlg.ui.ipLe.Text() + "】完成升级.")
                dlg.myMsg("提示信息", "解压系统升级包成功,请手动重启服务器【"+dlg.ui.ipLe.Text()+"】完成升级.", walk.MsgBoxIconInformation)
            }
        }
    }

    ok <- true
}
