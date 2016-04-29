package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

import (
	_ "github.com/mattn/go-adodb"
)

func main() {
	db := Mssql{
		// 如果数据库是默认实例（MSSQLSERVER）则直接使用IP，命名实例需要指明。
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

	var now time.Time
	// now = time.Now()
	// db.Insert()
	// fmt.Printf("[INFO] 插入用时 %v\n", time.Since(now))

	// now = time.Now()
	// db.Delete()
	// fmt.Printf("[INFO] 插入用时 %v\n", time.Since(now))

	// now = time.Now()
	// db.Inseret()
	// fmt.Printf("[INFO] 插入用时 %v\n", time.Since(now))

	now = time.Now()
	db.Select("info")
	fmt.Printf("[INFO] 查询用时 %v\n", time.Since(now))

}

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA

	dateIn []string
}

type SA struct {
	user   string
	passwd string
	port   int
}

func (m *Mssql) Open() (err error) {
	var buf bytes.Buffer
	buf.WriteString("Provider=SQLOLEDB" + ";" +
		"Data Source=" + m.dataSource + ";" +
		"Initial Catalog=" + m.database + ";")

	// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器
	// (需要在安装sqlserver时候设置)，
	// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
	if m.windows {
		buf.WriteString("integrated security=SSPI" + ";")
	} else {
		buf.WriteString("user id=" + m.sa.user + ";" +
			"password=" + m.sa.passwd + ";" +
			"port=" + fmt.Sprint(m.sa.port))
	}

	m.DB, err = sql.Open("adodb", buf.String())
	if err != nil {
		return err
	}
	return nil
}

func (m *Mssql) Insert(table string, columns []string) {
	for _, in := range m.dateIn {
		stmt, err := m.Prepare("INSERT info " + table + "(" + strings.Join(columns, ",") + ")")
	}
}

func (m *Mssql) Delete() {

}

func (m *Mssql) Update() {

}

func (m *Mssql) Select(table string) {
	// 执行SQL语句
	rows, err := m.Query("select name from " + table)
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	for rows.Next() {
		// 查询结果字段和声明变量数量相等，否则数据为空。
		var name string
		rows.Scan(&name)
		fmt.Printf("Name: %s\n", name)
		// var name string
		// var number int
		// rows.Scan(&name, &number)
		// fmt.Printf("Name: %s \t Number: %d\n", name, number)
	}
}
