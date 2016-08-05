package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const tag string = "-----------------------------------"

var deviceInof map[int]hupuTable

type hupuTable struct {
	db         *sql.DB
	SIP        string
	ID         int
	IP         string
	MAC        string
	ControlMAC string
}

type Result struct {
	Success            bool
	Info               string
	ProcessMillisecond int
}

func DBConnect(ip string) (*hupuTable, error) {
	dataSourceName := fmt.Sprintf("root:hupu12iman!@tcp(%s:3306)/hupunac?charset=utf8", ip)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	hupuTab := new(hupuTable)
	hupuTab.db = db
	hupuTab.SIP = ip

	return hupuTab, nil
}

func (h hupuTable) query() {
	rows, err := h.db.Query("SELECT ideviceid, sdeviceip, sdevicemac, sascuniqueid FROM tdevice WHERE sascuniqueid <> '' and sdeviceip <> ''")
	checkError(err, "rows 查询设备信息")
	defer rows.Close()

	deviceInof = make(map[int]hupuTable)
	for rows.Next() {
		err := rows.Scan(&h.ID, &h.IP, &h.MAC, &h.ControlMAC)
		checkError(err, "rows 获取设备信息")
		deviceInof[h.ID] = h
	}

	rowsNoneIP, err := h.db.Query("SELECT ideviceid, sdevicemac, sascuniqueid FROM tdevice WHERE sascuniqueid IS NOT NULL AND sascuniqueid != '' AND sascuniqueid <> '' AND (sdeviceip IS NULL OR sdeviceip = '')")
	checkError(err, "rowsNoneIP 查询设备信息")
	defer rowsNoneIP.Close()

	for rowsNoneIP.Next() {
		err := rowsNoneIP.Scan(&h.ID, &h.MAC, &h.ControlMAC)
		checkError(err, "rowsNoneIP 获取设备信息")
		h.IP = "0.0.0.0"
		deviceInof[h.ID] = h
		// fmt.Printf("没有IP的设备： %d %s %s\n", h.ID, h.MAC, h.ControlMAC)
	}
	h.db.Close()
}

func (h *hupuTable) AutoOnline() {
	h.query()
	for _, d := range deviceInof {
		err, info := h.httpGet(d.ID, d.IP, d.MAC, d.ControlMAC)
		if err != nil {
			fmt.Printf("[ERROR] 设备[ %-15s ][%18s ] %s.\n", d.IP, d.MAC, info)
		} else {
			fmt.Printf("[INFO ] 设备[ %-15s ][%18s ] %s.\n", d.IP, d.MAC, info)
		}
		// fmt.Printf("%5d %-15s %18s %18s \n", d.ID, d.IP, d.MAC, d.ControlMAC)
	}

}

func (h *hupuTable) httpGet(id int, ip, mac, conMac string) (error, string) {
	url := fmt.Sprintf("http://%s/approve/Users-deviceAuthByControl?asmType=2&authItemStr=&authPolicyId=1&authResult=0&controlMac=%s&criticalOnline=false&deviceip=%s&devicemac=%s&iauthidentity=0&ideptid=&ideviceid=%d&ifGuestAuth=false&iguestid=&iguesttypeid=&iuserid=&natType=0&scompanycode=10000000", h.SIP, conMac, ip, mac, id)
	resp, err := http.Get(url)
	if err != nil {
		return err, "服务器不可达."
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response falied."), "服务器响应失败."
	}

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err, "解析Json失败."
	}

	if !result.Success {
		return fmt.Errorf("%b", result.Success), result.Info
	}

	return nil, result.Info
}

func main() {
	var serverIP string
	for {
		fmt.Print("请输入iMan服务器IP: ")
		fmt.Scanf("%s\n", &serverIP)
		if serverIP != "" {
			break
		}
	}

	iman, err := DBConnect(serverIP)
	checkError(err, "连接数据库")
	iman.AutoOnline()
	fmt.Printf("%s\n", tag)
	fmt.Println("设备自动上线结束,请手动关闭窗口.")

	for {
	}

}

func checkError(err error, info string) {
	if err != nil {
		fmt.Printf("%s失败. %s\n", info, err)
		os.Exit(1)
	}
}
