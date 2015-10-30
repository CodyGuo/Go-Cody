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

func init() {
    iManPid := fmt.Sprint(os.Getpid())
    tmpDir := os.TempDir()

    if err := ProcEsxit(tmpDir); err == nil {
        pidFile, _ := os.Create(tmpDir + "\\imanPack.pid")
        defer pidFile.Close()

        pidFile.WriteString(iManPid)
    } else {
        os.Exit(1)
    }
}

func main() {
    err := RunMyWindow()
    if err != nil {
        log.Fatal(err)
    }
}

// 主窗体
func RunMyWindow() (err error) {
    mw := new(MyWindow)
    mw.init()

    // 居中
    mw.SetScreenCenter(true)

    // 设置主窗体在所有窗体之前
    mw.SetForegroundWindow()
    mw.SwitchToThisWindow(true)

    // 菜单--服务器设置
    mw.ui.ServerAction.Triggered().Attach(func() {
        if err := RunSetServer(mw, mw); err != nil {
            log.Fatal(err)
        }
    })

    // 菜单--关于
    mw.ui.AboutAction.Triggered().Attach(func() {
        mw.MyMsg("关于", "此程序正在开发阶段.", walk.MsgBoxIconInformation)
    })

    // 增加托盘图标
    err = mw.SetMyNotify()
    mw.checkError(err)
    defer mw.ni.Dispose()

    // 托盘图标默认隐藏.
    err = mw.ni.SetVisible(true)
    mw.checkError(err)

    // 增加托盘图标退出菜单
    err = mw.AddMyNotifyAction()
    mw.checkError(err)

    // 退出时隐藏到托盘图标
    err = mw.SetExitHide(false)
    mw.checkError(err)

    // 最小化时隐藏窗体到托盘图标
    mw.SizeChanged().Attach(func() {
        if mw.X() == -32000 && mw.Y() == -32000 {
            mw.Hide()
        }
    })

    // 鼠标左键显示主窗体
    mw.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
        if button == walk.LeftButton {
            mw.SwitchToThisWindow(true)
            mw.SetForegroundWindow()
        }
    })

    // 开始打包
    mw.ui.StartPackingBtn.Clicked().Attach(func() {
        mw.MyMsg("打包程序", "即将实现打包...", walk.MsgBoxIconInformation)
    })

    // 测试版本控制
    mw.DisablePackGb(true)
    mw.CheckAll(true)
    mw.ui.VersionTestRadio.CheckedChanged().Attach(func() {
        if mw.ui.VersionTestRadio.Checked() {
            mw.DisablePackGb(true)
        }
    })

    // 正式版本控制
    mw.ui.VersionOffRadio.CheckedChanged().Attach(func() {
        if mw.ui.VersionOffRadio.Checked() {
            mw.DisablePackGb(false)
        }
    })

    // 全选控制
    mw.ui.CheckAllCb.CheckedChanged().Attach(func() {
        if mw.ui.CheckAllCb.Checked() {
            mw.CheckAll(true)
        } else {
            mw.CheckAll(false)
        }

    })

    // PC上传按钮
    mw.ui.PcUploadBtn.Clicked().Attach(func() {
        pcFile := mw.UploadFile("PC助手")
        mw.ui.PcHelperLe.SetText(pcFile)
    })

    // Android 助手
    mw.ui.AndUploadBtn.Clicked().Attach(func() {
        androidFile := mw.UploadFile("Android助手")
        mw.ui.AndroidHelperLe.SetText(androidFile)
    })

    // Web 数据库
    mw.ui.WebSqlBtn.Clicked().Attach(func() {
        webSqlFile := mw.UploadFile("Web 数据库")
        mw.ui.WebSqlLe.SetText(webSqlFile)
    })

    mw.Show()

    ok := mw.Run()
    if ok != 0 {
        return errors.New("[ERROR] 运行主窗体错误.")
    }

    return nil
}

// 设置服务器窗体
func RunSetServer(owner walk.Form, mw *MyWindow) (err error) {
    dlg := new(DlgServer)

    dlg.init(owner)
    dlg.ui.IpLe.SetText(ConfSer.Ip)

    mw.ui.BuildServerLb.SetText("编译服务器IP: " + dlg.ui.IpLe.Text())

    ConfVerModel.Read()
    mw.ui.VersionTableView.SetModel(ConfVerModel)

    // 确定
    dlg.ui.AcceptPB.Clicked().Attach(func() {
        ipInputTrue, _ := regexp.MatchString(ipRegxp, dlg.ui.IpLe.Text())
        switch {
        case dlg.ui.IpLe.Text() == "" || !ipInputTrue:
            log.Println("[WARN] 请输入正确的服务器IP【 " + dlg.ui.IpLe.Text() + " 】.")
            walk.MsgBox(dlg, "警告信息", "请输入正确的服务器IP【 "+dlg.ui.IpLe.Text()+
                " 】.", walk.MsgBoxIconWarning)
        case dlg.ui.UserLe.Text() == "":
            log.Println("[WARN] 请输入用户名!")
            walk.MsgBox(dlg, "警告信息", "请输入用户名!", walk.MsgBoxIconWarning)
        case dlg.ui.PasswdLe.Text() == "":
            log.Println("[WARN] 请输入密码!")
            walk.MsgBox(dlg, "警告信息", "请输入密码!", walk.MsgBoxIconWarning)
        default:
            log.Println("[INFO] 您设置的服务器IP为【 " + dlg.ui.IpLe.Text() + " 】." +
                "用户名为 【 " + dlg.ui.UserLe.Text() + " 】." +
                "密码为【 " + dlg.ui.PasswdLe.Text() + " 】.")
            walk.MsgBox(dlg, "提示信息", "您设置的服务器IP为【 "+dlg.ui.IpLe.Text()+
                " 】.", walk.MsgBoxIconInformation)

            ConfSer.Ip = dlg.ui.IpLe.Text()
            ConfSer.User = dlg.ui.UserLe.Text()
            ConfSer.Passwd = dlg.ui.PasswdLe.Text()
            mw.ui.BuildServerLb.SetText("编译服务器IP: " + dlg.ui.IpLe.Text())
            // fmt.Println("iman:", ConfSer.Ip, ConfSer.User, ConfSer.Passwd)
            // 写入数据库
            ConfSer.Write()
            dlg.Close(0)
        }
    })

    ok := dlg.Run()
    if ok != 0 {
        return errors.New("[ERROR] 运行服务器设置窗体错误.")
    }

    return nil
}

// 判断进程是否启动
func ProcEsxit(tmpDir string) (err error) {
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
