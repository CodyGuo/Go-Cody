package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai", "10.10.1.1"},
		server{"Beijing", "10.10.2.1"})

	// output, err := xml.MarshalIndent(v, "", "    ")
	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}
