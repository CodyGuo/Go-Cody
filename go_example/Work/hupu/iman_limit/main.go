package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	exec.Command("cmd", "/c", "title iMan清空登陆限制工具").Run()

	var ip string
	fmt.Print("请输入iMan服务器IP: ")
	fmt.Scanf("%s\n", &ip)

	db, err := sql.Open("mysql", "root:hupu12iman!@tcp("+fmt.Sprint(ip)+":3306)/hupunac?charset=utf8")
	checkErr(err)
	defer db.Close()

	db.Query("TRUNCATE TABLE tloginmanage")
	db.Query("TRUNCATE TABLE tloginmanageother")

	fmt.Println("iMan所有平台登陆限制已清空，请重新使用WEB登陆。")
	var tmp string
	fmt.Scan(&tmp)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
