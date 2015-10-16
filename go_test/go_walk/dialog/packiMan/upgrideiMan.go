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

    return dlg.Run(), nil
}
