package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var serverIp string
	fmt.Print("请输入iMan服务器IP: ")
	fmt.Scan(&serverIp)

	db, err := sql.Open("mysql", "root:hupu12iman!@tcp("+serverIp+":3306)/hupunac")
	checkErr(err)
	defer db.Close()

	selectTb(db, "限制登录服务器的IP", "tloginmanage")
	fmt.Println("=================================================")
	selectTb(db, "限制登录控制器和云平台的IP", "tloginmanageother")
	fmt.Println("=================================================")
	fmt.Println("\n检查iMan登录限制结束.")

}

func selectTb(db *sql.DB, title, tb string) {
	rows, err := db.Query("SELECT sloginip as IP, sremark as Remark FROM " + tb)
	checkErr(err)

	fmt.Printf("\n%-15s  %s\n", title, "备注")
	var count int
	for rows.Next() {
		var ip string
		var remark string
		err = rows.Scan(&ip, &remark)
		checkErr(err)
		if ip != "" {
			count++
		}

		fmt.Printf("    %-15s  --  %s\n", ip, remark)
	}

	clearTable(db, tb, count)
}

func clearTable(db *sql.DB, tb string, count int) {
	var yesNo string
	if count >= 1 {
	OuterLoop:
		for {
			fmt.Print("是否要清空yes or no: ")
			fmt.Scan(&yesNo)
			switch yesNo {
			case "y", "yes", "YES":
				_, err := db.Query("TRUNCATE TABLE " + tb)
				checkErr(err)
				fmt.Println("清空限制成功.")
				break OuterLoop
			case "n", "no", "NO":
				break OuterLoop
			default:
				continue
			}
		}
	}

}
