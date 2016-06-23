package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	return uint16(^sum)
}
func main() {
	var (
		icmp  ICMP
		laddr net.IPAddr = net.IPAddr{IP: net.ParseIP("192.168.34.125")}
		raddr net.IPAddr = net.IPAddr{IP: net.ParseIP("114.114.114.114")}
	)
	conn, err := net.DialIP("ip4:icmp", &laddr, &raddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	icmp.Type = 8 //8->echo message  0->reply message
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0
	var (
		buffer bytes.Buffer
	)
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("send icmp packet success!")
}
