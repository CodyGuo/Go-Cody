package main

import (
    "errors"
    "fmt"
)

var (
    ConfSer *ConfigServer
)

func init() {
    ConfSer = NewConfigServer()
    ConfSer.Read()
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

    // 配置为空就插入服务器配置
    if err = cs.Read(); err != nil {
        fmt.Println("configServer - insert:", ip, user, passwd)
        SqlStmt, _ = template.Prepare("insert into setting(ip, user, passwd) values(?, ?, ?)")
        _, err = SqlStmt.Exec(ip, user, passwd)
    } else {
        // 配置存在就更新配置
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
