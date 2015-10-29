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

    mw.ui.VersionTableView.SetModel(ConfVerModel)

    // 刷新数据库数据到tableview
    mw.ui.PushButton1.Clicked().Attach(func() {
        ConfVerModel.Read()
        // 重置tableview
        mw.ui.VersionTableView.SetModel(ConfVerModel)
    })

    mw.ui.PushButton2.Clicked().Attach(func() {
        // 获取选择的项
        var getCv []*ConfigVersion
        getCv = ConfVerModel.GetChecked()
        fmt.Println("-----开始选择getcv-----")
        if len(getCv) > 0 {
            for i, _ := range getCv {
                fmt.Println("我选中了:", getCv[i].TagPath)
            }
        }
        fmt.Println("-----选择结束-----")
    })

    mw.Show()

    ok := mw.Run()
    if ok != 0 {
        return errors.New("[ERROR] 运行主窗体错误.")
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
