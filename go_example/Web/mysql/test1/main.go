package main

import (
	"fmt"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("codyguo", "研发部", "2016-12-04")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("UPDATE userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("郭肖", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var (
			uid        int
			username   string
			department string
			created    string
		)
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}
	fmt.Println("over")
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
