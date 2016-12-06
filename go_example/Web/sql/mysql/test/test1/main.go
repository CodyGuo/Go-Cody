package main

import (
	"fmt"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:hupu12iman!@tcp(10.10.2.227:3306)/hupunac?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT ideviceid, sdeviceip, sdevicemac FROM tdevice")
	checkErr(err)

	columns, err := rows.Columns()
	checkErr(err)

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		checkErr(err)

		var value string
		for i, col := range values {
			if col == nil {
				value = ""
			} else {
				value = string(col)
			}
			fmt.Printf("%s: %s  ", columns[i], value)
		}
		fmt.Println()
	}
	defer rows.Close()
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
