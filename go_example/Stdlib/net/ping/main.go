package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	ICMP_ECHO_REQUEST = 8
	ICMP_ECHO_REPLY   = 0
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host", os.Args[0])
		os.Exit(1)
	}

	dst := os.Args[1]
	raddr, err := net.ResolveIPAddr("ip4", dst)
	checkError(err)
	ipconn, err := net.DialIP("ip4:icmp", nil, raddr)
	checkError(err)

	sendid := os.Getpid() & 0xffff
	sendseq := 1
	pingpktlen := 64
	for {
		sendpkt := makePingRequest(sendid, sendseq, pingpktlen, []byte(""))
		start := int64(time.Now().Nanosecond())
		_, err := ipconn.WriteToIP(sendpkt, raddr)
		checkError(err)

		resp := make([]byte, 1024)
		for {
			n, from, err := ipconn.ReadFrom(resp)
			checkError(err)
			fmt.Printf("%d bytes from %s: icmp_req = %d time = %.2f ms\n", n,
				from, sendseq, elapsedTime(start))

			if resp[0] != ICMP_ECHO_REPLY {
				continue
			}

			rcvid, rcvseq := parsePingReply(resp)
			if rcvid != sendid || rcvseq != sendseq {
				fmt.Println("Ping reply saw id ", rcvid, rcvseq, sendid, sendseq)
			}

			break
		}

		if sendseq == 4 {
			break
		} else {
			sendseq++
		}

		time.Sleep(1e9)
	}

}

func makePingRequest(id, seq, pktlen int, filler []byte) []byte {
	p := make([]byte, pktlen)
	copy(p[8:], bytes.Repeat(filler, (pktlen-8)/(len(filler)+1)))
	p[0] = ICMP_ECHO_REQUEST
	p[1] = 0
	p[2] = 0
	p[3] = 0
	p[4] = uint8(id >> 8)
	p[5] = uint8(id & 0xff)
	p[6] = uint8(seq >> 8)
	p[7] = uint8(seq & 0xff)
	chlen := len(p)
	s := uint32(0)
	for i := 0; i < (chlen - 1); i += 2 {
		s += uint32(p[i+1]<<8) | uint32(p[i])
	}

	if chlen&1 == 1 {
		s += uint32(p[chlen-1])
	}

	s = (s >> 16) + (s & 0xffff)
	s = s + (s >> 16)
	p[2] ^= uint8(^s & 0xff)
	p[3] ^= uint8(^s >> 8)

	return p
}

func parsePingReply(p []byte) (id, seq int) {
	id = int(p[4])<<8 | int(p[5])
	seq = int(p[6])<<8 | int(p[7])

	return
}

func elapsedTime(start int64) float32 {
	t := float32((int64(time.Now().Nanosecond()) - start) / 1000000.0)

	return t
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
