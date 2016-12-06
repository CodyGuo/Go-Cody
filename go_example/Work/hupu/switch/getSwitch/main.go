package main

import (
	"fmt"
)

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type Switch struct {
	iswitchportid int
	iswitchid     int
	sport         string
	sifmanaged    int
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:hupu12iman!@tcp(10.10.3.227:3306)/hupunac?charset=utf8")
	checkErr(err)
	defer engine.Close()
	ok, err := engine.IsTableExist("tdevice")
	checkErr(err)
	if ok {
		fmt.Println("表存在")
	} else {
		fmt.Println("表不存在")
		return
	}

	var switchPort []Switch
	engine.Where("sifmanaged = ?", 1).Find(&switchPort)
	fmt.Println(switchPort)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
