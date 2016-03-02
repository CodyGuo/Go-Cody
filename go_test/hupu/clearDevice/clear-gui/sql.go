package main

import (
	"fmt"
)
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ipFlag = iota
	macFlag
)

type SQL struct {
	user     string
	passwd   string
	port     int
	database string
	serverIP string
	DB       *sql.DB
}

func doSQL(serverIP string, flag int, str string) error {
	sqlConn := new(SQL)
	sqlConn.serverIP = serverIP
	sqlConn.user = "root"
	sqlConn.passwd = "hupu12iman!"
	sqlConn.port = 3306
	sqlConn.database = "hupunac"

	db, err := sql.Open("mysql", sqlConn.user+":"+sqlConn.passwd+
		"@tcp("+sqlConn.serverIP+":"+fmt.Sprint(sqlConn.port)+")/"+
		sqlConn.database)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlConn.DB = db
	switch flag {
	case ipFlag:
		return sqlConn.doClearIP(str)
	case macFlag:
		return sqlConn.doClearMAC(str)
	}
	return nil
}

func (s *SQL) doClearIP(ip string) error {
	_, err := s.DB.Query("DELETE FROM tdevice WHERE sdeviceip = ?", ip)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQL) doClearMAC(mac string) error {
	_, err := s.DB.Query("DELETE FROM tdevice WHERE sdevicemac = ?", mac)
	if err != nil {
		return err
	}
	return nil
}
