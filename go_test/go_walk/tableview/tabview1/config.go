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
    dbFile  = "./iman.db"
    ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

var (
    _VERSION_ = "cody.guo"

    ConfVer      *ConfigVersion
    SqlStmt      *sql.Stmt
    ConfVerModel *ConfigVersionModel
)

func init() {

    ConfVer = NewConfgVersion()
    ConfVerModel = NewConfigVersionModel()

    ConfVer.Read()

}

type ConfigVersion struct {
    Index     int
    Version   string
    MasterVer string
    Pack      string
    Tag       string
    TagPath   string
    PackTime  time.Time
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
    m := new(ConfigVersionModel)
    // m.ResetRows()
    return m
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

func (m *ConfigVersionModel) ResetRows(n int, conf *ConfigVersion) {
    // Create some random data.
    // m.items = append(m.items, conf)
    m.items = make([]*ConfigVersion, 10000)

    m.items[n] = &ConfigVersion{
        Index:     conf.Index,
        MasterVer: conf.MasterVer,
        Version:   conf.Version,
        Pack:      conf.Pack,
        Tag:       conf.Tag,
        TagPath:   conf.TagPath,
        // PackTime:  conf.PackTime,
    }
    fmt.Println("m.timets", m.items[0], len(m.items))

    // now := time.Now()

    // for i := range m.items {
    //     m.items[i] = &ConfigVersion{
    //         Index:     i,
    //         MasterVer: "3",
    //         Version:   "3.20",
    //         Pack:      "pack",
    //         Tag:       "tag",
    //         TagPath:   "tagpath",
    //         PackTime:  time.Unix(rand.Int63n(now.Unix()), 0),
    //     }
    // }

    // Notify TableView and other interested parties about the reset.
    m.PublishRowsReset()

    m.Sort(m.sortColumn, m.sortOrder)
}

func (cv *ConfigVersion) Read() (err error) {
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

    var n int
    for {
        switch rows.Next() {
        case false:
            fmt.Println("找到不到历史版本记录了.")
            return errors.New("找到不到历史版本记录了.")
        case true:
            rows.Scan(&cv.Index, &cv.MasterVer, &cv.Version, &cv.Pack, &cv.Tag, &cv.TagPath, &cv.PackTime)
            ConfVerModel.ResetRows(n, cv)
            fmt.Println("n:", n)
            n += 1
            fmt.Println("version - read:", cv.Index, cv.MasterVer, cv.Version, cv.Pack, cv.Tag, cv.TagPath, cv.PackTime)
        }
    }

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
