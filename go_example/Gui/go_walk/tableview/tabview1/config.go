package main

import (
    "errors"
    "fmt"
    // "math/rand"
    "os"
    // "sort"
    "time"

    "database/sql"
)

import (
    "github.com/lxn/walk"

    _ "github.com/mattn/go-sqlite3"
)

const (
    dbFile     = "./iman.db"
    layoutTime = "2006-01-02 15:04:05"
    ipRegxp    = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

var (
    _VERSION_ = "cody.guo"

    SqlStmt      *sql.Stmt
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
    checked   bool
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

// Called by the TableView to retrieve if a given row is checked.
func (m *ConfigVersionModel) Checked(row int) bool {
    return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *ConfigVersionModel) SetChecked(row int, checked bool) error {
    m.items[row].checked = checked

    return nil
}

//获取被选中的结果
func (m *ConfigVersionModel) GetChecked() (cv []*ConfigVersion) {
    for idx, item := range m.items {
        if m.Checked(idx) {
            cv = append(cv, item)
        }
    }
    return cv
}

func (m *ConfigVersionModel) RowCount() int {
    return len(m.items)
}

func (m *ConfigVersionModel) Value(row, col int) interface{} {
    item := m.items[row]

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

// 排序
// func (m *ConfigVersionModel) Sort(col int, order walk.SortOrder) error {
//     m.sortColumn, m.sortOrder = col, order

//     sort.Stable(m)

//     return m.SorterBase.Sort(col, order)
// }

func (m *ConfigVersionModel) Len() int {
    return len(m.items)
}

func (m *ConfigVersionModel) Less(i, j int) bool {
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

func (m *ConfigVersionModel) Swap(i, j int) {
    m.items[i], m.items[j] = m.items[j], m.items[i]
}

func (m *ConfigVersionModel) ResetRows(conf *ConfigVersion) {
    m.items = append(m.items, &ConfigVersion{
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
    m.PublishRowsReset()

    m.Sort(m.sortColumn, m.sortOrder)
}

func (m *ConfigVersionModel) Read() (err error) {
    db, err := openDb()
    // defer db.Close()
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
            fmt.Println("找到不到历史版本记录了.")
            return errors.New("找到不到历史版本记录了.")
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

func openDb() (db *sql.DB, err error) {
    // db配置文件是否存在
    dbExsit, err := os.Open(dbFile)
    defer dbExsit.Close()
    if err != nil {
        // 初始化db配置文件
        err = initDb()
        if err != nil {
            return nil, err
        }
    }

    // 打开db配置文件
    db, err = sql.Open("sqlite3", dbFile)
    // defer db.Close()
    if err != nil {
        return nil, err
    }

    return
}

func initDb() (err error) {
    db, err := sql.Open("sqlite3", dbFile)
    defer db.Close()
    if err != nil {
        return err
    }

    sqlInit := `
    create table setting (ip text not null primary key, user text, passwd text);
    delete from setting;
    create table version ("index" int not null primary key, masterver text, version text, pack text, tag text, tagpath text, packtime datatype);
    delete from version;
    `
    _, err = db.Exec(sqlInit)
    if err != nil {
        return err
    }

    return nil
}
