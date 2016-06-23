package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)

import (
	_ "github.com/mattn/go-adodb"
)

var (
	local    bool
	remoteIP string
	remoteDS string
)

func init() {
	flag.BoolVar(&local, "local", true, "set window connect.")
	flag.StringVar(&remoteIP, "remoteIP", "192.168.1.104", "set up remote mssql of ip.")
	flag.StringVar(&remoteDS, "remoteDS", "MSSQLSERVER", "set up remote mssql of datasource.")
}

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         *SA
}

type SA struct {
	user   string
	passwd string
	port   int
}

func NewMssql() *Mssql {
	mssql := new(Mssql)
	dataS := "localhost"
	if !local {
		dataS = fmt.Sprintf("%s\\%s", remoteIP, remoteDS)
	}

	mssql = &Mssql{
		// 如果数据库是默认实例（MSSQLSERVER）则直接使用IP，命名实例需要指明。
		// dataSource: "192.168.1.104\\MSSQLSERVER",
		dataSource: dataS,
		database:   "iman",
		// windows: true 为windows身份验证，false 必须设置sa账号和密码
		windows: local,
		sa: &SA{
			user:   "sa",
			passwd: "123456",
			port:   1433,
		},
	}

	return mssql

}

func (m *Mssql) Open() error {
	config := fmt.Sprintf("Provider=SQLOLEDB;Initial Catalog=%s;Data Source=%s",
		m.database, m.dataSource)

	if m.windows {
		config = fmt.Sprintf("%s;Integrated Security=SSPI", config)
	} else {
		// sql 2000的端口写法和sql 2005以上的有所不同，在Data Source 后以逗号隔开。
		config = fmt.Sprintf("%s,%d;user id=%s;password=%s",
			config, m.sa.port, m.sa.user, m.sa.passwd)
	}

	var err error
	m.DB, err = sql.Open("adodb", config)
	fmt.Println(config)

	return err
}

func (m *Mssql) Select() {
	rows, err := m.Query("select id, name from users")
	if err != nil {
		fmt.Printf("select query err: %s\n", err)
	}

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		fmt.Printf("id = %s, name = %s\n", id, name)
	}
}

func main() {
	flag.Parse()

	mssql := NewMssql()
	err := mssql.Open()
	checkError(err)

	mssql.Select()

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
