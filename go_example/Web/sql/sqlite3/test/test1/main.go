package main

import (
	"fmt"
)

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./user.db")
	checkErr(err)
	defer db.Close()

	ctuserinfo := `CREATE TABLE "userinfo" (
"uid" INTEGER PRIMARY KEY AUTOINCREMENT,
"username" VARCHAR(64) NULL,
"departname" VARCHAR(64) NULL,
"created" DATE NULL
)`

	ctuserdeatail := `CREATE TABLE "userdeatail" (
"uid" INT(10) NULL,
"intro" TEXT NULL,
"profile" TEXT NULL,
PRIMARY KEY ("uid")
)`

	_, err = db.Exec(ctuserinfo)
	checkErr(err)
	_, err = db.Exec(ctuserdeatail)
	checkErr(err)
	fmt.Println("done")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
