package main

import (
	"fmt"
)

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Tdevice struct {
	Ideviceid  int `gorm:"primary_key"`
	Sdevicemac string
}

func (t Tdevice) TableName() string {
	return "tdevice"
}

func main() {
	db, err := gorm.Open("mysql", "root:hupu12iman!@tcp(10.10.2.227:3306)/hupunac?charset=utf8")
	checkErr(err)
	defer db.Close()
	db.LogMode(true)

	checkTable := "tdevice"
	fmt.Println(checkTable, db.HasTable(checkTable))

	var tdevice Tdevice

	db.Table("tdevice").Where("ideviceid = ?", 6505).First(&tdevice)
	fmt.Println(tdevice)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
