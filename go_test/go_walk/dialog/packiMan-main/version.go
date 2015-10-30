package main

import (
    "errors"
    "fmt"
    "time"
)

import (
    "github.com/lxn/walk"
)

const (
    layoutTime = "2006-01-02 15:04:05"
)

var (
    ConfVerModel *ConfigVersionModel
)

func init() {
    ConfVerModel.Read()
}

type ConfigVersion struct {
    Index     int
    Version   string
    MasterVer string
    Pack      string
    Tag       string
    TagPath   string
    PackTime  time.Time
    Checked   bool
}

type ConfigVersionModel struct {
    walk.TableModelBase
    walk.SorterBase
    sortColumn int
    sortOrder  walk.SortOrder
    items      []*ConfigVersion
}

func NewConfgVersion() *ConfigVersion {
    return new(ConfigVersion)
}

func NewConfigVersionModel() *ConfigVersionModel {
    return new(ConfigVersionModel)
}

func (cm *ConfigVersionModel) RowCount() int {
    return len(cm.items)
}

func (cm *ConfigVersionModel) Checked(row int) bool {
    return cm.items[row].Checked
}

func (cm *ConfigVersionModel) SetChecked(row int, checked bool) error {
    cm.items[row].Checked = checked
    return nil
}

func (cm *ConfigVersionModel) Value(row, col int) interface{} {
    item := cm.items[row]

    switch col {
    case 0:
        return item.Index
    case 1:
        return item.MasterVer
    case 2:
        return item.Version
    case 3:
        return item.Pack
    case 4:
        return item.Tag
    case 5:
        return item.TagPath
    case 6:
        return item.PackTime
    }

    panic("unexpected col")
}

// func (cm *ConfigVersionModel) Sort(col int, order walk.SortOrder) error {
//     cm.sortColumn, cm.sortOrder = col, order

//     sort.Stable(m)

//     return cm.SorterBase.Sort(col, order)
// }

func (cm *ConfigVersionModel) Len() int {
    return len(cm.items)
}

func (cm *ConfigVersionModel) Less(i, j int) bool {
    a, b := cm.items[i], cm.items[j]
    c := func(ls bool) bool {
        if cm.sortOrder == walk.SortAscending {
            return ls
        }
        return !ls
    }

    switch cm.sortColumn {
    case 0:
        return c(a.Index < b.Index)
    case 1:
        return c(a.Version < b.Version)
    case 2:
        return c(a.MasterVer < b.MasterVer)
    case 3:
        return c(a.Pack < b.Pack)
    case 4:
        return c(a.Tag < b.Tag)
    case 5:
        return c(a.TagPath < b.TagPath)
    case 6:
        return c(a.PackTime.Before(b.PackTime))
    }

    panic("unreachable")
}

func (cm *ConfigVersionModel) Swap(i, j int) {
    cm.items[i], cm.items[j] = cm.items[j], cm.items[i]
}

func (cm *ConfigVersionModel) ResetRows(conf *ConfigVersion) {
    cm.items = append(cm.items, &ConfigVersion{
        Index:     conf.Index,
        MasterVer: conf.MasterVer,
        Version:   conf.Version,
        Pack:      conf.Pack,
        Tag:       conf.Tag,
        TagPath:   conf.TagPath,
        PackTime:  conf.PackTime,
    })

    // fmt.Println("ResetRows:", conf.Index, conf.MasterVer, conf.Version, conf.Pack, conf.Tag, conf.TagPath, conf.PackTime)

    // Notify TableView and other interested parties about the reset.
    cm.PublishRowsReset()

    cm.Sort(cm.sortColumn, cm.sortOrder)
}

func (cm *ConfigVersionModel) Read() (err error) {
    db, err := openDb()
    defer db.Close()
    if err != nil {
        return err
    }

    rows, err := db.Query("select * from version")
    defer rows.Close()
    if err != nil {
        return err
    }

    ConfVerModel = NewConfigVersionModel()

    for {
        ConfVer := NewConfgVersion()
        var packtime string

        switch rows.Next() {
        case false:
            fmt.Println("历史记录查询完毕.")
            return errors.New("历史记录查询完毕.")
        case true:
            rows.Scan(&ConfVer.Index, &ConfVer.MasterVer, &ConfVer.Version, &ConfVer.Pack, &ConfVer.Tag, &ConfVer.TagPath, &packtime)
            ConfVer.PackTime, _ = time.Parse(layoutTime, packtime)
            ConfVerModel.ResetRows(ConfVer)
            fmt.Println("version - read:", ConfVer.Index, ConfVer.MasterVer, ConfVer.Version, ConfVer.Pack, ConfVer.Tag, ConfVer.TagPath, ConfVer.PackTime)
        }
    }

    db.Close()

    return nil
}
