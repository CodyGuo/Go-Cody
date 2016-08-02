package main

import (
    "os"

    "database/sql"
)

import (
    _ "github.com/mattn/go-sqlite3"
)

var (
    SqlStmt *sql.Stmt
)

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
