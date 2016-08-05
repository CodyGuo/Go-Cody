package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Success            bool
	Info               string `json:"info"`
	ProcessMillisecond int
}

func main() {
	s := "http://10.10.3.227/approve/Users-deviceAuthByControl?asmType=2&authItemStr=&authPolicyId=1&authResult=0&controlMac=B0-51-8E-03-9C-93&criticalOnline=false&deviceip=10.10.2.112&devicemac=AC-BC-32-C9-5D-ED&iauthidentity=0&ideptid=&ideviceid=74&ifGuestAuth=false&iguestid=&iguesttypeid=&iuserid=&natType=0&scompanycode=10000000"
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("http get:", err)
	}
	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("json Decoder:", err)
	}

	resp.Body.Close()
	fmt.Println(result)

}
