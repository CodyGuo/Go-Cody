package main

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func Usage() {
	arg_num := len(os.Args)
	if arg_num < 2 {
		fmt.Print(
			"Please runAs [super user] in [terminal].\n",
			"Usage:\n",
			"\tgoping url\n",
			"\texample: goping www.baidu.com",
		)
		time.Sleep(5e9)
		os.Exit(1)
	}
}

func main() {
	Usage()
	var (
		icmp     ICMP
		laddr    = net.IPAddr{IP: net.ParseIP("0.0.0.0")}
		raddr, _ = net.ResolveIPAddr("ip", os.Args[1])
	)
	conn, err := net.DialIP("ip4:icmp", &laddr, raddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer conn.Close()

	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.Checksum = CheckSum(buffer.Bytes())
	buffer.Reset()

	binary.Write(&buffer, binary.BigEndian, icmp)

	fmt.Printf("\n正在 Ping %s 具有0字节的数据:\n", raddr.String())
	recv := make([]byte, 1024)
	statistic := list.New()
	sended_packets := 0
	for i := 4; i > 0; i-- {
		if _, err := conn.Write(buffer.Bytes()); err != nil {
			fmt.Println(err.Error())
			return
		}
		sended_packets++
		t_start := time.Now()
		conn.SetReadDeadline((time.Now().Add(5 * time.Second)))
		_, err := conn.Read(recv)
		if err != nil {
			fmt.Println("请求超时")
			continue
		}
		t_end := time.Now()
		dur := t_end.Sub(t_start).Nanoseconds() / 1e6
		fmt.Printf("来自 %s 的回复: 时间 = %dms\n", raddr.String(), dur)
		statistic.PushBack(dur)
	}

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
