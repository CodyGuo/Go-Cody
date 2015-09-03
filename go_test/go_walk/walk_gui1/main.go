package main

import (
    "fmt"
    "log"
    // "strconv"
    "github.com/codyguo/skinh"
    "strings"
    "time"

    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func main() {
    skinh.Attach()

    Input := new(Model)
    mw := new(MyMainWindow)

    date, _ := time.Parse("2006-02-03 03:04:05", "2016-01-01 08:08:08")

    if err := (MainWindow{
        AssignTo: &mw.MainWindow,
        DataBinder: DataBinder{
            AssignTo:       &db,
            DataSource:     Input,
            ErrorPresenter: ErrorPresenterRef{&ep},
        },
        Title:   "iMan - 测试程序",
        MinSize: Size{1100, 700},
        Layout:  VBox{},
        Children: []Widget{
            Composite{
                AssignTo: &setting,
                Layout:   Grid{Columns: 1},
                Children: []Widget{
                    RadioButtonGroupBox{
                        AssignTo:   &radioBtnGpBox,
                        Title:      "注册模式",
                        MinSize:    Size{0, 30},
                        Layout:     HBox{},
                        DataMember: "Int",
                        Buttons: []RadioButton{
                            {Text: iMode.RunMode[0].Key, Value: iMode.RunMode[0].Int},
                            {Text: iMode.RunMode[1].Key, Value: iMode.RunMode[1].Int},
                            {Text: iMode.RunMode[2].Key, Value: iMode.RunMode[2].Int},
                        },
                    },
                },
            },
            Composite{
                AssignTo: &setting,
                Layout:   Grid{Columns: 2},
                Children: []Widget{
                    CheckBox{
                        AssignTo:   &cBox1,
                        Text:       "盒子",
                        CheckState: walk.CheckUnchecked,
                        Tristate:   true,
                        OnClicked: func() {
                            fmt.Println(cBox1.CheckState())
                            if cBox1.CheckState() == 0 {
                                mw.openAddd_Triggered("您关闭了盒子.")
                            } else if cBox1.CheckState() == 1 {
                                mw.openAddd_Triggered("您打开了盒子.")
                            } else {
                                mw.openAddd_Triggered("您的盒子是未知状态.")
                            }
                        },
                        OnCheckedChanged: func() {
                            mw.openAddd_Triggered("您修改了盒子的状态.")
                        },
                        OnCheckStateChanged: func() {
                            mw.openAddd_Triggered("您的盒子状态被改变了...")
                        },
                        // CheckState: false,
                    },

                    DateEdit{
                        AssignTo: &de1,
                        // Format:   "Y-M-d H:m:s",
                        Format:     "yyyy-MM-dd H:mm:ss",
                        NoneOption: true,
                        MinDate:    time.Now(),
                        MaxDate:    time.Now().AddDate(1, 0, 0),
                        Date:       date,
                    },
                    PushButton{
                        AssignTo: &addIpBtn,
                        Text:     "添加",
                        OnClicked: func() {
                            ok := mw.addAction_ServerIp()
                            if ok == 0 {
                                fmt.Println("添加服务器", "添加"+iplist.Ip+"成功.")
                            }

                            iplist.Ip = ""
                            iplist.Remark = ""
                        },
                    },

                    TableView{
                        MinSize:                    Size{200, 100},
                        MaxSize:                    Size{550, 450},
                        AlternatingRowBGColor:      walk.RGB(255, 255, 224),
                        CheckBoxes:                 true,
                        ItemStateChangedEventDelay: 1,
                        ColumnsOrderable:           false,
                        ColumnsSizable:             false,
                        SingleItemSelection:        false, // 单选
                        LastColumnStretched:        false, // 显示最后一列
                        Columns: []TableViewColumn{
                            {Title: "序号"},
                            {Title: "IP"},
                            {Title: "备注"},
                            {Title: "添加时间", Format: "2006-01-02 15:04:05", Width: 150},
                        },

                        Model: &model,
                    },

                    PushButton{
                        AssignTo: &addIpBtn,
                        Text:     "提交",
                        MinSize:  Size{60, 80},
                        MaxSize:  Size{100, 80},

                        ToolTipText: "提示信息，请填写数字",
                        OnClicked: func() {

                            fmt.Println("-----ok----------")
                            InputIp.ServerLists = model.GetChecked()
                            var tmpListIp []string
                            if len(InputIp.ServerLists) > 0 {
                                fmt.Println("--------------------选中了哪个-----------")
                                // fmt.Print(InputIp.ServerLists[0].Ip)
                                for index, _ := range InputIp.ServerLists {
                                    tmpListIp = append(tmpListIp, InputIp.ServerLists[index].Ip)
                                    // if len(InputIp.ServerLists) == 1 {
                                    //     tmpListIp = InputIp.ServerLists[index].Ip
                                    // } else if len(InputIp.ServerLists) == 2 {
                                    //     tmpListIp = tmpListIp + " , " + InputIp.ServerLists[index].Ip
                                    // } else {
                                    //     tmpListIp = tmpListIp + " , " + InputIp.ServerLists[index].Ip
                                    // }

                                }
                                fmt.Println(tmpListIp)
                                mw.openAddd_Triggered(strings.Join(tmpListIp, ", "))
                            }
                        },
                    },
                },
            },
        },
    }).Create(); err != nil {
        // log.Fatal(err)
        fmt.Println("错误来了：", err)
        log.Println(err)
    }

    // addIpBtn.SetMinMaxSize(walk.Size{30, 50}, walk.Size{50, 60})
    // mw.SetPersistent(true)
    mw.Run()
}

func (mw *MyMainWindow) openAddd_Triggered(message string) {
    walk.MsgBox(mw, "提示信息", message, walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) addAction_ServerIp() int {

    cmd, err := RunServerListDialog(mw, &iplist)

    switch {
    case err != nil:
        log.Print(err)
        return 1
    case cmd == walk.DlgCmdOK:
        fmt.Println(cmd)
        fmt.Println("----------添加-----------")
        // fmt.Println(&iplist.Ip, &iplist.Remark)
        fmt.Println(iplist)
        model.AddServerIp(iplist.Ip, iplist.Remark)
        fmt.Println("----------添加-END-----------")
    default:
        fmt.Println(cmd, err)
        return 1
    }

    return 0

}

// func (mw *MyMainWindow) addServerStart() {
//     if err := db.Submit(); err != nil {
//         // outTE.SetText("请选择一个模式...")
//         mw.openAddd_Triggered("请选择一个模式...")
//         // return
//     }
//     fmt.Println("-----ok----------")

//     Input.ServerLists = model.GetChecked()

//     if len(Input.ServerLists) > 0 {
//         fmt.Println("--------------------选中了哪个-----------")
//         // fmt.Print(Input.ServerLists[0].Ip)
//         // for index, _ := range Input.ServerLists {
//         //     tmpListIp = strings.Join(string(Input.ServerLists[index].Ip), ",")
//         // }
//     }
//     // mw.openAddd_Triggered(tmpListIp)
//     // if Input.Int == 1 {
//     //     // outTE.SetText("注册为控制器.")

//     //     // lineE1.SetText(strconv.FormatFloat(numB1.Value(), 'f', 2, 32))
//     //     mw.openAddd_Triggered("控制器注册成功.")
//     //     fmt.Println("注册为控制器.")
//     // } else if Input.Int == 2 {
//     //     // outTE.SetText("注册为服务器.")
//     //     mw.openAddd_Triggered("服务器注册成功.")
//     //     fmt.Println("注册为服务器")
//     //     // fmt.Println(numB1.Font())
//     // } else if Input.Int == 3 {
//     //     // inTE.SetText("注册为控制器+服务器.")
//     //     fmt.Println("注册为控制器+服务器")
//     //     mw.openAddd_Triggered("控制器+服务器注册成功.")
//     // }

//     // 读取任务
//     // Input.ServerLists = model.GetChecked()

//     // if len(Input.ServerLists) == 0 {
//     //     log.Println(" *     —— 亲，任务列表不能为空哦~")
//     //     return
//     // }

//     // fmt.Print(&Input.ServerLists)
//     // mw.openAddd_Triggered(Input.ServerLists)

// }
