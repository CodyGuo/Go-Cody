package main

import (
    "github.com/lxn/walk"
    "time"
    // . "github.com/lxn/walk/declarative"
)

var (
    addIpBtn           *walk.PushButton
    editIpBtn          *walk.PushButton
    deleteIpBtn        *walk.PushButton
    inTE, outTE        *walk.TextEdit
    tabView1           *walk.TableView
    numB1              *walk.NumberEdit
    lineE1             *walk.LineEdit
    ComBox1            *walk.ComboBox
    cBox1              *walk.CheckBox
    de1                *walk.DateEdit
    recentMenu         *walk.Menu
    openAction         *walk.Action
    radioBtnGpBox      *walk.GroupBox
    splitter           *walk.Splitter
    setting            *walk.Composite
    dlg                *walk.Dialog
    acceptPB, cancelPB *walk.PushButton
    db                 *walk.DataBinder
    ep                 walk.ErrorPresenter

    iplist    IpList
    iplistNum int

    model ServerListModel
    m     *walk.MainWindow
)

const (
    CONTROL = 1 + iota
    SERVER
    CONTROLSERVER
)

type MyMainWindow struct {
    *walk.MainWindow
}

type Model struct {
    Key string
    Int int
}

var iMode = struct {
    RunMode []*Model
}{
    RunMode: []*Model{
        {Key: "控制器", Int: CONTROL},
        {Key: "服务器", Int: SERVER},
        {Key: "控制器+服务器", Int: CONTROLSERVER},
    },
}

type IpList struct {
    Ip     string
    Remark string
}

type ServerList struct {
    Index   int
    Ip      string
    Remark  string
    AddTime time.Time
    Checked bool
}

type ServerListModel struct {
    walk.TableModelBase
    walk.SorterBase

    sortOrder  walk.SortOrder
    evenBitmap *walk.Bitmap
    sortColumn int
    items      []*ServerList
}

type Inputor struct {
    ServerLists []*ServerList
}

var InputIp = &Inputor{}
