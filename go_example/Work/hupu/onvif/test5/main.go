package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 3))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	resp, err := client.Post("http://10.10.2.16/onvif/device_service", "application/soap+xml;charset=utf-8", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(string(data))

}
