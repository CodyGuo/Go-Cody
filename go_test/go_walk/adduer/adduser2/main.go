package main

import (
    "encoding/gob"
    "fmt"
    "log"
    "os"

    // "strconv"
    "github.com/codyguo/skinh"

    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func (u *User) LoadConfig() {
    file, err := os.Open(configFile)
    if err != nil {
        log.Println("打开配置文件失败." + err.Error())
        UserConfig.UserName = "用户名"
        UserConfig.Password = "密码"
        UserConfig.JobNumber = 1
        UserConfig.Name = "姓名"
        UserConfig.Email = "邮箱"
        UserConfig.Phone = "联系电话"
        return
    }
    defer file.Close()

    dec := gob.NewDecoder(file)
    err = dec.Decode(UserConfig)
    if err != nil {
        UserConfig.UserName = "用户名"
        UserConfig.Password = "密码"
        UserConfig.JobNumber = 1
        UserConfig.Name = "姓名"
        UserConfig.Email = "邮箱"
        UserConfig.Phone = "联系电话"
    }
}

func (u *User) SaveConfig() {
    file, err := os.Create(configFile)
    if err != nil {
        log.Fatal("创建配置文件失败." + err.Error())
    }
    defer file.Close()

    enc := gob.NewEncoder(file)
    err = enc.Encode(UserConfig)
    if err != nil {
        log.Fatal("保存配置文件失败." + err.Error())
    }

}

func init() {
    UserConfig.LoadConfig()
}

func main() {
    skinh.Attach()
    mw := new(MyMainWindow)

    // 禁止最大化
    mw.SetMaximizeBox(false)

    // 禁止最小化
    mw.SetMinimizeBox(false)

    // 固定窗体大小
    mw.SetFixedSize(true)

    if err := (MainWindow{
        AssignTo: &mw.MainWindow,
        DataBinder: DataBinder{
            AssignTo:       &db,
            DataSource:     UserConfig,
            ErrorPresenter: ErrorPresenterRef{&ep},
        },
        Title:        "Cody - 测试程序",
        ScreenCenter: true, // 屏幕居中
        MinSize:      Size{300, 262},
        Layout:       VBox{Spacing: 2},
        Children: []Widget{
            Composite{
                Layout: Grid{Columns: 2, Spacing: 10},
                Children: []Widget{
                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "用户名:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{

                            LineEdit{
                                AssignTo: &ltu,
                                MinSize:  Size{160, 0},
                                Text:     Bind("UserName", SelRequired{}),
                            },
                        },
                    },

                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "密码:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{

                            LineEdit{
                                AssignTo:     &ltu,
                                PasswordMode: true,
                                MinSize:      Size{160, 0},
                                Text:         Bind("Password"),
                            },
                        },
                    },

                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "姓名:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{
                            LineEdit{
                                AssignTo: &ltu,
                                MinSize:  Size{160, 0},
                                Text:     Bind("Name"),
                            },
                        },
                    },

                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "工号:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{
                            NumberEdit{
                                AssignTo: &ltj,
                                Value:    Bind("JobNumber", Range{1, 10}),
                            },
                        },
                    },

                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "邮箱:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{
                            LineEdit{
                                AssignTo: &lte,
                                MinSize:  Size{160, 0},
                                Text:     Bind("Email"),
                            },
                        },
                    },

                    VSplitter{
                        Children: []Widget{
                            Label{
                                MinSize: Size{18, 0},
                                Text:    "手机:",
                            },
                        },
                    },
                    VSplitter{
                        Children: []Widget{

                            LineEdit{
                                AssignTo: &ltphone,
                                MinSize:  Size{160, 0},
                                Text:     Bind("Phone"),
                            },
                        },
                    },

                    // VSplitter{},
                    VSplitter{
                        ColumnSpan: 2,
                        MinSize:    Size{100, 0},
                        Children: []Widget{
                            LineErrorPresenter{
                                AssignTo:   &ep,
                                ColumnSpan: 2,
                            },

                            PushButton{
                                AssignTo:  &addUserBtn,
                                MinSize:   Size{90, 0},
                                Text:      "提交",
                                OnClicked: mw.openAddd_Triggered,
                            },
                        },
                    },
                },
            },
        },
    }.Create()); err != nil {
        // log.Fatal(err)
        fmt.Println("错误来了：", err)
        log.Println(err)
    }
    // addIpBtn.SetMinMaxSize(walk.Size{18, 18}, walk.Size{18, 18})

    mw.Run()
    mw.Dispose()
}

func (mw *MyMainWindow) openAddd_Triggered() {
    fmt.Println(mw.Size())
    if err := db.Submit(); err != nil {
        log.Println(err)
        walk.MsgBox(mw, "错误提示", err.Error(), walk.MsgBoxIconError)

        return
    }
    UserConfig.SaveConfig()
    walk.MsgBox(mw, "提示信息", "保存用户成功.", walk.MsgBoxIconInformation)

}
