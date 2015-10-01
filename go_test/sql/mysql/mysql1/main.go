package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type TestMysql struct {
    db *sql.DB
}

func Init() (*TestMysql, error) {
    test := new(TestMysql)
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hupuwebsite?charset=utf8")
    if err != nil {
        return nil, err
    }

    test.db = db
    return test, nil
}

func (t *TestMysql) Read() {
    if t.db == nil {
        return
    }
    rows, err := t.db.Query("select sname,saccount,spassword from tadmin where iadminid >= ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    cols, _ := rows.Columns()
    for i := range cols {
        fmt.Println(cols[i])
    }
    var sname string
    var saccount string
    var spassword string
    for rows.Next() {
        err := rows.Scan(&sname, &saccount, &spassword)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("name: ", sname, "account:", saccount, "spassword:", spassword)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

}

func (t *TestMysql) Write() {
    stmt, err := t.db.Prepare("insert into tadmin(sname, saccount, spassword, sifinlay, sdesc)values(?,?,?,?,?)")
    if err != nil {
        log.Println(err)
    }
    rs, err := stmt.Exec("普通管理员", "codyguo", "123", "2", "我用的.")
    if err != nil {
        log.Println(err)
    }
    //我们可以获得插入的id
    id, _ := rs.LastInsertId()
    //可以获得影响行数
    affect, _ := rs.RowsAffected()
    fmt.Println("id:", id, "affect: ", affect)
}
func main() {
    if test, err := Init(); err == nil {
        test.Read()
        test.Write()
        test.Read()
    } else {
        log.Fatal(err)
    }

}
