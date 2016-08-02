package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
)

type Setting struct {
    Ip     string
    User   string
    Passwd string
}

func main() {

    // os.Remove("./foo.db")
    //
    var settings Setting

    // settings := new(Setting)
    _, err := os.Open("./foo.db")
    if err == nil {
        log.Fatalln("文件已存在.")
    }

    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sqlStmt := `
    create table setting (ip text not null primary key, user text, passwd text);
    delete from setting;
    `
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }

    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    stmt, err := tx.Prepare("insert into setting(ip, user, passwd) values(?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    for i := 150; i < 200; i++ {
        _, err = stmt.Exec("10.10.2."+fmt.Sprint(i), "root", "passwd")
        if err != nil {
            fmt.Println("10.10.2."+fmt.Sprint(i), err.Error())
        }
    }

    tx.Commit()

    rows, err := db.Query("select * from setting")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        rows.Scan(&settings.Ip, &settings.User, &settings.Passwd)
        fmt.Println(settings.Ip, settings.User, settings.Passwd)
    }

}
