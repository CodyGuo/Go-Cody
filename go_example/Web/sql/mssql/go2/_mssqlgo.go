func (b Box) mssqlReadWrite(name string, zid []string) {

	db := Mssql{
		dataSource: "10.184.31.143",
		database:   "A01_JiZhanKu",
		// windwos: true Ϊwindows�����֤��false ��������sa�˺ź�����
		windows: false,
		sa: SA{
			user:   "sa",
			passwd: "sa123!",
			port:   1433,
		},
	}
	// �������ݿ�
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()
	// ִ��SQL���

	wm := b.wm
	arrval := b.arrval

	sql := "delete from " + name + " where EXCHID='" + wm + "'"
	db.Query(sql)

	t := time.Now().Unix()
	t2 := strings.Split(time.Unix(t, 0).String(), "+")[0]
	for _, vals := range arrval {
		a1 := strings.Split(vals, "|")
		sql = "INSERT INTO " + name
		sql = sql + "(EXCHID,"
		for _, d := range zid {
			sql = sql + d + ","
		}
		sql = sql + "DATE)"
		sql = sql + " VALUES("
		sql = sql + "'" + wm + "'," //0
		for i := 0; i < len(a1); i++ {
			sql = sql + "'" + a1[i] + "',"
		}
		sql = sql + "'" + t2 + "')"
		db.Query(sql)
	}

}