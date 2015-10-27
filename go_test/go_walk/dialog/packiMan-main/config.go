package main

import (
    "errors"
    "fmt"
    "os"
    "time"

    "database/sql"
)

import (
    // "github.com/lxn/walk"
    _ "github.com/mattn/go-sqlite3"
)

const (
    dbFile  = "./iman.db"
    ipRegxp = "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-4])$"
)

var (
    _VERSION_ = "cody.guo"

    ConfSer *ConfigServer
    ConfVer *ConfigVersion
    SqlStmt *sql.Stmt
)

func init() {
    ConfSer = NewConfigServer()
    ConfSer.Read()

    ConfVer = NewConfgVersion()
    ConfVer.Read()
}

type ConfigServer struct {
    Ip     string
    User   string
    Passwd string
}

func NewConfigServer() *ConfigServer {
    return new(ConfigServer)
}

func (cs *ConfigServer) Read() (err error) {
    db, err := openDb()
    defer db.Close()
    if err != nil {
        return err
    }

    rows, err := db.Query("select * from setting")
    defer rows.Close()
    if err != nil {
        return err
    }

    switch rows.Next() {
    case false:
        fmt.Println("服务器设置为空.")
        return errors.New("服务器设置为空.")
    case true:
        rows.Scan(&cs.Ip, &cs.User, &cs.Passwd)
    }

    fmt.Println("configServer - read:", cs.Ip, cs.User, cs.Passwd)

    return nil
}

func (cs *ConfigServer) Write() (err error) {
    db, err := openDb()
    defer db.Close()
    if err != nil {
        return err
    }

    template, err := db.Begin()
    if err != nil {
        return err
    }

    var ip, user, passwd string
    ip = cs.Ip
    user = cs.User
    passwd = cs.Passwd

    if err = cs.Read(); err != nil {
        fmt.Println("configServer - insert:", ip, user, passwd)
        SqlStmt, _ = template.Prepare("insert into setting(ip, user, passwd) values(?, ?, ?)")
        _, err = SqlStmt.Exec(ip, user, passwd)

    } else {
        fmt.Println("configServer - update:", ip, user, passwd, cs.Ip)
        SqlStmt, _ = template.Prepare("update setting set ip=?, user=?, passwd=? where ip=?")
        _, err = SqlStmt.Exec(ip, user, passwd, cs.Ip)
    }
    defer SqlStmt.Close()
    if err != nil {
        return err
    }

    template.Commit()

    return nil
}

type ConfigVersion struct {
    Index     int
    Version   string
    MasterVer string
    Pack      string
    Tag       string
    TagPath   string
    PackTime  time.Duration
}

// type ConfigVersionModel struct {
//     walk.TableModelBase
//     walk.SorterBase
//     sortColumn int
//     sortOrder  walk.SortOrder
//     items      []*ConfigVersion
// }

// func NewConfigVersionModel() *ConfigVersionModel {
//     m := new(ConfigVersionModel)
//     m.ResetRows()
//     return m
// }

// func (m *ConfigVersionModel) RowCount() int {
//     return len(m.items)
// }

// func (m *ConfigVersionModel) Value(row, col int) interface{} {
//     item := m.items[row]

//     switch col {
//     case 0:
//         return item.Index

//     case 1:
//         return item.Version

//     case 2:
//         return item.MasterVer

//     case 3:
//         return item.Pack
//     case 4:
//         return item.Tag
//     case 5:
//         return item.TagPath
//     case 6:
//         return item.PackTime
//     }

//     panic("unexpected col")
// }

// func (m *ConfigVersionModel) Checked(row int) bool {
//     return m.items[row].checked
// }

// func (m *ConfigVersionModel) SetChecked(row int, checked bool) error {
//     m.items[row].checked = checked

//     return nil
// }

// func (m *ConfigVersionModel) Sort(col int, order walk.SortOrder) error {
//     m.sortColumn, m.sortOrder = col, order

//     sort.Stable(m)

//     return m.SorterBase.Sort(col, order)
// }

// func (m *ConfigVersionModel) Len() int {
//     return len(m.items)
// }

// func (m *ConfigVersionModel) Less(i, j int) bool {
//     a, b := m.items[i], m.items[j]
//     c := func(ls bool) bool {
//         if m.sortOrder == walk.SortAscending {
//             return ls
//         }
//         return !ls
//     }

//     switch m.sortColumn {
//     case 0:
//         return c(a.Index < b.Index)
//     case 1:
//         return c(a.Version < b.Version)
//     case 2:
//         return c(a.MasterVer < b.MasterVer)
//     case 3:
//         return c(a.Pack < b.Pack)
//     case 4:
//         return c(a.Tag < b.Tag)
//     case 5:
//         return c(a.TagPath < b.TagPath)
//     case 6:
//         return c(a.PackTime.Before(b.PackTime))
//     }

//     panic("unreachable")
// }

// func (m *ConfigVersionModel) Swap(i, j int) {
//     m.items[i], m.items[j] = m.items[j], m.items[i]
// }

func NewConfgVersion() *ConfigVersion {
    return new(ConfigVersion)
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

    switch rows.Next() {
    case false:
        fmt.Println("历史版本记录为空.")
        return errors.New("历史版本记录为空.")
    case true:
        for rows.Next() {
            rows.Scan(&cv.Index, &cv.Version, &cv.MasterVer, &cv.Pack, &cv.Tag, &cv.TagPath, &cv.PackTime)
        }
    }

    fmt.Println("version - read:", cv.Index, cv.Version, cv.MasterVer, cv.Pack, cv.Tag, cv.TagPath, cv.PackTime)

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
    create table version ("index" int not null primary key, version text, masterver text, pack text, tag text, tagpath text, packtime datatype);
    delete from version;
    `
    _, err = db.Exec(sqlInit)
    if err != nil {
        return err
    }

    return nil
}
