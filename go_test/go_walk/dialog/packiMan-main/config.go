package main

import (
    "errors"
    "fmt"
    "os"
    "time"

    "database/sql"
)

import (
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
        rows.Scan(&cv.Index, &cv.Version, &cv.MasterVer, &cv.Pack, &cv.Tag, &cv.TagPath, &cv.PackTime)
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
