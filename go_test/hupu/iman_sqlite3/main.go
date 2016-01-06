package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
    // "time"
)

const (
    LOGINFO    string = "提示信息"
    LOGWARRING string = "警告信息"
    LOGERROR   string = "错误信息"
)

type Slogs struct {
    LogType   int
    WriteTime string
    LogInfo   string
}

var sysPath, dbFile, logtype string

func init() {
    sysPath = os.Getenv("systemroot")
    dbFile = sysPath + "\\hpNacIdm\\nacdata.db"
}

func main() {
    db, err := sql.Open("sqlite3", dbFile)
    checkErr(err)
    defer db.Close()

    rows, err := db.Query("select * from systemlog")
    checkErr(err)
    defer rows.Close()

    var syslogs []Slogs = make([]Slogs, 0)
    for rows.Next() {
        var slog Slogs
        rows.Scan(&slog.LogType, &slog.WriteTime, &slog.LogInfo)
        syslogs = append(syslogs, slog)
    }

    for k, v := range syslogs {
        switch v.LogType {
        case 1:
            logtype = LOGINFO
        case 2:
            logtype = LOGWARRING
        case 3:
            logtype = LOGERROR
        }

        fmt.Printf("序号: %d | 日志类型: %s | 时间: %s | 信息: %s\n", k+1, logtype, v.WriteTime, v.LogInfo)
    }

    fmt.Println("请手动关闭窗口....")
    var pause string
    fmt.Scan(&pause)
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
