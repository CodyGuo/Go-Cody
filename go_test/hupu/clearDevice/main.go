package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	if runtime.GOOS == "windows" {
		exec.Command("cmd", "/c", "cls & title", "清理故障设备信息工具").Run()
	}
}

func main() {
	var serverIp, sIP, sMac string
	fmt.Print("请输入iMan服务器IP: ")
	fmt.Scanf("%s\n", &serverIp)

	for {
		fmt.Print("请输入设备的IP MAC，以空格分隔: ")
		n, err := fmt.Scanf("%s %s\n", &sIP, &sMac)
		if n >= 2 && err == nil && sIP != "" && sMac != "" {
			break
		}
		fmt.Print("\n")
	}

	db, err := sql.Open("mysql", "root:hupu12iman!@tcp("+serverIp+":3306)/hupunac")
	checkErr(err)
	defer db.Close()

	clearDevice(db, sIP, sMac)

	var pause string
	log.Println("请手动关闭工具窗口.")
	fmt.Scan(&pause)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func clearDevice(db *sql.DB, ip, mac string) {
	_, err := db.Query("DELETE FROM tdevice WHERE sdeviceip = ? OR sdevicemac = ?", ip, mac)
	checkErr(err)
	fmt.Println("设备", ip, mac, "清理结束.")
	fmt.Println("------------------------------------------")
}
