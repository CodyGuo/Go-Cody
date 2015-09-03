package main

import (
    "fmt"
    "log"
    "sort"
    "time"

    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

func (m *ServerListModel) RowCount() int {
    return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *ServerListModel) Value(row, col int) interface{} {
    item := m.items[row]

    switch col {
    case 0:
        return item.Index

    case 1:
        return item.Ip

    case 2:
        return item.Remark

    case 3:
        return item.AddTime
    }

    panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *ServerListModel) Checked(row int) bool {
    return m.items[row].Checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *ServerListModel) SetChecked(row int, checked bool) error {
    m.items[row].Checked = checked

    return nil
}

//获取被选中的结果
func (m *ServerListModel) GetChecked() []*ServerList {
    rc := []*ServerList{}
    for idx, item := range m.items {
        if m.Checked(idx) {
            rc = append(rc, item)
        }
    }
    return rc
}

// Called by the TableView to sort the model.
func (m *ServerListModel) Sort(col int, order walk.SortOrder) error {
    m.sortColumn, m.sortOrder = col, order

    sort.Stable(m)

    return m.SorterBase.Sort(col, order)
}

func (m *ServerListModel) Len() int {
    return len(m.items)
}

func (m *ServerListModel) Less(i, j int) bool {
    a, b := m.items[i], m.items[j]

    c := func(ls bool) bool {
        if m.sortOrder == walk.SortAscending {
            return ls
        }

        return !ls
    }

    switch m.sortColumn {
    case 0:
        return c(a.Index < b.Index)

    case 1:
        return c(a.Ip < b.Ip)

    case 2:
        return c(a.Remark < b.Remark)

    case 3:
        return c(a.AddTime.Before(b.AddTime))
    }

    panic("unreachable")
}

func (m *ServerListModel) Swap(i, j int) {
    m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Called by the TableView to retrieve an item image.
func (m *ServerListModel) Image(row int) interface{} {
    if m.items[row].Index%2 == 0 {
        return m.evenBitmap
    }
    return m
}

func (m *ServerListModel) AddServerIp(ip, remark string) {
    fmt.Println("----------------IP------------------")
    fmt.Println(ip, remark)

    now := time.Now()

    m.items = append(m.items, &ServerList{
        Index:   iplistNum,
        Ip:      ip,
        Remark:  remark,
        AddTime: now,
    })

    fmt.Println("----------------IP-END------------------")
    // Notify TableView and other interested parties about the reset.
    m.PublishRowsReset()

    iplistNum += 1

}

func RunServerListDialog(owner walk.Form, iplist *IpList) (int, error) {
    return Dialog{
        AssignTo:      &dlg,
        Title:         "添加服务器",
        DefaultButton: &acceptPB,
        CancelButton:  &cancelPB,
        DataBinder: DataBinder{
            AssignTo:       &db,
            DataSource:     iplist,
            ErrorPresenter: ErrorPresenterRef{&ep},
        },
        MinSize: Size{300, 300},
        Layout:  VBox{},
        Children: []Widget{
            Composite{
                Layout: Grid{Columns: 2},
                Children: []Widget{
                    Label{
                        Text: "IP:",
                    },
                    LineEdit{
                        Text: Bind("Ip", Regexp{"^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])$"}),
                    },

                    Label{
                        Text: "备注",
                    },

                    TextEdit{
                        ColumnSpan: 2,
                        MinSize:    Size{100, 50},
                        Text:       Bind("Remark", ValidatorRef{}),
                    },

                    VSpacer{
                        ColumnSpan: 2,
                        Size:       8,
                    },

                    LineErrorPresenter{
                        AssignTo:   &ep,
                        ColumnSpan: 2,
                    },
                },
            },
            Composite{
                Layout: HBox{},
                Children: []Widget{
                    HSpacer{},
                    PushButton{
                        AssignTo: &acceptPB,
                        Text:     "确定",
                        OnClicked: func() {
                            if err := db.Submit(); err != nil {
                                log.Print(err)
                                return
                            }
                            fmt.Println("我草，又来一次。")

                            fmt.Println("===========db.submit===========")
                            fmt.Println(db.Submit())
                            fmt.Println("-----------确定-------------")
                            fmt.Println(iplist.Ip, iplist.Remark)
                            fmt.Println("-----------确定--END-------------")

                            dlg.Accept()
                        },
                    },
                    PushButton{
                        AssignTo:  &cancelPB,
                        Text:      "取消",
                        OnClicked: func() { dlg.Cancel() },
                    },
                },
            },
        },
    }.Run(owner)
}
