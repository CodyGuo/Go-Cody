package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	// Svs     []server `xml:"server"`
	// ServerName []string `xml:"server>serverName"`
	// ServerIP   []string `xml:"server>serverIP"`
}

// type server struct {
// 	ServerName string `xml:"serverName"`
// 	ServerIP   string `xml:"serverIP"`
// }

func main() {
	v := &Servers{Version: "1"}
	// v.ServerName = append(v.ServerName, "上海")
	// v.ServerName = append(v.ServerName, "北京")
	// v.ServerIP = append(v.ServerIP, "127.0.0.1")
	// v.ServerIP = append(v.ServerIP, "127.0.0.2")
	// v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	// v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
