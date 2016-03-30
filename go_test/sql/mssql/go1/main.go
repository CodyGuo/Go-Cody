package main

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器
		// (需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)
	conf = append(conf, "port="+fmt.Sprint(m.sa.port))

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db := Mssql{
		dataSource: "10.10.2.140\\SQLEXPRESS",
		database:   "test",
		// windows: true 为windows身份验证，false 必须设置sa账号和密码
		windows: false,
		sa: SA{
			user:   "sa",
			passwd: "123456",
			port:   1433,
		},
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()

	// 执行SQL语句
	rows, err := db.Query("select name from info")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	for rows.Next() {
		//结果几个字段就要命名几个变量
		var name string
		rows.Scan(&name)
		fmt.Printf("Name: %s\n", name)
		// var number int
		// rows.Scan(&name, &number)
		// fmt.Printf("Name: %s \t Number: %d\n", name, number)
	}
}
